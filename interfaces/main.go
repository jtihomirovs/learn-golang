package main

import "fmt"

//
type bot interface {
	// anything that kind of matches this description right here Is now of type bot
	// and so we can then use those other types in any other location where we expect to see a type of bot
	getGreeting() string
}

// English bot and Spanish bot are now also of type bot
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Remove eb if we would not use them in function
// By default it would be:
// func (eb englishBot) getGreeting() string {
func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi there!"
}

// The same for the spanish bot - remove sb
func (spanishBot) getGreeting() string {
	// Very custom logic for generatin an spanish greeting
	return "Hola!"
}
