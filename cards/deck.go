package main

import "fmt"

// Create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
}

// The actual copy of the deck we're working with
// is available in the function as variable "d"

// Every variable of type "deck" can call this function on itself
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
