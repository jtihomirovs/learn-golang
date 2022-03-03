package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

// Remove eb if we would not use them in function
// By default it would be:
// func (eb englishBot) getGreeting() string {
func (englishBot) getGreeting() string {
	// Very custom logic for generatin an english greeting
	return "Hi there!"
}

// The same for the spanish bot - remove sb
func (spanishBot) getGreeting() string {
	// Very custom logic for generatin an spanish greeting
	return "Hola!"
}
