package gojsonld

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Dataset struct {
	Graphs map[string][]*Triple
}

func NewDataset() *Dataset {
	dataset := &Dataset{}
	dataset.Graphs = make(map[string][]*Triple, 0)
	dataset.Graphs["@default"] = make([]*Triple, 0)
	return dataset
}

func (d *Dataset) Serialize() string {
	var result bytes.Buffer
	for name, triples := range d.Graphs {
		for _, triple := range triples {
			tripleString := serializeTriple(name, triple)
			result.WriteString(tripleString)
			result.WriteString("\n")
		}
	}
	return result.String()
}

func serializeTriple(name string, triple *Triple) string {
	if name == "@default" {
		return triple.String()
	}
	subjectString := "nil"
	if !isNil(triple.Subject) {
		subjectString = triple.Subject.String()
	}
	predicateString := "nil"
	if !isNil(triple.Predicate) {
		predicateString = triple.Predicate.String()
	}
	objectString := "nil"
	if !isNil(triple.Object) {
		objectString = triple.Object.String()
	}
	return fmt.Sprintf("%s %s %s %s .", subjectString, predicateString,
		objectString, name)
}

func ParseDataset(input []byte) (*Dataset, error) {
	dataset := NewDataset()
	readBuffer := bytes.NewBuffer(input)
	reader := bufio.NewReader(readBuffer)
	for {
		line, _, readErr := reader.ReadLine()
		if !isNil(readErr) {
			break
		}
		lineString := string(line)
		if len(lineString) == 0 {
			continue
		}
		parts := STATEMENT.SubexpNames()
		match := STATEMENT.FindAllStringSubmatch(lineString, -1)[0]
		partValues := make(map[string]string, 0)
		for index, matchValue := range match {
			partValues[parts[index]] = matchValue
		}
		subject, subjectErr := parseSubject(partValues["subject"])
		predicate, predicateErr := parsePredicate(partValues["predicate"])
		object, objectErr := parseObject(partValues["object"])
		var graph string
		var graphErr error
		if len(partValues["graph"]) > 0 {
			graph, graphErr = parseGraph(partValues["graph"])
		} else {
			graph, graphErr = "@default", error(nil)
		}
		if !isNil(subjectErr) {
			return nil, subjectErr
		}
		if !isNil(predicateErr) {
			return nil, predicateErr
		}
		if !isNil(objectErr) {
			return nil, objectErr
		}
		if !isNil(graphErr) {
			return nil, graphErr
		}
		appendTriple(dataset, graph, NewTriple(subject, predicate, object))
	}
	return dataset, nil
}

func parseSubject(value string) (Term, error) {
	if IRIREF.MatchString(value) {
		return NewResource(value[1:(len(value) - 1)]), nil
	} else if BLANK_NODE_LABEL.MatchString(value) {
		return NewBlankNode(value[2:]), nil
	} else {
		return nil, errors.New("Invalid subject")
	}
}

func parsePredicate(value string) (Term, error) {
	return parseSubject(value)
}

func parseObject(value string) (Term, error) {
	if LITERAL.MatchString(value) {
		return parseLiteral(value), nil
	} else if IRIREF.MatchString(value) {
		return NewResource(value[1:(len(value) - 1)]), nil
	} else if BLANK_NODE_LABEL.MatchString(value) {
		return NewBlankNode(value[2:]), nil
	} else {
		return nil, errors.New("Invalid subject")
	}
}

func parseLiteral(value string) Term {
	literalQuote := STRING_LITERAL_QUOTE.FindString(value)
	unescapedValue := unescapeValue(literalQuote[1:(len(literalQuote) - 1)])
	dataType := DATATYPE.FindString(value[len(literalQuote):])
	language := LANGTAG.FindString(value[len(literalQuote):])
	var dataTypeTerm Term
	if dataType == "" {
		dataTypeTerm = NewResource(XSD_STRING)
	} else {
		dataTypeTerm = NewResource(dataType[3:(len(dataType) - 1)])
	}
	if language != "" {
		language = language[1:]
	}
	return NewLiteralWithLanguageAndDatatype(unescapedValue, language, dataTypeTerm)
}

func parseGraph(value string) (string, error) {
	if IRIREF.MatchString(value) {
		return value[1:(len(value) - 1)], nil
	} else if BLANK_NODE_LABEL.MatchString(value) {
		return value[2:], nil
	} else {
		return "", errors.New("Invalid graph")
	}
	return "", nil
}

func appendTriple(dataset *Dataset, graph string, triple *Triple) {
	if _, hasGraph := dataset.Graphs[graph]; !hasGraph {
		dataset.Graphs[graph] = make([]*Triple, 0)
	}
	dataset.Graphs[graph] = append(dataset.Graphs[graph], triple)
}

func unescapeValue(value string) string {
	value = strings.Replace(value, "\\\\", "\\", -1)
	value = strings.Replace(value, "\\\"", "\"", -1)
	value = strings.Replace(value, "\\n", "\n", -1)
	value = strings.Replace(value, "\\r", "\r", -1)
	value = strings.Replace(value, "\\t", "\t", -1)
	return value
}

func (d *Dataset) Equal(other *Dataset) bool {
	dMap := createFrequencyMap(d)
	otherMap := createFrequencyMap(other)
	return reflect.DeepEqual(dMap, otherMap)
}

func createFrequencyMap(d *Dataset) map[string]int {
	frequencyMap := make(map[string]int, 0)
	serializedDataset := d.Serialize()
	readBuffer := bytes.NewBuffer([]byte(serializedDataset))
	reader := bufio.NewReader(readBuffer)
	for {
		line, _, readErr := reader.ReadLine()
		if !isNil(readErr) {
			break
		}
		lineString := string(line)
		if lineString == "" {
			continue
		}
		if _, hasLine := frequencyMap[lineString]; !hasLine {
			frequencyMap[lineString] = 0
		}
		frequencyMap[lineString] += 1
	}
	return frequencyMap
}

func (d *Dataset) ToGraphs() []*Graph {
	graphs := make([]*Graph, 0)
	for graph := range d.Graphs {
		newGraph := NewGraph(graph)
		for _, triple := range d.Graphs[graph] {
			newGraph.triples[triple] = true
		}
		graphs = append(graphs, newGraph)
	}
	return graphs
}

func (d *Dataset) IterTriples() (ch chan *Triple) {
	ch = make(chan *Triple)
	go func() {
		for graph := range d.Graphs {
			for _, triple := range d.Graphs[graph] {
				ch <- triple
			}
		}
		close(ch)
	}()
	return ch
}
