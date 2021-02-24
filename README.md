# Interfaces

What we know:

* Every value has a type
* Every function has to specify the type if it's arguments

So does that mean....

Every function that we ever write has to be rewritten to accomidate different types even if the logic in it is identical.

That is where interfaces come in. Below you can see some code that doesn't quite work out the way we want

```go
package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetingEnglish(eb)
	printGreetingSpanish(sb)
}

// printGreetingEnglish and printGreetingSpanish
// share the same logic but take a different type
// But how can we write one method that accepts two different types?
// ideally we would have one function called "printGreeting"
// that could be used for english bot and spanish bot types
func printGreetingEnglish(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingSpanish(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	// custom logic for english
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// custom logic for spanish
	return "Hola"
}

```

The solution is to use an interface:

```go

package main

import "fmt"

// To whom it may concern
// Our program has a new type called "bot"
// If you are a type in this program with a function named "getGreeting"
// and you return a string, you are now an honorary member of type "bot"
type bot interface {
	getGreeting() string
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

```

## Format


Below is the format for declaring an interface:

```

type <name> interface {
    <functon-name>(<list-of-argument-types>) (<list-of-return-types>)
}

```

Example:

```go

type user struct{
    name string
}

type bot interface {
    getGreeting(string, int) (string, error)
    getBotVersion() int
    respondToUser(user) string
}

```

## Concrete vs Interface types

**Concrete types** - you can create a value out of it directly and then access and change it.

* map
* int
* struct
* string
* englishBot
* spanishBot

**Interfact types** - you cannot create a value directly out of this type:

* bot