package main

import "fmt"

func main() {
	// 1. create a new variable and assign to a playing card
	// var cards string = "Ace of Spades"
	cards := []string{"Ace of Diamonds", newCard()}

	// append to slice
	cards = append(cards, "Six of Spades")

	// iterate
	for index, card := range cards {
		fmt.Println(index, card)
	}

	// 2. print cards
	fmt.Println(cards)
}

// be explicit regardless return or props
func newCard() string {
	return "Five of Diamonds"
}
