package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func NewDeck() deck {
	cards := deck{
	}
	suits := []string{"Diamonds", "Spades", "Clubs", "Hearts"}
	values := []string{"Five", "Two", "Three", "Four"}

	for _, s := range suits {
		for _, v := range values {
			c := v + " Of " + s
			cards = append(cards, c)
		}
	}

	return cards
}
func newCard() (d string) {

	d = "Ace of Diamonds"

	return
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, ":", card)
	}

}

func (d deck) cut(n int) (deck, deck) {
	return d[1:n], d[n:14]
}

func (d deck) combineCards() string {
	return strings.Join(d, ",")
}

func (d deck) writeIntoFile(filename string) {
	ioutil.WriteFile(filename, []byte(d.combineCards()), 0666)
}

func newDeckByReadingFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}
