package main

import "fmt"

func main() {
	// cards := newDeck()
	// // cards = append(cards, "Six of Spades")

	// // cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// // fmt.Println("-----------")
	// remainingCards.print()

	cards := newDeck()
	fmt.Println(cards.toString())
}

// func newcard() string {
// 	return "Five of Diamonds"
// }
