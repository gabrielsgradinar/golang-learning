package main

func main() {
	cards := newDeck()
	cards.saveToFile("cards.txt")

	new_deck := newDeckFromFile("card.txt")
	new_deck.print()

}
