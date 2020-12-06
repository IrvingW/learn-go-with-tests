package racer

import (
	"net/http"
	"time"
)

// HTTPError 自定义Error
type HTTPError string

func (e HTTPError) Error() string {
	return string(e)
}

const (
	// TimeoutError 当服务器超时时返回
	TimeoutError = HTTPError("This url has timeout.")
)

// Racer 返回两个url中返回快的那个
func Racer(a, b string, timeout time.Duration) (string, error) {
	// select 会同时监听两个channel
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", TimeoutError
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
