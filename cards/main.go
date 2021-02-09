package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	cards.saveToFile("fileName.txt")
	s := newDeckFromFile("fileName.txt")
	// fmt.Println("The values are ",s)
	fmt.Println()
	s.print()

}
