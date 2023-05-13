package cards

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spaces, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	test_filename := "_decktesting"

	os.Remove(test_filename)

	d := newDeck()
	d.saveToFile(test_filename)

	loadedDeck := newDeckFromFile(test_filename)

	if len(loadedDeck) != len(d) {
		t.Errorf("Expected %v cards in deck, but got %v", len(loadedDeck), len(d))
	}

	os.Remove(test_filename)
}
