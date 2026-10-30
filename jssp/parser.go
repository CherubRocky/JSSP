package jssp

import (
	"errors"
	"strconv"
	"strings"
)

var ErrWrongData = errors.New("Found something that is not an int in the file")
var ErrWrongHeaderFormat = errors.New("The number of arguments in header is different from two")
var ErrJobsMismatch = errors.New("The number of jobs declared don't match the number of jobs given (represented by lines)")
var ErrNoContent = errors.New("The text is basically empty")

func ParseJSSP(text string) (Instance, error) {
	ins := Instance{}
	text = strings.TrimSpace(text)
	rawLines := strings.Split(text, "\n")
	var lines []string
	for _, l := range rawLines {
		l = strings.TrimSpace(l)
		if l != "" {
			lines = append(lines, l)
		}
	}
	if len(lines) == 0 {
		return ins, ErrNoContent
	}
	fLine := strings.Fields(lines[0])
	if len(fLine) != 2 {
		return ins, ErrWrongHeaderFormat
	}
	var err error
	
	ins.N, err = strconv.Atoi(fLine[0])
	if err != nil {
		return ins, ErrWrongData
	}
	ins.M, err = strconv.Atoi(fLine[1])
	if err != nil {
		return ins, ErrWrongData
	}

	if ins.N != len(lines) - 1 {
		return ins, ErrJobsMismatch
	}

	jobsSlice := make([]Operation, 0, ins.N * ins.M)
	counter := 0
	for i := 1; i < len(lines); i++ {
		iLine := strings.Fields(lines[i])
		for j := 0; j < len(iLine); j += 2 {
			mac, err := strconv.Atoi(iLine[j])
			if err != nil {
				return ins, ErrWrongData
			}
			sp, err := strconv.Atoi(iLine[j + 1])
			if err != nil {
				return ins, ErrWrongData
			}
			op := Operation{ID: counter, Machine: mac, Job: i - 1, Span: sp}
			jobsSlice = append(jobsSlice, op)
			counter++
		}
	}
	ins.Jobs = jobsSlice
	return ins, nil
}
