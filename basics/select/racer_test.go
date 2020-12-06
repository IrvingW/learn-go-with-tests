package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Second * 1)
			w.WriteHeader(http.StatusOK)
		}))
	fastServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
	slowURL := slowServer.URL
	fastURL := fastServer.URL
	winner := Racer(slowURL, fastURL)
	if winner != fastURL {
		t.Errorf("got %v but want %v", winner, fastURL)
	}
}
