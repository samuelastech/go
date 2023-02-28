package main

func main() {
	cards := newDeck()
	cards.saveToFile("test")
	newDeck := newDeckFromFile("test")
	newDeck.print()
}
