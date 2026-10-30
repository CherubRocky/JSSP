package jssp

import (
	"math/rand/v2"
	"slices"
	"fmt"
)

type Instance struct {
	UB int // Upper bound
	LB int // Lower bound
	N int// number of jobs
	M int// number of machines
	OpsXJob int // number of operations per job
	OpsXMach int // number of operations per machine
	Jobs []Operation
	Machines []Operation
	rng *rand.Rand
}

func NewInstance(textInstance string, rng *rand.Rand) (*Instance, error) {
	ins, err := ParseJSSP(textInstance)
	if err != nil {
		return nil, fmt.Errorf("Error creating instance. %w", err)
	}
	ins.rng = rng
	ins.Machines = ins.opsByMachine()
	ins.UB = ins.UpperBound()
	ins.LB = ins.LowerBound()
	ins.OpsXJob = len(ins.Jobs) / ins.N
	ins.OpsXMach = len(ins.Machines) / ins.M
	return &ins, nil
}

func (ins *Instance) opsByMachine() []Operation {
	machs := slices.Clone(ins.Jobs)
	f := func(a, b Operation) int {
		if a.Machine < b.Machine {
			return -1
		}
		if a.Machine == b.Machine {
			return 0
		}
		return 1
	}
	slices.SortStableFunc(machs, f)
	return machs
}

func (ins *Instance) UpperBound() int {
	sum := 0
	for _, v := range ins.Jobs {
		sum += v.Span
	}
	return sum
}

func (ins *Instance) LowerBound() int {
	n := len(ins.Jobs) / ins.N
	m := len(ins.Machines) / ins.M
	return max(maxSlice(ins.Jobs, n), maxSlice(ins.Machines, m))
}

func maxSlice(s []Operation, n int) int {
	maxVal := 0
	partSum := 0
	for i, _ := range s {
		if i % n == 0 {
			maxVal = max(partSum, maxVal)
			partSum = 0
		}
		partSum += s[i].Span
	}
	maxVal = max(partSum, maxVal)
	return maxVal
}

func (ins *Instance) RandomSol() []Operation {
	sol := slices.Clone(ins.Jobs)
	ins.connectJob(sol)
	ins.connectMachine(sol)
	return sol
}

func (ins *Instance) connectJob(op []Operation) {
	opsJob := len(op) / ins.N
	for i, _ := range op {
		if i % opsJob != opsJob - 1 {
			op[i].NextJob = op[i].ID + 1
			op[i + 1].InDegree += 1 
		}
	}
}

func (ins *Instance) connectMachine(op []Operation) {
	opsMach := len(op) / ins.M
	for i := 0; i < ins.M; i++ {
		perm := ins.rng.Perm(opsMach)
		for j, v := range perm {
			if j < len(perm) - 1 {
				actualInd := (i * opsMach) + v
				nextActualInd := (i * opsMach) + perm[j + 1]
				op[actualInd].NextMachine = op[nextActualInd].ID
				op[nextActualInd].InDegree += 1
			}
		}
	}
}


type Operation struct {
	ID int
	Machine int
	Job int
	Span int
	NextMachine int // Id of next machine in jobflow
	NextJob int // Id of next operation in jobflow
	InDegree int
	TEnd int
}
