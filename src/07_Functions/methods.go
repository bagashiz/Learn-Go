package main

import "fmt"

//* Methods are functions that are associated with a type
//* Methods are called by using the dot operator (.)
//* Methods are called by using the receiver variable and the dot operator (.)

func main() {
	g := greeter{"Hello", "World"}
	g.greet() // prints "Hello World"
}

//* greeter is a struct that contains a greeting and a name
type greeter struct {
	greeting string
	name     string
}

//* Method that is associated with the greeter struct
func (g greeter) greet() { // (g greeter) means that the receiver variable is g
	fmt.Println(g.greeting, g.name)
}
