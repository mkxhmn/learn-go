package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	PERMISSION_READ_AND_WRITE := fs.FileMode(0666)
	return ioutil.WriteFile(filename, []byte(d.toString()), PERMISSION_READ_AND_WRITE)
}

func newDeckFromFile(filename string) deck {
	byteSlice, error := ioutil.ReadFile(filename)

	if error != nil {
		// Option #1 - log the error and return new call to newDeck()
		// [/] Option #2 - log the error and entirely quit the program

		fmt.Println("Error:", error)

		ERROR_CODE := 1
		os.Exit(ERROR_CODE)
	}

	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)

		// swap card by index
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
