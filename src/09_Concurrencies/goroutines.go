package main

import (
	"fmt"
	"time"
)

//* Goroutines are lightweight threads of execution that can be run in parallel.
//* Goroutines run in the same address space as the main program.
//* Goroutines are created by calling the "go" keyword before the func() function.

func main() {
	//* Call the sayHello function in a new goroutine.
	go sayHello() // this will create a new goroutine and run it in the background.
	//* The output does not print because the function is executed in a goroutine and the main function ended early.
	time.Sleep(100 * time.Millisecond) // sleep for 100 milliseconds.
	//* Give goroutines time to run.

	//* Execute anonymous function that prints a value from a variable in a goroutine.
	var msg string = "World"
	go func() {
		fmt.Println(msg) // prints "World"
	}() // this will create a new goroutine and run it in the background.
	time.Sleep(100 * time.Millisecond) // sleep for 100 milliseconds.

	//! The function above create a dependency between the variable in the main function and the variable in the goroutine.
	go func() {
		fmt.Println(msg) // it should have print "World"
	}() // this will create a new goroutine and run it in the background.
	msg = "Goodbye"                    // this will change the value of the variable in the main function.
	time.Sleep(100 * time.Millisecond) // sleep for 100 milliseconds.
	//! The goroutine will print "Goodbye" because the variable in the function that runs in goroutine is dependent of the variable in the main function.

	//* To mitigate this, use parameters in the function.
	msg = "Hello"
	go func(msg string) { // the parameter is a copy of the variable in the main function.
		fmt.Println(msg) // prints "Hello"
	}(msg) // this will create a new goroutine and run it in the background.
	msg = "Goodbye"                    // this will change the value of the variable in the main function.
	time.Sleep(100 * time.Millisecond) // sleep for 100 milliseconds.
	//* The goroutine will print "Hello" because the variable in the function that runs in goroutine is independent of the variable in the main function.
}

//* Function to be executed in a goroutine.
func sayHello() {
	fmt.Println("Hello") // prints "Hello"
}
