package main

func main() {
	cards := newDeck()
	// cards = append(cards, "Six of Spades")

	// cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	// fmt.Println("-----------")
	remainingCards.print()
}

// func newcard() string {
// 	return "Five of Diamonds"
// }
