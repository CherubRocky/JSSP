package jssp

import "testing"

func TestMakespan(t *testing.T) {
	j := make([]Operation, 9)
	copy(j, []Operation{Operation{ID: 1, Machine: 1, Job: 1, NextMachine: 4,}, Operation{ID: 2, Machine: 2, Job: 1, NextMachine: 6,}, Operation{ID: 3, Machine: 3, Job: 1, NextMachine: 9,}, Operation{ID: 4, Machine: 1, Job: 2, NextMachine: 8,}, Operation{ID: 5, Machine: 3, Job: 2, NextMachine: 3,}, Operation{ID: 6, Machine: 2, Job: 2, NextMachine: -1,}, Operation{ID: 7, Machine: 2, Job: 3, NextMachine: 2,}, Operation{ID: 8, Machine: 1, Job: 3, NextMachine: -1,}, Operation{ID: 9, Machine: 3, Job: 3, NextMachine: -1,}})
	g := JSSPDigraph(j)
	want := 16
	got := g.Makespan()
	gotWant(t, got, want)
}

func gotWant(t testing.TB, got, want int) {
	t.Helper()
	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}
}
