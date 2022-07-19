package main

func main() {
	// 1. create a new variable and assign to a playing card
	// var cards string = "Ace of Spades"
	cards := newDeckFromFile("deck_of_cards")
	cards.print()
}
