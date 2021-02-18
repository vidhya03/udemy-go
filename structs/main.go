package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}
type contactInfo struct {
	email   string
	zipcode int
}

func main() {

	salt := person{
		firstName: "Salt",
		lastName:  "Theni",
		contactInfo: contactInfo{
			email:   "test@va.com",
			zipcode: 625531},
	}
	// saltPointer := &salt
	salt.print()
	salt.updateName("Pepper")
	salt.print()
	// saltPointer.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}
