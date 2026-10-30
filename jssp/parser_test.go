package jssp

import (
	"testing"
)

func jobsList() []Operation {
	jobs := []Operation {
		{ID: 0, Machine: 0, Job: 0, Span: 3},
		{ID: 1, Machine: 1, Job: 0, Span: 3},
		{ID: 2, Machine: 2, Job: 0, Span: 3},
		{ID: 3, Machine: 0, Job: 1, Span: 2},
		{ID: 4, Machine: 2, Job: 1, Span: 3},
		{ID: 5, Machine: 1, Job: 1, Span: 4},
		{ID: 6, Machine: 1, Job: 2, Span: 3},
		{ID: 7, Machine: 0, Job: 2, Span: 2},
		{ID: 8, Machine: 2, Job: 2, Span: 1},
	}
	return jobs
}


func TestParser(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		want := Instance{N: 3, M: 3, Jobs: jobsList()}
		cad := `3 3
0 3 1 3 2 3
0 2 2 3 1 4
1 3 0 2 2 1`
		got, err := ParseJSSP(cad)
		assertNoError(t, err)
		if !instanceEquals(got, want) {
			t.Error("Values differ")
		}
	})
	t.Run("header with non-integers", func(t *testing.T) {
		cad := `100veintitrés 100
no sé qué siga`
		_, err := ParseJSSP(cad)
		assertError(t, err, ErrWrongData)
	})
	t.Run("jobs declared don't match with jobs given 1", func(t *testing.T) {
		cad := `2 3
1 3 4 5`
		_, err := ParseJSSP(cad)
		assertError(t, err, ErrJobsMismatch)
	})
	t.Run("jobs declared don't match with jobs given 2", func(t *testing.T) {
		cad := `2 3
1 3 4 5
4 5 6 7
8 9 10 11`
		_, err := ParseJSSP(cad)
		assertError(t, err, ErrJobsMismatch)
	})
	t.Run("Incorrect header", func(t *testing.T) {
		cad := `2 3 67 69
1 3 4 5
4 5 6 7
8 9 10 11`
		_, err := ParseJSSP(cad)
		assertError(t, err, ErrWrongHeaderFormat)
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

func instanceEquals(a, b Instance) bool {
	if a.N != b.N || a.M != b.M || len(a.Jobs) != len(b.Jobs) {
		return false
	}
	for i, v := range a.Jobs {
		if !operationEquals(v, b.Jobs[i]) {
			return false
		}
	} 
	return true
}

func operationEquals(a, b Operation) bool {
	if !(a.ID == b.ID && a.Job == b.Job && a.Machine == b.Machine && a.Span == b.Span) {
		return false
	}
	return true
}
