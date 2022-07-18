package main

import "fmt"

// Create a new type of 'deck', which is a slice of strings

type deck []string
type suites []string
type values []string

func newDeck() deck {
	cards := deck{}
	suites := suites{"Spades", "Diamond", "Heart", "Clubs"}
	values := values{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, suit := range suites {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
