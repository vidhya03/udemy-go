package main

import "testing"

// this is for testing
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expecting 16 but got %v ", len(d))
	}
}
