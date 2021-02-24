package main

import "fmt"

// To whom it may concern
// Our program has a new type called "bot"
// If you are a type in this program with a function named "getGreeting"
// and you return a string, you are now an honorary member of type "bot"
type bot interface {
	getGreeting() string // function set
}

type englishBot struct{} // honorary member of type "bot"
type spanishBot struct{} // honorary member of type "bot"

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// now that you are an honorary member of type "bot",
// you can now call this function called "printGreeting"
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// custom logic for english
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// custom logic for spanish
	return "Hola"
}

// $ go run main.go
