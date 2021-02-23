package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"math/rand"
	"time"
)
// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print (){
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(deckOfCards deck, handSize int) (deck, deck) {
	return deckOfCards[:handSize], deckOfCards[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - logt the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	slice_strings := strings.Split(string(bs), ",")
	return deck(slice_strings)
}

func (deckOfCards deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range deckOfCards{
		newPosition := r.Intn(len(deckOfCards) - 1)

		deckOfCards[i], deckOfCards[newPosition] = deckOfCards[newPosition], deckOfCards[i]
	}
}