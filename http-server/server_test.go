package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	store := StubStore{
		map[string]int{
			"Sheldon": 100,
			"Penny":   90,
		},
	}
	server := PlayerServer{&store}
	t.Run("get penny score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Penny", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "90"
		if response.Code != http.StatusOK {
			t.Error("http response status not ok")
		}
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})
	t.Run("get sheldon score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Sheldon", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "100"
		if response.Code != http.StatusOK {
			t.Error("http response status not ok")
		}
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("return 404 when player does not exist.", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Lenord", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Code
		want := http.StatusNotFound
		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}

	})
}

func TestAddScoreThenRetrive(t *testing.T) {
	store := StubStore{map[string]int{}}
	server := PlayerServer{&store}
	request, _ := http.NewRequest(http.MethodPost, "/players/Lenord", nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)
	server.ServeHTTP(response, request)
	server.ServeHTTP(response, request)
	request, _ = http.NewRequest(http.MethodGet, "/players/Lenord", nil)
	response = httptest.NewRecorder()
	server.ServeHTTP(response, request)
	got := response.Body.String()
	want := "3"
	if got != want {
		t.Errorf("expected %v but got %v", want, got)
	}

}
