package racer

import (
	"net/http"
	"time"
)

// Racer 返回两个url中返回快的那个
func Racer(a, b string) string {
	startTime := time.Now()
	http.Get(a)
	durationA := time.Since(startTime)

	startTime = time.Now()
	http.Get(b)
	durationB := time.Since(startTime)

	if durationA > durationB {
		return b
	}
	return a
}
