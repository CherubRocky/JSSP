package jssp

import (
	"github.com/gammazero/deque"
	"fmt"
)

type DiGraph struct {
	Ops []*Operation
}

func JSSPDigraph(op []*Operation) *DiGraph {
	graph := &DiGraph{Ops: op,}
	return graph
}

func (dg *DiGraph) Makespan() (int, int) {
	inDegrees := dg.inDegreeSlice()
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
		if u.NextJob != -1 {
			dg.printG()
			v := dg.Ops[u.NextJob - 1]
			v.TEnd = max(v.TEnd, u.TEnd + v.Span)
			inDegrees[v.ID - 1]--
			if inDegrees[v.ID - 1] == 0 {
				q.PushBack(v)
			}
		}
		if u.NextMachine != -1 {
			v := dg.Ops[u.NextMachine - 1]
			v.TEnd = max(v.TEnd, u.TEnd + v.Span)
			inDegrees[v.ID - 1]--
			if inDegrees[v.ID - 1] == 0 {
				q.PushBack(v)
			}
		}
	}
	makespan := dg.FindMax()
	vic := vInCicle(inDegrees)
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

func (dg *DiGraph) inDegreeSlice() []int {
	s := make([]int, len(dg.Ops))
	for i, v := range dg.Ops {
		s[i] = v.InDegree
	}
	return s
}

func  vInCicle(inDegrees []int) int {
	vic := 0
	for _, v := range inDegrees {
		if v > 0 {
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
