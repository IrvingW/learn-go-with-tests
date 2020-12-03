package main

import (
	"bytes"
	"testing"
	"time"
)

type SpySleeper struct {
	CallCount int
}

func (s *SpySleeper) Sleep(d time.Duration) {
	s.CallCount++
}

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	CountDown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
	if spySleeper.CallCount != 4 {
		t.Errorf("got %d sleep but want %d sleep", spySleeper.CallCount, 4)
	}
}
