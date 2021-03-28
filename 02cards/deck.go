package main

import "fmt"

type deck []string

// Creates new deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}
	return cards
}

// Prints cards in the deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Deal a hand
func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	remainingCards := d[handSize:]
	return hand, remainingCards
}
