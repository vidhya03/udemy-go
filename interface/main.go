package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}
type deck string

func main() {

	eb := englishBot{}
	sb := spanishBot{}
	var s deck = "vde"

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(s)

}
func (eb englishBot) getGreeting() string {
	// very custom logic for generation an english greetings
	return "Hi how are you!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (ss deck) getGreeting() string {
	return "deck string"
}

func (eb spanishBot) getGreeting() string {
	// very custom logic for generation an english greetings
	return "hola como estas!"
}
