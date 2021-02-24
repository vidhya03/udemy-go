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

	salt.print()
	salt.updateName("Pepper")
	salt.print()

	name := "Bill"
	fmt.Println(*&name)

	mySlice := []string{"1", "2"}
	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "3"
}
func mapExample() {

}
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}
