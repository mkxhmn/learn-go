package main

func main() {
	// 1. create a new variable and assign to a playing card
	// var cards string = "Ace of Spades"
	cards := newDeck()
  cards.saveToFile("deck_of_cards")

	hand, remainingCards := deal(cards, 5)

  hand.print()
  remainingCards.print()
}
