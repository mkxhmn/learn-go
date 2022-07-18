package main

import "fmt"

func main() {
	fmt.Println("hello")

	// 1. create a new variable and assign to a playing card
	// var cards string = "Ace of Spades"
	cards := newCard()

	// 2. print cards
	fmt.Println(cards)
}

// be explicit regardless return or props
func newCard() string {
	return "Five of Diamonds"
}
