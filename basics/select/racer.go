package racer

import (
	"net/http"
)

// Racer 返回两个url中返回快的那个
func Racer(a, b string) string {
	// select 会同时监听两个channel
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
