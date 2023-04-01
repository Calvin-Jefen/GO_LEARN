package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	card := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Heart", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card = append(card, value+" of "+suit)

		}
	}
	return card
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
