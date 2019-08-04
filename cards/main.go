package main

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())
	// hand, remainingCards := deal(cards, 2)

	// hand.print()
	// remainingCards.print()

}
