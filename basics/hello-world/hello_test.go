package main

import (
	"testing"
)

// TestHello test Hello function
func TestHello(t *testing.T) {
	assertMessageCorrect := func(t *testing.T, got string, want string) {
		// tell test kit this function is a helper func
		// the error report will report the real line where error happened
		t.Helper()
		if got != want {
			t.Errorf("got %q, but expected %q", got, want)
		}
	}
	t.Run("say hello to some people", func(t *testing.T) {
		got := Hello("Go", en)
		expected := "Hello, Go"
		assertMessageCorrect(t, got, expected)
	})

	t.Run("say hello world when input str is empty", func(t *testing.T) {
		got := Hello("", en)
		want := "Hello, World"
		assertMessageCorrect(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Go", spanish)
		want := "Hola, Go"
		assertMessageCorrect(t, got, want)
	})

	t.Run("say hello with default language(en)", func(t *testing.T) {
		got := Hello("Go", "Go")
		want := "Hello, Go"
		assertMessageCorrect(t, got, want)
	})
}
