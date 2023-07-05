package main

func main() {
	filename := "my_cards"

	cards := newDeck()
	cards.shuffle()

	hand, _ := deal(cards, 5)
	hand.print()

	hand.saveToFile(filename)
	hand = newDeckFromFile(filename)

	hand.print()
}
