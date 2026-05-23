package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		//options: log the error and return a call to newDeck()
		//option 2: log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",") //Ace of Spades,Two of Spades,Three of Spades,Four of Spades,Ace of Diamonds,Two of Diamonds,Three of Diamonds,Four of Diamonds,Ace of Hearts,Two of Hearts,Three of Hearts,Four of Hearts,Ace of Clubs,Two of Clubs,Three of Clubs,Four of Clubs
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}

// }
