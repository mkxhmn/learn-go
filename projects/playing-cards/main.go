package main

func main() {
	// 1. create a new variable and assign to a playing card
	// var cards string = "Ace of Spades"
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

  hand.print()
  remainingCards.print()
}
