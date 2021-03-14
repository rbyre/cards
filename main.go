package main

func main() {
	// var cards av type deck har tilgang til print() metode.
	cards := deck{newCard(), newCard()}

	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Hearts"
}
