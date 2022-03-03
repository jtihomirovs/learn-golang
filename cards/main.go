package main

func main() {
	// cards := newDeck()
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// // We take this cards variable and this cards variable is passed to print function as the variable d
	//
	//cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()
}
