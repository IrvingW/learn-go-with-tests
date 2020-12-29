package main

import (
	"fmt"
	"net/http"
)

// PlayerStore store player's score
// and implements query func & update func
type PlayerStore interface {
	GetScore(name string) int
	AddScore(name string)
}

// PlayerServer call store's functions to
// server the http requests
type PlayerServer struct {
	store PlayerStore
}

// StubStore implement PlayerStore interface
type StubStore struct {
	scoreMap map[string]int
}

// GetScore returns the score value stored in storeMap
func (s *StubStore) GetScore(name string) int {
	return s.scoreMap[name]
}

// AddScore add a score for input player
func (s *StubStore) AddScore(name string) {
	_, ok := s.scoreMap[name]
	if !ok {
		s.scoreMap[name] = 1
	} else {
		s.scoreMap[name]++
	}
}

// ServeHTTP supplies two function
// one of them let you check specified player's score
// another one let you add score for specified player
func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	player := request.URL.Path[len("/players/"):]
	switch request.Method {
	case http.MethodGet:
		p.showScore(response, player)
	case http.MethodPost:
		p.processWin(response, player)
	default:
		response.WriteHeader(http.StatusNotFound)
	}

}

func (p *PlayerServer) showScore(response http.ResponseWriter, name string) {
	score := p.store.GetScore(name)
	if score == 0 {
		response.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(response, score)
}

func (p *PlayerServer) processWin(response http.ResponseWriter, name string) {
	p.store.AddScore(name)
	response.WriteHeader(http.StatusAccepted)
}
