package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}
func (eb englishBot) getGreeting() string {
	// very custom logic for generation an english greetings
	return "Hi how are you!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func (eb spanishBot) getGreeting() string {
	// very custom logic for generation an english greetings
	return "hola como estas!"
}
