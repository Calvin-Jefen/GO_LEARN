package main

func main() {
	// cards := newDeck()
	// // cards = append(cards, "Six of Spades")

	// // cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// // fmt.Println("-----------")
	// remainingCards.print()

	// cards := newDeck()
	// // fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newcard() string {
// 	return "Five of Diamonds"
// }
