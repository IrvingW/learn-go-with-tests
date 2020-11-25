package main

import (
	"fmt"
)

const en = "ENGLISH"
const spanish = "SPANISH"
const enHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// Hello says hello to someone(name) in some language(language)
// This func is public
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := helloPrefix(language)
	return prefix + name
}

func helloPrefix(language string) (prefix string) {
	switch language {
	case en:
		prefix = enHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return
}

func main() {
	// just an example
	fmt.Println(Hello("Irving", "en"))
}
