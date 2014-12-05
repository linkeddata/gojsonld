package gojsonld

import (
	"math"
	"strconv"
	"strings"
)

func toRDF(activeContext *Context, element interface{}) (*Dataset, error) {
	// 1)
	expanded, expandErr := Expand(element, activeContext.options)
	if !isNil(expandErr) {
		return nil, expandErr
	}
	// 2)
	nodeMap := make(map[string]interface{}, 0)
	var idGenerator = BlankNodeIdGenerator{}
	idGenerator.counter = 0
	idGenerator.identifierMap = make(map[string]string, 0)
	defaultArg := "@default"
	err := generateNodeMap(expanded, nodeMap, &defaultArg, nil,
		nil, nil, &idGenerator)
	if !isNil(err) {
		return nil, err
	}
	// 3)
	dataset := NewDataset()
	// 4)
	keys := sortedKeys(nodeMap)
	for _, graphName := range keys {
		graph := nodeMap[graphName]
		// 4.1)
		if isRelativeIri(graphName) {
			continue
		}
		// 4.2)
		triples := make([]*Triple, 0)
		// 4.3)
		graphMap := graph.(map[string]interface{})
		graphKeys := sortedKeys(graphMap)
		for _, id := range graphKeys {
			node := graphMap[id]
			// 4.3.1)
			if !isAbsoluteIri(id) {
				continue
			}
			nodeMapValue := node.(map[string]interface{})
			keysNode := sortedKeys(nodeMapValue)
			for _, property := range keysNode {
				values := nodeMapValue[property]
				// 4.3.2.1)
				if property == "@type" {
					property = RDF_TYPE
					// 4.3.2.2)
				} else if isKeyword(property) {
					continue
					// 4.3.2.3)
				} else if strings.HasPrefix(property, "_:") &&
					!activeContext.options.ProduceGeneralizedRdf {
					continue
					// 4.3.2.4)
				} else if isRelativeIri(property) {
					continue
				}
				// RDF subject
				var subject Term
				if strings.HasPrefix(id, "_:") {
					subject = NewBlankNode(id[2:])
				} else {
					subject = NewResource(id)
				}
				// RDF predicate
				var predicate Term
				if strings.HasPrefix(property, "_:") {
					predicate = NewBlankNode(property[2:])
				} else {
					predicate = NewResource(property)
				}
				valuesArray := values.([]interface{})
				for _, item := range valuesArray {
					if isListObject(item) {
						list := item.(map[string]interface{})["@list"].([]interface{})
						listHead, listTriples := listToRDF(list, &idGenerator)
						triples = append(triples, NewTriple(subject, predicate,
							listHead))
						for _, triple := range listTriples {
							triples = append(triples, triple)
						}
					} else {
						object := objectToRDF(item)
						if !isNil(object) {
							triples = append(triples, NewTriple(subject, predicate,
								object))
						}
					}
				}
			}
		}
		// 4.4 + 4.5)
		if strings.HasPrefix(graphName, "_:") {
			graphName = graphName[2:]
		}
		dataset.Graphs[graphName] = triples
	}
	return dataset, nil
}

func objectToRDF(item interface{}) Term {
	//TODO spec is wrong
	if itemString, isString := item.(string); isString {
		if strings.HasPrefix(itemString, "_:") {
			return NewBlankNode(itemString[2:])
		} else {
			return NewResource(itemString)
		}
	}
	id, hasID := item.(map[string]interface{})["@id"]
	language, hasLanguage := item.(map[string]interface{})["@language"]
	// 1)
	if isNodeObject(item) && hasID && isRelativeIri(id.(string)) {
		return nil
	}
	// 2)
	if isNodeObject(item) {
		if strings.HasPrefix(id.(string), "_:") {
			idString := id.(string)
			return NewBlankNode(idString[2:])
		} else {
			return NewResource(id.(string))
		}
	}
	// 3)
	value := item.(map[string]interface{})["@value"]
	// 4)
	datatype, hasDatatype := item.(map[string]interface{})["@type"]
	if !hasDatatype {
		datatype = ""
	}
	valueBool, isBool := value.(bool)
	valueFloat, isFloat := value.(float64)
	valueInt, isInt := value.(int64)
	// 5)
	if isBool {
		if valueBool {
			value = "true"
		} else {
			value = "false"
		}
		if datatype == "" {
			datatype = XSD_BOOLEAN
		}
		// 6)
	} else if isFloat {
		var tmpDatatype string
		if valueFloat != math.Trunc(valueFloat) || datatype.(string) == XSD_DOUBLE {
			value = strconv.FormatFloat(valueFloat, 'E', -1, 64)
			value = convertFloatValue(value.(string))
			tmpDatatype = XSD_DOUBLE
		} else {
			value = strconv.FormatFloat(valueFloat, 'f', 0, 64)
			tmpDatatype = XSD_INTEGER
		}
		if datatype == "" {
			datatype = tmpDatatype
		}
		// 7
	} else if isInt || datatype.(string) == XSD_INTEGER {
		value = strconv.FormatInt(valueInt, 10)
		if datatype == "" {
			datatype = XSD_INTEGER
		}
		// 8)
	} else {
		if datatype == "" {
			if hasLanguage {
				//TODO fix spec
				//datatype = RDF_LANGSTRING
				datatype = XSD_STRING
			} else {
				datatype = XSD_STRING
			}
		}
	}
	datatypeTerm := NewResource(datatype.(string))
	if hasLanguage {
		return NewLiteralWithLanguageAndDatatype(value.(string), language.(string),
			datatypeTerm.(Term))
	}
	return NewLiteralWithDatatype(value.(string), datatypeTerm.(Term))
}

func listToRDF(list []interface{}, idGenerator *BlankNodeIdGenerator) (Term, []*Triple) {
	listTriples := make([]*Triple, 0)
	// 1)
	if len(list) == 0 {
		return NewResource(RDF_NIL), listTriples
	}
	// 2)
	bnodes := make([]Term, 0)
	for i := 0; i < len(list); i++ {
		id := idGenerator.generateBlankNodeIdentifier(nil)
		bnodes = append(bnodes,
			NewBlankNode(id[2:]))
	}
	// 3)
	listTriples = make([]*Triple, 0)
	// 4)
	for i := 0; i < len(list); i++ {
		subject := bnodes[i]
		item := list[i]
		// 4.1)
		object := objectToRDF(item)
		// 4.2)
		if !isNil(object) {
			listTriples = append(listTriples,
				NewTriple(subject, NewResource(RDF_FIRST), object))
		}
		// 4.3)
		var rest Term
		if i == len(list)-1 {
			rest = NewResource(RDF_NIL)
		} else {
			rest = bnodes[i+1]
		}
		listTriples = append(listTriples,
			NewTriple(subject, NewResource(RDF_REST), rest))
	}
	// 5)
	return bnodes[0], listTriples
}

func fromRDF(dataset *Dataset, useNativeTypes bool,
	useRdfType bool) []interface{} {
	// 1)
	defaultGraph := make(map[string]interface{}, 0)
	// 2)
	graphMap := make(map[string]interface{})
	graphMap["@default"] = defaultGraph
	//TODO possible draft error
	nodeUsagesMap := make(map[string]interface{}, 0)
	// 3)
	for name, graph := range dataset.Graphs {
		// 3.2)
		if _, hasGraph := graphMap[name]; !hasGraph {
			graphMap[name] = make(map[string]interface{}, 0)
		}
		// 3.3)
		_, hasName := defaultGraph[name]
		if name != "@default" && !hasName {
			tmpMap := make(map[string]interface{}, 0)
			tmpMap["@id"] = name
			defaultGraph[name] = tmpMap

		}
		// 3.4)
		nodeMap := graphMap[name].(map[string]interface{})
		// 3.5)
		for _, triple := range graph {
			subject := triple.Subject.RawValue()
			predicate := triple.Predicate.RawValue()
			object := triple.Object.RawValue()
			// 3.5.1)
			if _, hasSubject := nodeMap[subject]; !hasSubject {
				tmpMap := make(map[string]interface{}, 0)
				tmpMap["@id"] = subject
				nodeMap[subject] = tmpMap
			}
			// 3.5.2)
			node := nodeMap[subject].(map[string]interface{})
			// 3.5.3)
			_, hasObject := nodeMap[object]
			if (isBlankNodeIdentifier(object) || isIRI(object)) && !hasObject {
				tmpMap := make(map[string]interface{}, 0)
				tmpMap["@id"] = object
				nodeMap[object] = tmpMap
			}
			// 3.5.4)
			if predicate == RDF_TYPE && !useRdfType &&
				(isBlankNodeIdentifier(object) || isIRI(object)) {
				mergeValue(node, "@type", object)
				continue
			}
			// 3.5.5)
			value := rdfToObject(triple.Object, useNativeTypes)
			// 3.5.6+7)
			mergeValue(node, predicate, value)
			// 3.5.8)
			if isIRI(object) || isBlankNodeIdentifier(object) {
				nodeObjectMap := nodeMap[object].(map[string]interface{})
				tmpMap := make(map[string]interface{}, 0)
				tmpMap["node"] = node
				tmpMap["property"] = predicate
				tmpMap["value"] = value
				mergeValue(nodeObjectMap, "usages", tmpMap)
				nodeMap[object] = nodeObjectMap
				//TODO spec is wrong
				mergeValue(nodeUsagesMap, object, name+" "+node["@id"].(string))
			}
		}
	}
	// 4)
	for name := range graphMap {
		graphObject := graphMap[name].(map[string]interface{})
		// 4.1)
		if _, hasNil := graphObject[RDF_NIL]; !hasNil {
			continue
		}
		// 4.2)
		//Spec defines this variable's name as nil, but nil is a reserved
		//keyword in go so I named it as nilValue instead
		nilValue := graphObject[RDF_NIL].(map[string]interface{})
		// 4.3)
		usages := nilValue["usages"].([]interface{})
		for index := range usages {
			usage := usages[index].(map[string]interface{})
			// 4.3.1)
			node := usage["node"].(map[string]interface{})
			property := usage["property"].(string)
			head := usage["value"].(map[string]interface{})
			// 4.3.2)
			list := make([]interface{}, 0)
			//TODO check type of listNodes
			listNodes := make([]interface{}, 0)
			// 4.3.3)
			for RDF_REST == property && isWellFormedListNode(node) &&
				len(nodeUsagesMap[node["@id"].(string)].([]interface{})) == 1 {
				// 4.3.3.1)
				list = append(list, node[RDF_FIRST].([]interface{})[0])
				// 4.3.3.2)
				listNodes = append(listNodes, node["@id"])
				// 4.3.3.3)
				nodeUsage := node["usages"].([]interface{})[0].(map[string]interface{})
				// 4.3.3.4)
				node = nodeUsage["node"].(map[string]interface{})
				property = nodeUsage["property"].(string)
				head = nodeUsage["value"].(map[string]interface{})
				// 4.3.3.5)
				if !isBlankNodeIdentifier(node["@id"].(string)) {
					break
				}
			}
			// 4.3.4)
			if property == RDF_FIRST {
				// 4.3.4.1)
				if RDF_NIL == node["@id"].(string) {
					continue
				}
				//4.3.4.3)
				headID := head["@id"].(string)
				// 4.3.4.4)
				head = graphObject[headID].(map[string]interface{})
				// 4.3.4.5)
				head = head[RDF_REST].([]interface{})[0].(map[string]interface{})
				// 4.3.4.6)
				list = list[:(len(list) - 1)]
				listNodes = listNodes[:(len(listNodes) - 1)]
			}
			// 4.3.5)
			delete(head, "@id")
			// 4.3.6)
			for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
				list[i], list[j] = list[j], list[i]
			}
			// 4.3.7)
			head["@list"] = list
			// 4.3.8)
			for _, nodeID := range listNodes {
				delete(graphObject, nodeID.(string))
			}
		}

	}
	// 5)
	result := make([]interface{}, 0)
	// 6)
	keys := sortedKeys(defaultGraph)
	for _, subject := range keys {
		node := defaultGraph[subject].(map[string]interface{})
		// 6.1)
		if _, hasSubject := graphMap[subject]; hasSubject {
			// 6.1.1)
			node["@graph"] = make([]interface{}, 0)
			// 6.1.2)
			keysGraph := sortedKeys(graphMap[subject].(map[string]interface{}))
			for _, s := range keysGraph {
				n := graphMap[subject].(map[string]interface{})[s]
				nMap := n.(map[string]interface{})
				_, hasID := nMap["@id"]
				delete(nMap, "usages")
				if len(nMap) == 1 && hasID {
					continue
				}
				node["@graph"] = append(node["@graph"].([]interface{}), nMap)
			}
		}
		// 6.2)
		delete(node, "usages")
		if _, hasID := node["@id"]; !(len(node) == 1 && hasID) {
			result = append(result, node)
		}
	}
	// 7)
	return result
}

func isWellFormedListNode(node interface{}) bool {
	nodeMap := node.(map[string]interface{})
	//TODO spec has no mention of @id
	for key := range nodeMap {
		if key != "@id" && key != "@type" && key != RDF_FIRST &&
			key != RDF_REST && key != "usages" {
			return false
		}
	}
	if !isBlankNodeIdentifier(nodeMap["@id"].(string)) {
		return false
	}
	if usages, hasUsages := nodeMap["usages"]; hasUsages {
		usagesArray, isArray := usages.([]interface{})
		if !(isArray && len(usagesArray) == 1) {
			return false
		}
	} else {
		return false
	}
	if first, hasKey := nodeMap[RDF_FIRST]; hasKey {
		firstArray, isArray := first.([]interface{})
		if !(isArray && len(firstArray) == 1) {
			return false
		}
	} else {
		return false
	}
	if rest, hasKey := nodeMap[RDF_REST]; hasKey {
		restArray, isArray := rest.([]interface{})
		if !(isArray && len(restArray) == 1) {
			return false
		}
	} else {
		return false
	}
	if typeValue, hasType := nodeMap["@type"]; hasType {
		typeArray, isArray := typeValue.([]interface{})
		if !(isArray && len(typeArray) == 1 && RDF_LIST == typeArray[0]) {
			return false
		}
	}
	return true
}

func rdfToObject(value Term, useNativeTypes bool) map[string]interface{} {
	// 1)
	if isTermResource(value) || isTermBlankNode(value) {
		returnValue := make(map[string]interface{}, 0)
		returnValue["@id"] = value.RawValue()
		return returnValue
	}
	//2)
	valueLiteral := value.(*Literal)
	// 2.1)
	result := make(map[string]interface{}, 0)
	// 2.2)
	var convertedValue interface{}
	convertedValue = valueLiteral.Value
	// 2.3)
	var typeValue interface{}
	typeValue = nil
	// 2.4)
	if useNativeTypes {
		// 2.4.1)
		if valueLiteral.Datatype.RawValue() == XSD_STRING {
			//TODO java version is different and does not add
			//string to result in this case
			//does nothing
			// 2.4.2)
		} else if valueLiteral.Datatype.RawValue() == XSD_BOOLEAN {
			if convertedValue == "true" {
				convertedValue = true
			} else if convertedValue == "false" {
				convertedValue = false
			}
			// 2.4.3)
		} else if valueLiteral.Datatype.RawValue() == XSD_DOUBLE ||
			valueLiteral.Datatype.RawValue() == XSD_INTEGER {
			floatValue, floatErr := strconv.ParseFloat(convertedValue.(string), 64)
			if isNil(floatErr) {
				convertedValue = floatValue
			}
		} else {
			typeValue = valueLiteral.Datatype.RawValue()
		}
		// 2.5)
	} else if valueLiteral.Language != "" {
		result["@language"] = valueLiteral.Language
		// 2.6)
	} else {
		if valueLiteral.Datatype.RawValue() != XSD_STRING {
			typeValue = valueLiteral.Datatype.RawValue()
		}
	}
	// 2.7)
	result["@value"] = convertedValue
	// 2.8)
	if !isNil(typeValue) {
		result["@type"] = typeValue
	}
	// 2.9)
	return result
}
