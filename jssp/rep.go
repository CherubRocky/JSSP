package jssp

import (
	"fmt"
)

type Instance struct {
	Jobs int
	Machines int
	Ul int
	Ll int
	Jobss []Job
}

type Operation struct {
	Machine int
	Job int
	Span int
}

type Job []Operation

func (j Job) Equals(cmp Job) bool {
	if len(j) != len(cmp) {
		return false
	}
	for i, op := range j {
		if op != cmp[i] {
			return false
		}
	}
	return true
}

func (i Instance) Equals(cmp Instance) bool {
	if i.Ul == cmp.Ul && i.Ll == cmp.Ll && i.Jobs == cmp.Jobs && i.Machines == cmp.Machines {
		if len(i.Jobss) != len(cmp.Jobss) {
			return false
		}
		for i, jobcy := range i.Jobss {
			if !jobcy.Equals(cmp.Jobss[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (i Instance) String() string {
	return fmt.Sprintf("Upper Limit: %d, Lower Limit: %d, Jobs: %d, Machines: %d", i.Ul, i.Ll, i.Jobs, i.Machines)
}
