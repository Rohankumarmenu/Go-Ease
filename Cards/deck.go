package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

// ["red", "blue", "green"] slice of strings convert
// "red,blue,green"  to a single string
// cards.toString()
// func toByteSlice(){

func (d deck) saveTofile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// }
