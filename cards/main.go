package cards

func main() {
	cards := newDeck()
	// cards.saveToFile("cards.txt")

	// new_deck := newDeckFromFile("card.txt")
	cards.shuffle()
	cards.print()

}
