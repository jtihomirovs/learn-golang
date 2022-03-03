package main

import "fmt"

type ColorMap map[string]string

func main() {
	// 1 Approach

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// 2 Approach
	// We use VAR keyword If we want to kind of figure out later on what values or what key value pairs we want to add onto it
	var colors2 map[string]string

	// 3 Approach
	colors3 := make(map[string]string)

	// Add a new key-value pair
	colors3["white"] = "#ffffff"

	// Access key's value
	fmt.Println(colors3["white"])

	// Delete
	delete(colors3, "white")

	fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)

	printMap(colors)

	// Test stuff
	colors4 := ColorMap{
		"blue":  "#ff0000",
		"black": "#4bf745",
	}

	colors4.print2()

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func (c ColorMap) print2() {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
