package main

import "fmt"

func main() {
	cards := deck{newCard(), "Four of Spades"}
	cards = append(cards, "vidhya is a good boy")

	cards.print()

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"

}
