package main

import (
	"bytes"
	"testing"
)

func TestCoundDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	CountDown(buffer)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
