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
	sol.ComputeCost()
	return sol
}

func (s *Solution) ComputeCost() float64 {
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

func (s *Solution) Cost() float64 {
	if s.cost < 0 {
		return s.ComputeCost()
	}
	return s.cost
}

func (s *Solution) Neighbour() *Solution {
	neigh := s.clone()
	machSelected := s.ins.rng.IntN(s.ins.M)
	machOp1 := s.ins.rng.IntN(s.ins.OpsXMach)
	machOp2 := s.ins.rng.IntN(s.ins.OpsXMach)
	for machOp1 == machOp2 {
		machOp2 = s.ins.rng.IntN(s.ins.OpsXMach)
	}
	neigh.swapMachine(machSelected, machOp1, machOp2)
	neigh.ComputeCost()
	return neigh
}

func (s *Solution) swapMachine(mach, posA, posB int) {
	truePosA := mach * s.ins.OpsXMach + posA
	truePosB := mach * s.ins.OpsXMach + posB
	tmp := s.machines[truePosA]
	// put b in original position a
	s.machines[truePosA] = s.machines[truePosB]
	// put a in original position b
	s.machines[truePosB] = tmp

	// reconfigure
	s.reconfigure(truePosA)
	s.reconfigure(truePosB)
}

func (s *Solution) reconfigure(machPos int) {
	id := s.machines[machPos].ID
	s.machines[machPos].InDegree = 0
	s.machines[machPos].NextMachine = -1
	s.machines[machPos].NextJob = -1
	if s.HasJobPredecessor(id) {
		s.digraph.Ops[id - 1].NextJob = id
		s.machines[machPos].InDegree += 1
	}
	if s.HasMachinePredecessor(machPos) {
		s.machines[machPos - 1].NextMachine = id
		s.machines[machPos].InDegree += 1
	}
	if s.HasJobSuccessor(id) {
		s.digraph.Ops[id].NextJob = id + 1
	}
	if s.HasMachineSuccessor(machPos) {
		idSucc := s.machines[machPos + 1].ID
		s.machines[machPos].NextMachine = idSucc
	}
}

func (s *Solution) HasJobPredecessor(opID int) bool {
	if opID % s.ins.OpsXJob == 0 {
		return false
	}
	return true
}

func (s *Solution) HasMachinePredecessor(machPos int) bool {
	if machPos % s.ins.OpsXMach == 0 {
		return false
	}
	return true
}

func (s *Solution) HasJobSuccessor(opID int) bool {
	if (opID + 1) % s.ins.OpsXJob == 0 {
		return false
	}
	return true
}

func (s *Solution) HasMachineSuccessor(machPos int) bool {
	if (machPos + 1) % s.ins.OpsXMach == 0 {
		return false
	}
	return true
}

func (s *Solution) clone() *Solution {
	newDigraph := s.digraph.clone()
	newMachines := make([]*Operation, len(s.machines))
	for i, mOp := range s.machines {
		newMachines[i] = newDigraph.Ops[mOp.ID]
	}
	newSol := &Solution{
		ins: s.ins,
		digraph: newDigraph,
		cost: s.cost,
		machines: newMachines,
	}
	return newSol
}

