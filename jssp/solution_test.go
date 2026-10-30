package jssp

import "testing"

func TestInputsAsIs(t *testing.T) {
	t.Run("input 40", func(t *testing.T) {
		ins, _ := NewInstance(input40x40(), nil)
		sol := setup(ins)
		got := sol.Cost()
		want := 2.1666
		if sol.cost != 2.1666 {
			t.Errorf("got %f, want %f", got, want)
		}
		if ins.UB != 205 {
			t.Errorf("got %d, want %d", ins.UB, 205)
		}
		if ins.LB != 103 {
			t.Errorf("got %d, want %d", ins.LB, 103)
		}
		ms, _ := sol.digraph.Makespan()
		if ms != 324 {
			t.Errorf("got %d, want %d", ms, 324)
		}
	})
	t.Run("input 3 two", func(t *testing.T) {
		ins, _ := NewInstance(input3x3Two(), nil)
		sol := setup(ins)
		got := sol.Cost()
		want := 2.1666
		if sol.cost != 2.1666 {
			t.Errorf("got %f, want %f", got, want)
		}
		if ins.UB != 20 {
			t.Errorf("got %d, want %d", ins.UB, 20)
		}
		if ins.LB != 8 {
			t.Errorf("got %d, want %d", ins.LB, 8)
		}
		ms, _ := sol.digraph.Makespan()
		if ms != 324 {
			t.Errorf("got %d, want %d", ms, 324)
		}
	})
	
}

func setup(ins *Instance) *Solution {
	x := make([]Operation, len(ins.Jobs))
	copy(x, ins.Jobs)
	jobSlice := pointerSlice(x)
	ins.connectJob(jobSlice)
	connectMachAsIs(ins, jobSlice)
	return NewSolution(ins, jobSlice, nil)
}

func connectMachAsIs(ins *Instance, op []*Operation) {
	for i, _ := range ins.Machines {
		id := ins.Machines[i].ID
		if i % ins.OpsXMach != ins.OpsXMach - 1 {
			op[id].NextMachine = ins.Machines[i + 1].ID
			op[ins.Machines[i + 1].ID].InDegree++ 
		} else {
			op[id].NextMachine = -1
		}
	}
}
