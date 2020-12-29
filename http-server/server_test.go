package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testServer(t *testing.T) {
	t.Run("get someone's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Penny", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})
}
