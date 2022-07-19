package main

import "testing"

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
