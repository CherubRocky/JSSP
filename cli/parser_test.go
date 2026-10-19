package cli

import (
	"testing"
	"github.com/CherubRocky/JSSP/jssp"
)

func jobsList() []jssp.Job {
	jobs := make([]jssp.Job, 0, 3)
	job1 := make([]jssp.Operation, 0, 3)
	job1 = append(job1, jssp.Operation{1, 1, 3})
	job1 = append(job1, jssp.Operation{2, 1, 3})
	job1 = append(job1, jssp.Operation{3, 1, 3})
	job2 := make([]jssp.Operation, 0, 3)
	job2 = append(job2, jssp.Operation{1, 2, 2})
	job2 = append(job2, jssp.Operation{3, 2, 3})
	job2 = append(job2, jssp.Operation{2, 2, 4})
	job3 := make([]jssp.Operation, 0, 3)
	job3 = append(job3, jssp.Operation{2, 3, 3})
	job3 = append(job3, jssp.Operation{1, 3, 2})
	job3 = append(job3, jssp.Operation{3, 3, 1})
	jobs = append(jobs, job1)
	jobs = append(jobs, job2)
	jobs = append(jobs, job3)
	return jobs
}

func TestParser(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		want := jssp.Instance{3, 3, 10, 10, jobsList()}
		cad := `3 3 10 10
1 3 2 3 3 3
1 2 3 3 2 4
2 3 1 2 3 1`
		got, err := ParseJSSP(cad)
		assertNoError(t, err)
		if !got.Equals(want) {
			t.Errorf("got %q want %q", got.String(), want.String())
		}
	})
	t.Run("header with non-integers", func(t *testing.T) {
		cad := `100 100 12momos1 56381
no sé qué siga`
		_, err := ParseJSSP(cad)
		assertError(t, err, ErrWrongData)
	})
	t.Run("jobs declared don't match with jobs given 1", func(t *testing.T) {
		cad := `2 3 54 12
1 3 4 5`
		_, err := ParseJSSP(cad)
		assertError(t, err, ErrJobsMismatch)
	})
	t.Run("jobs declared don't match with jobs given 2", func(t *testing.T) {
		cad := `2 3 54 12
1 3 4 5
4 5 6 7
8 9 10 11`
		_, err := ParseJSSP(cad)
		assertError(t, err, ErrJobsMismatch)
	})
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Got an error: %v", err)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("did not got an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
