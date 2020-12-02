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

func assertValue(t *testing.T, dict Dict, key string, want string) {
	t.Helper()
	got, err := dict.Search(key)
	assertError(t, err, nil)
	assertString(t, got, want)
}

func TestAdd(t *testing.T) {
	t.Run("add new key", func(t *testing.T) {
		dict := make(Dict)
		err := dict.Add("test", "answer")
		assertError(t, err, nil)
		assertValue(t, dict, "test", "answer")
	})
	t.Run("add existing key", func(t *testing.T) {
		dict := Dict{"test": "test"}
		err := dict.Add("test", "answer")
		assertError(t, err, ErrorKeyExist)
		assertValue(t, dict, "test", "test")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("normal case", func(t *testing.T) {
		dict := Dict{"test": "test"}
		err := dict.Update("test", "update")
		assertError(t, err, nil)
		assertValue(t, dict, "test", "update")
	})

	t.Run("key not exist case", func(t *testing.T) {
		dict := make(Dict)
		err := dict.Update("test", "update")
		assertError(t, err, ErrorKeyNotExist)
	})
}
