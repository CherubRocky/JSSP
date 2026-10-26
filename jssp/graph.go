package jssp

type DiGraph struct {
	VertexList []Vertex
}

type Vertex struct {
	InDegree int
	Neighbours []Vertex
}

func NewDependencyGraph(sourceNeigh, sinkNeigh Vertex) *DiGraph {
	dg := Digraph{VertexList: make([]Vertex, 2, 2),}
	source := Vertex{0, sourceNeigh}
	sink := Vertex{len(sinkNeigh), nil}
	for _, u := range sinkNeigh {
		u.Neighbours = append(u.Neighbours, sink)
	}
	dg.VertexList[0] = &source
	gd.VertexList[1] = &sink
	
}

func NewJSSPGraph() *DiGraph {
	
}
