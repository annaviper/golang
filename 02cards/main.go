package main

import "fmt"

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 4)
	// cards.print()
	fmt.Println(hand)

}

