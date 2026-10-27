package jssp

type DiGraph struct {
	Jobs []Operation
	Machine []Operation
}

func JSSPDigraph(op []Operation) *DiGraph {
	return nil
}

func (dg *DiGraph) Makespan() int {
	return 0
}
