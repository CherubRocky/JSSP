package jssp

import "fmt"

type Solution struct {
	ins *Instance
	digraph *DiGraph
	cost float64
	machines []*Operation
}

func NewSolution(i *Instance, jobs, machs []*Operation) *Solution {
	sol := &Solution{ins: i, machines: machs,}
	sol.digraph = &DiGraph{jobs}
	sol.Cost()
	return sol
}

func (s *Solution) Cost() float64 {
	if s.cost > 0 {
		return s.cost 
	}
	makespan, vic := s.digraph.Makespan()
	ms := float64(makespan)
	lb := float64(s.ins.LB)
	ub := float64(s.ins.UB)
	inCicle := float64(vic)

	c := (ms - lb) / (ub - lb) + lb * inCicle
	s.cost = c

	fmt.Printf("ms: %f, lb: %f, ub: %f, inCicle: %f\n", ms, lb, ub, inCicle)
	return c
}

func (s *Solution) Neighbour() *Solution {
	return nil
}
