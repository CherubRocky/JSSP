package cli

import (
	"errors"
	"strconv"
	"strings"
	"github.com/CherubRocky/JSSP/jssp"
)

var ErrWrongData = errors.New("Found something that is not an int in the file")
var ErrWrongHeaderFormat = errors.New("The number of arguments in header is different from four")
var ErrJobsMismatch = errors.New("The number of jobs declared don't match the number of jobs given (represented by lines)")
var ErrNoContent = errors.New("The text is basically empty")

func ParseJSSP(text string) (jssp.Instance, error) {
	ins := jssp.Instance{}
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
	if len(fLine) != 4 {
		return ins, ErrWrongHeaderFormat
	}
	var err error
	
	ins.Jobs, err = strconv.Atoi(fLine[0])
	if err != nil {
		return ins, ErrWrongData
	}
	ins.Machines, err = strconv.Atoi(fLine[1])
	if err != nil {
		return ins, ErrWrongData
	}
	ins.Ul, err = strconv.Atoi(fLine[2])
	if err != nil {
		return ins, ErrWrongData
	}
	ins.Ll, err = strconv.Atoi(fLine[3])
	if err != nil {
		return ins, ErrWrongData
	}

	if ins.Jobs != len(lines) - 1 {
		return ins, ErrJobsMismatch
	}

	jobsArr := make([]jssp.Job, 0, ins.Jobs)
	for i := 1; i < len(lines); i++ {
		iLine := strings.Fields(lines[i])
		var iJob jssp.Job = make([]jssp.Operation, 0, 2 * ins.Jobs)
		for j := 0; j < len(iLine); j += 2 {
			mac, err := strconv.Atoi(iLine[j])
			if err != nil {
				return ins, ErrWrongData
			}
			sp, err := strconv.Atoi(iLine[j + 1])
			if err != nil {
				return ins, ErrWrongData
			}
			op := jssp.Operation{mac, i, sp}
			iJob = append(iJob, op)
		}
		jobsArr = append(jobsArr, iJob)
	}
	ins.Jobss = jobsArr
	return ins, nil
}
