package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{firstName: "alex", lastName: "aderson"}
	fmt.Println(alex.firstName)

	var alex2 person
	alex2.firstName = "alex2"
	alex2.lastName = "aderson2"

	fmt.Println(alex2)
	fmt.Printf("%+v", alex2)

	jim := person{
		firstName: "jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	fmt.Println(jim)

	// pass by value example wth classic type
	// jimPointer.updateName("jimmy")
	// jim.print()

	// pass by pointer with classic way
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()

	// pass by pointer with shortcut way
	jim.updateName("jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
