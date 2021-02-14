package main

import (
	"os"
	"testing"
)

// this is for testing
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expecting 16 but got %v ", len(d))
	}

	if d[len(d)-1] != "four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadDeck := newDeckFromFile("_decktesting")
	if len(loadDeck) != 16 {
		t.Errorf("Expecting 16 but got %v ", len(deck))
	}
	os.Remove("_decktesting")

}
