package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	firstElementFromDeck := d[0]
	if firstElementFromDeck != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", firstElementFromDeck)
	}

	lastElementFromDeck := d[len(d)-1]
	if lastElementFromDeck != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs but got %v", lastElementFromDeck)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deck_testing")

	d := newDeck()
	d.saveToFile("_deck_testing")

	loadedDeck := newDeckFromFile("_deck_testing")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}
}
