package jssp

import (
	"github.com/gammazero/deque"
	"fmt"
)

type DiGraph struct {
	Ops []*Operation
	Machine []Operation
}

func JSSPDigraph(op []*Operation) *DiGraph {
	graph := &DiGraph{Ops: op,}
	return graph
}

func (dg *DiGraph) Makespan() (int, int) {
	q := &deque.Deque[*Operation]{}
	for _, u := range dg.Ops {
		if u.InDegree == 0 {
			q.PushBack(u)
			u.TEnd = u.Span
			fmt.Println(u.TEnd)
		}
	}
	dg.printG()
	for q.Len() != 0 {
		u := q.PopFront()
		fmt.Printf("ID: %d, ", u.ID)
		fmt.Printf("Span: %d, ", u.Span)
		if u.NextJob != -1 {
			dg.printG()
			v := dg.Ops[u.NextJob - 1]
			v.TEnd = max(v.TEnd, u.TEnd + v.Span)
			v.InDegree--
			if v.InDegree == 0 {
				q.PushBack(v)
			}
		}
		if u.NextMachine != -1 {
			v := dg.Ops[u.NextMachine - 1]
			v.TEnd = max(v.TEnd, u.TEnd + v.Span)
			v.InDegree--
			if v.InDegree == 0 {
				q.PushBack(v)
			}
		}
	}
	makespan := dg.FindMax()
	vic := dg.vInCicle()
	return makespan, vic
}

func (dg *DiGraph) FindMax() int {
	max := 0
	for _, v := range dg.Ops {
		if max < v.TEnd {
			max = v.TEnd
		}
	}
	return max
}

func (dg *DiGraph) vInCicle() int {
	vic := 0
	for _, v := range dg.Ops {
		if v.InDegree > 0 {
			vic++
		}
	}
	return vic
}

func (dg *DiGraph) printG() {
	for _, v := range dg.Ops {
		fmt.Printf("%d, ", v.TEnd)
	}
	fmt.Println()
}
