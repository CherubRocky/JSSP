package cli

import (
	"testing"
	"errors"
	"os"
)

func TestMyReadFile(t *testing.T) {
	t.Run("Happy case", func(t *testing.T) {
		got, _ := MyReadFile("bait")
		want := "hola, babuino\n"
		assertMessage(t, got, want)
	})
	t.Run("File not found", func(t *testing.T) {
		_, got := MyReadFile("rutafalsaxd")
		assertError(t, got, os.ErrNotExist)
	})
	t.Run("No permission to open file", func(t *testing.T) {
		_, got := MyReadFile("no_accesible")
		assertError(t, got, os.ErrPermission)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Test failed. Got: %q, want: %q", got, want);
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if !errors.Is(got, want) {
		t.Fatalf("Got %q, want %q", got, want)
	}
}
