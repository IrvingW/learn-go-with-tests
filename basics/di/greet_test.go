package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "TestFunc")
	got := buffer.String()
	want := "Hello, TestFunc"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
