package main

import "fmt"

//* Interfaces are a way to define a contract
//* Interfaces are used to define the behavior of a type

//* Syntax:
// type interface_name interface {
//		method_name(parameters) return_type
// }

func main() {
	//* Example:
	//* Create a variable that is of type Writer and set it to a ConsoleWriter
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!")) // call the Write method on the variable

	//* Example2:
	//* Create a variable that is of type IntCounter and set it to 0
	myInt := IntCounter(0)       // myInt is a variable of type IntCounter and it's value is 0
	var inc Incrementer = &myInt // inc is a variable of type Incrementer and it's value is the address of myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment()) // call the Increment method on the variable
	}
}

//* Example:
//* Writer is an interface that has a write method
type Writer interface {
	Write([]byte) (int, error) // Write method takes a slice of bytes and returns an int and an error
}

//* Implement the Writer interface to create a type that can write to a file
type ConsoleWriter struct{} // create a type that implements the Writer interface

//* Implement the Write method to write to the console
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data)) // convert the slice of bytes to a string and print it
	return n, err
}

//* Example2:
//* Incrementer is an interface that has a increment method
type Incrementer interface {
	Increment() int // Increment method returns an int
}

//* Implement the Incrementer method to create a type that can increment a number
type IntCounter int // create a type that implements the Incrementer interface

//* Implement the Increment method to increment the value of the variable
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

//* Compose the Incrementer and Writer interfaces to create a type that can write to a file and increment a number
type FileWriter struct { // FileWriter is a composed type that implements the Incrementer and Writer interfaces
	Writer      // Writer is a type that implements the Writer interface
	Incrementer // Incrementer is a type that implements the Incrementer interface
}

//* Empty interface is a way to define a contract that doesn't specify any methods
type Empty interface{}

//* Interface Best Practices:
//* 1. Use many, small interfaces to define many, small types
//* 2. Don't export interfaces for types that will be consumed
//* 3. Do export interfaces for types that will be used by package
//* 4. Design functions and methods to receive interfaces whenever possible
