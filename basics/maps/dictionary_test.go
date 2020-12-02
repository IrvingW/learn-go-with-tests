package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("normal case", func(t *testing.T) {
		dict := Dict{"test": "test"}
		got, err := dict.Search("test")
		assertError(t, err, nil)
		want := "test"
		assertString(t, got, want)
	})
	t.Run("unknown key case", func(t *testing.T) {
		dict := Dict{"test": "test"}
		got, err := dict.Search("unknown")
		assertError(t, err, ErrorNotFound)
		want := ""
		assertString(t, got, want)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %s but want error %s", got, want)
	}
}
