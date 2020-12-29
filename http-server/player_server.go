package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{&StubStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Could not start player server at port 5000, error: %v", err)
	}
}
