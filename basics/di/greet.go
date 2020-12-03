package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet will output a say hello string into the first argument(who implements io.Writer.Write())
func Greet(writer io.Writer, name string) {
	fmt.Fprint(writer, "Hello, "+name)
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Greet(w, "http server")
	}))
}
