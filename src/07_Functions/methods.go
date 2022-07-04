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
//* syntax
// func (receiver_variable_name) function_name(parameters) return_type {
//		function_body
// }
//* receiver_variable_name is the name of the variable that will be used to call the method

func (g greeter) greet() { // (g greeter) means that the receiver variable is g with the type greeter
	fmt.Println(g.greeting, g.name)
}

//* Create setter and getter methods for the greeter struct
func (g *greeter) setGreeting(greeting string) { // use pointer to change the value of the greeting variable
	g.greeting = greeting
}
func (g *greeter) setName(name string) { // use pointer to change the value of the name variable
	g.name = name
}
