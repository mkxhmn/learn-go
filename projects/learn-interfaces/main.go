package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)
}

func printGreetings(b bot) {
	fmt.Println(b.getGreeting())
}

// this also works, just include the struct/ "empty function"
func (englishBot) getGreeting() string {
	return "hello there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
