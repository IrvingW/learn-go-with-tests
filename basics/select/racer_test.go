package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeMockServer(t, 1*time.Second)
	fastServer := makeMockServer(t, 0)
	// 在某个函数调用前加上 defer 前缀会在 包含它的函数结束时 调用它
	defer slowServer.Close()
	defer fastServer.Close()
	slowURL := slowServer.URL
	fastURL := fastServer.URL
	winner := Racer(slowURL, fastURL)
	if winner != fastURL {
		t.Errorf("got %v but want %v", winner, fastURL)
	}
}

func makeMockServer(t *testing.T, delay time.Duration) *httptest.Server {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
	return server
}
