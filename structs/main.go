package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
type contactInfo struct {
	email   string
	zipcode int
}

func main() {

	salt := person{
		firstName: "Salt",
		lastName:  "Theni",
		contact: contactInfo{
			email:   "test@va.com",
			zipcode: 625531},
	}

	fmt.Printf("%+v", salt)

}
