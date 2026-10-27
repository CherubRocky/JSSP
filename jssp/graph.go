package jssp

type DiGraph struct {
	Ops []Operation
	Machine []Operation
}

func JSSPDigraph(op []Operation) *DiGraph {
	return nil
}

func (dg *DiGraph) Makespan() int {
	counter := 0
	q := make([]*Operation, len(dg.Ops))
	for _, u := range dg.Ops {
		if true {// Missing logic
			q[counter] = &u
			counter++
			// Tfin logic
		}
	}
	counter = 0
	for counter != -1 {
		counter--
		u := q[counter]
		return u.Machine
	}
	return 0
}
