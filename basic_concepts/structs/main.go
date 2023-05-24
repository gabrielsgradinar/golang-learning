package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName, lastName string
	contactInfo
}

func main() {

	// julia := person{"Julia", "Sofreio"}
	// julia := person{firstName: "Julia", lastName: "Sofreio"}
	julia := person{
		firstName: "Julia",
		lastName:  "Sofreio",
		contactInfo: contactInfo{
			email:   "julia@gmail.com",
			zipCode: 12345,
		},
	}
	julia.updateName("Xulia")
	julia.print()

	tiago := struct { // anonymous struct
		firstName, lastName string
		age                 int
	}{
		firstName: "Tiago",
		lastName:  "Grandinar",
		age:       14,
	}

	fmt.Printf("%#v\n", tiago)

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
