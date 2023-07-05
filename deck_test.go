package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	{
		expectedValue := 52
		if len(d) != expectedValue {
			t.Errorf("Expected deck length of 20, but got %v", len(d))
		}
	}

	{
		expectedValue := "Ace of Spades"
		if d[0] != expectedValue {
			t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
		}
	}

	{
		expectedValue := "King of Clubs"
		if d[len(d)-1] != expectedValue {
			t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
		}
	}
}
