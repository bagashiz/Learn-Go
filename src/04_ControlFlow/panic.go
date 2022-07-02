package main

import "fmt"

func main() {
	//* panic keyword
	//* panic keyword is used to cause a function to panic.
	//* works like exception in other languanges
	//* syntax:
	// panic(value)
	//* panic is used to stop the execution of a function and cause a program to exit.
	
	//a, b := 1, 0
	//ans := a / b
	//fmt.Println(ans) //! panic: runtime error: integer divide by zero

	//* panic() function
	fmt.Println("Start")
	panic("Something went wrong!") // prints "Start" and then throw panic: "Something went wrong!"
	// fmt.Println("End") //! all lines after panic() become unreachable
}