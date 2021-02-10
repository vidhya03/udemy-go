package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDesk()

	if len(d) != 2 {
		t.Errorf("Expecting 20 but got ", len(d))
	}
}
