package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	// WRITE_OP write operation
	writeOp = "write"
	// SLEEP_OP sleep operation
	sleepOp = "sleep"
)

type CountDownOpSpy struct {
	Calls []string
}

func (c *CountDownOpSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, writeOp)
	return
}

func (c *CountDownOpSpy) Sleep(d time.Duration) {
	c.Calls = append(c.Calls, sleepOp)
}

func TestCountDown(t *testing.T) {
	t.Run("print 3 to Go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		CountDown(buffer, &CountDownOpSpy{})
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}

	})

	t.Run("sleep before every print", func(t *testing.T) {
		spy := &CountDownOpSpy{}
		CountDown(spy, spy)
		want := []string{
			sleepOp,
			writeOp,
			sleepOp,
			writeOp,
			sleepOp,
			writeOp,
			sleepOp,
			writeOp,
		}
		if !reflect.DeepEqual(spy.Calls, want) {
			t.Errorf("got %v but want %v", spy.Calls, want)
		}
	})
}
