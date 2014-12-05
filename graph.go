package gojsonld

type Graph struct {
	triples map[*Triple]bool

	uri  string
	term Term
}

func NewGraph(uri string) *Graph {
	if uri != "@default" && uri[:5] != "http:" &&
		uri[:6] != "https:" {
		panic(uri)
	}
	return &Graph{
		triples: make(map[*Triple]bool),

		uri:  uri,
		term: NewResource(uri),
	}
}

func (g *Graph) Len() int {
	return len(g.triples)
}

func (g *Graph) Term() Term {
	return g.term
}

func (g *Graph) URI() string {
	return g.uri
}

func (g *Graph) AddToDataset(d *Dataset) {
	if _, hasGraph := d.Graphs[g.URI()]; !hasGraph {
		d.Graphs[g.URI()] = make([]*Triple, 0)
	}
	for triple := range g.triples {
		d.Graphs[g.URI()] = append(d.Graphs[g.URI()], triple)
	}
}
