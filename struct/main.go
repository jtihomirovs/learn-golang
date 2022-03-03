package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type contactInfo2 struct {
	address string
}

type person struct {
	firstName string
	lastName  string
	// Approach1
	contact contactInfo
	// Approach 2
	contactInfo2
}

func main() {
	// Approach 1
	//alex := person{"Alex", "Anderson"}

	// Approach 2
	alex_2 := person{firstName: "Alex", lastName: "Anderson"}

	// Declare a new struct in Go
	var bill person

	//fmt.Println(alex)

	fmt.Println(alex_2)

	fmt.Printf("%+v", bill)

	fmt.Println()

	//fmt.Println(alex.firstName)

	// Update values
	//alex.firstName = "Alex 2"
	//alex.lastName = "Alex 2"

	//fmt.Println(alex)

	// Use contactinfo
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@party.com",
			zipCode: 9400,
		},
		contactInfo2: contactInfo2{
			address: "street 1",
		},
	}
	// Old code:
	// We turn &jim into a memory address or a pointer and then we assign that value to jimPointer
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	// jim.print()

	// Old code:
	// fmt.Println(jim)
	// jim.updateName("Jimmy")
	// jim.print()
	jim.updateName("jimmy")
	jim.print()
}

// Old code:
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// *person - a pointer that points at a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// Take this memory address *pointerToPerson and turn it into an actual value.
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
