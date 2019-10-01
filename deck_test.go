package main

import "testing"

func TestNewDeck(t *testing.T)  {

	d := NewDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16 cards but got %d cards", len(d))
	}

	if d[0] != "One Of Diamonds" {
		t.Errorf("Expected first card One Of Diamonds but got %s card", d[0])
	}


}

func TestNewDeckByReadingFile(t *testing.T)  {
	newDeckByReadingFile()
}