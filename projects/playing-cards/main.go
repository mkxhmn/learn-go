package main

func main() {
	// 1. create a new variable and assign to a playing card
	// var cards string = "Ace of Spades"
	cards := newDeck()

	// append to slice
	cards = append(cards, "Six of Spades")

	// iterate
	cards.print()
}
