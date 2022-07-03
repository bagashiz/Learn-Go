package main

import "fmt"

func main() {
	//* if statement is used to execute a block of code if a condition is true

	//* if statement syntax:
	// if condition {
	// 	block of code
	// }
	//* if condition is true, the block of code is executed
	//* if condition is false, the block of code is not executed
	//* if condition is not true or false, the program will panic

	//* if else statement syntax:
	// if condition {
	// 	block of code
	// } else {
	// 	block of code
	// }

	//* if else if statement syntax:
	// if condition {
	// 	block of code
	// } else if condition {
	// 	block of code
	// } else {
	// 	block of code
	// }

	//* if statements in golang always require curly braces even if there is only one line of code
	if true {
		fmt.Println("This is true")
	}

	//* if statement can also be used to check if a key exists in a map using the ok variable (works with switch statement too)
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}

	//* Initializer syntax:
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	} else {
		fmt.Println("Florida is not a state")
	} // prints 20612439 because Florida is a key in the map
	//! fmt.Println(pop)
	//! variable pop out of scope because it is initiated in the if statement

	//* switch statement is used to execute a block of code based on a condition

	//* switch statement syntax:
	// switch condition {
	// 	case value1:
	// 		block of code
	// 	case value2:
	// 		block of code
	// 	default:
	// 		block of code
	// }

	//* tag less switch statement syntax:
	// switch condition {
	// 	case value1:
	// 		block of code
	// 	case value2:
	// 		block of code
	// 	default:
	// 		block of code
	// }

	//* initializer syntax:
	switch state := "Florida"; state {
	case "California":
		fmt.Println("California")
	case "Texas":
		fmt.Println("Texas")
	case "Florida":
		fmt.Println("Florida")
	default:
		fmt.Println("Unknown state")
	} // prints Florida

	//* switch statement with multiple cases
	switch state := "Florida"; state {
	case "California", "Texas": // comma-separated list of cases
		fmt.Println("West coast")
	case "Florida":
		fmt.Println("East coast")
	default:
		fmt.Println("Unknown state")
	} // prints East coast

	//* Tag less switch statement
	var i = 10
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
	case i > 10:
		fmt.Println("i is greater than 10")
	default:
		fmt.Println("i is equal to 10")
	} // prints i is equal to 10

	//* fallthrough keyword is used to execute the next case in the switch statement
	switch {
	case i <= 10:
		fmt.Println("i is less than 10")
		fallthrough // fallthrough keyword is optional
	case i <= 20:
		fmt.Println("i is less than 20")
	default:
		fmt.Println("i is greater than 20")
	} // prints i is less than 10 and i is less than 20

	//* break keyword is used to break out of a switch statement
	switch {
	case i <= 10:
		fmt.Println("i is less than 10")
		break // break keyword is optional
		//! fmt.Println("This doesn't print")
	case i <= 20:
		fmt.Println("i is less than 20")
	default:
		fmt.Println("i is greater than 20")
	} // prints i is less than 10

	//* Typed switch statement
	var x interface{} = "hello"
	// x is an interface type, it can be any type of data type (string, int, float, etc)
	switch x.(type) { // type assertion is used to cast an interface type to a concrete type (string, int, float, etc)
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is of unknown type")
	} // prints x is a string
}
