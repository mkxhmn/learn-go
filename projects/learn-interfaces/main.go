package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)
}

func printGreetings(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// acknowledge the error for now
func printGreetings(sb englishBot) {
	fmt.Println(sb.getGreeting())
}

// this also works, just include the struct/ "empty function"
func (englishBot) getGreeting() string {
	return "hello there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
