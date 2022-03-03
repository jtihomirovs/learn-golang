package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// With "type deck []string" we create a new custom type of 'deck', which is a slice of strings
// So our deck essentially behaves like a slice of string but we now have the ability to tie some additional
// functions specifically to anything that is of type deck.

// This type deck we are creating extends all the behaviour of slice of string
// So we can say or guarantee that type deck === []string
type deck []string

func newDeck() deck {
	// Create a new variable called cards of type deck
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Any instance of the class deck has the print method
// Before the print method we add a reciever - (d deck). Any variable inside our app with type deck now gets access to the print method.
// Receiver sets up methods on the variables we create
// d - instance of variable we are working with. In this example d is reference to cards array.
// deck - reference to the type that we want attach this method to
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// We can call this function with first argument of type deck, and the second argument of type int
// Signature is: deal(var type, var type)
// Adding receiver underlies on whether or not we are modifying that underlying array (slice) of type deck
// So here we are creating a new slice of type deck so we don't need a receiver
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// (d deck) is a receiver and later we can use cards.toString()
func (d deck) toString() string {
	// type conversion from deck to string:
	// []byte('Hi there')
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// We don't need any receiver to the function, because we don't have a deck. We would call this thing because we want access to the deck.
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	// If there is an error
	if err != nil {
		// What can we do if an error accurs
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the error nd entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
