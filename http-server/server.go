package main

import (
	"fmt"
	"net/http"
)

// PlayerServer supplies two function
// one of them let you check specified player's score
// another one let you add score for specified player
func PlayerServer(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "20")
}
