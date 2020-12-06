package racer

import (
	"net/http"
	"time"
)

// Racer 返回两个url中返回快的那个
func Racer(a, b string) string {
	if measureDuration(a) > measureDuration(b) {
		return b
	}
	return a
}

func measureDuration(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	duration := time.Since(startTime)
	return duration
}
