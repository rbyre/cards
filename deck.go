package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// (d deck) receiver of a function.. gir alle variabler av type deck tilgang til print() funksjonen.
// d er placeholder for variabelen som evt tar funksjonen i bruk? d pga d er første bokstav i deck.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
