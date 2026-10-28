package jssp

import "testing"

func TestMakespan(t *testing.T) {
	j := newTestOperations3x3()
	g := JSSPDigraph(j)
	want := 12
	makespan, vic := g.Makespan()
	gotWant(t, makespan, want, vic)
}

func gotWant(t testing.TB, got, want, vic int) {
	t.Helper()
	if want != got {
		t.Errorf("got %d, want %d, vic %d", got, want, vic)
	}
	if vic > 0 {
		t.Errorf("vic %d", vic)
	}
}

func newTestOperations3x3() []*Operation {
	return []*Operation {
		// Job 1
		{ID: 1, Machine: 1, Job: 1, NextMachine: 4, NextJob: 2, InDegree: 0, Span: 3,},
		{ID: 2, Machine: 2, Job: 1, NextMachine: 6, NextJob: 3, InDegree: 2, Span: 3,},
		{ID: 3, Machine: 3, Job: 1, NextMachine: 9, NextJob: -1, InDegree: 2, Span: 3,},
		// Job 2
		{ID: 4, Machine: 1, Job: 2, NextMachine: 8, NextJob: 5, InDegree: 1, Span: 2,},
		{ID: 5, Machine: 3, Job: 2, NextMachine: 3, NextJob: 6, InDegree: 1, Span: 3,},
		{ID: 6, Machine: 2, Job: 2, NextMachine: -1, NextJob: -1, InDegree: 2, Span: 4,},
		// Job 3
		{ID: 7, Machine: 2, Job: 3, NextMachine: 2, NextJob: 8, InDegree: 0, Span: 3,},
		{ID: 8, Machine: 1, Job: 3, NextMachine: -1, NextJob: 9, InDegree: 2, Span: 2,},
		{ID: 9, Machine: 3, Job: 3, NextMachine: -1, NextJob: -1, InDegree: 2, Span: 1,},
	}
	
}
