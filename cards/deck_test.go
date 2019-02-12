package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length 12, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expeted Ace of Spades but got %v", d[0])
	}
}
