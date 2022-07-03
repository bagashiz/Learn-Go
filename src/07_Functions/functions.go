package main

import "fmt"

//* Function is a block of code that takes input and returns output
//* Functions are called by their name and arguments
//* Functions can be called from anywhere in the program
//* Functions can be called from multiple places in the program
//* Functions can be called from multiple functions in the program
//* Functions can be called from the main function
//* Functions can be called from the same function

//* Function declaration syntax
// func functionName(parameterName type) returnType {
// 		function body
//		 return expression
// }
//* without return type
// func functionName(parameterName type) {
// 		function body
// }

//* Function call syntax
// functionName(argumentName)

//* Function naming conventions in go
//* PascalCase: public from the package
//* camelCase: private to the package

//* main() function is the entry point of the program
//* main() function is the only function that can be called from the main package

func main() {
	sayMessage("Hello World")                // call function sayMessage that prints "Hello World"
	sayMessageWithIndex("Hello World 2", 10) // call function sayMessageWithIndex that prints "Hello World 2" and prints "The value of the index is 10"

	//* Passing by value
	//* When a function is called, the values of the arguments are copied to the function's local variables
	//* When the function returns, the values of the local variables are copied to the caller's arguments
	greeting := "Assalamu alaikum"
	name := "Abdul"
	sayMessage2(greeting, name) // prints "Assalamu alaikum Abdul" because the values of the arguments are copied to the local variables
	fmt.Println(name)           // prints "Abdul" because the value of the local variable is copied to the argument

	//* Passing by reference
	//* When a function is called, the address of the arguments are copied to the function's local variables
	//* When the function returns, the address of the local variables are copied to the caller's arguments
	sayMessage3(&greeting, &name)
	fmt.Println(name) // prints "Mohammad" because the address of the local variable is copied to the caller's argument

	//* Calling function with variadic parameters
	sum(1, 2, 3, 4, 5)                 // prints "[1 2 3 4 5]" and prints "The sum is 15"
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) // prints "[1 2 3 4 5 6 7 8 9 10]" and prints "The sum is 55"

	//* Calling function with return value
	s := sum2(1, 2, 3, 4, 5)      // prints "[1 2 3 4 5]" but does not print the value of total variable
	fmt.Println("The sum is ", s) // prints "The sum is 15"

	//* Calling function with return value and pointer
	s2 := sum3(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", *s2) // / *s2 because the return type is a pointer to an int so we need to dereference the pointer to get the value

	//* Anonymous function
	//* Anonymous functions are functions that are defined without a name
	//* Anonymous functions are useful when we want to pass a function as an argument to another function
	//* Anonymous functions are executed immediately after the function that calls them is executed
	func() {
		msg := "Hello Go"
		fmt.Println(msg) // prints "Hello Go"
	}() // brackets are used to call the anonymous function

	//* Calling function with multiple return values with error checking
	d, err := divide(5, 0) // d is the first return value and err is the second return value
	if err != nil {
		fmt.Println(err) // prints "Cannot divide by zero"
		return
	} else {
		fmt.Println(d)
	}
}

//* function named sayMessage that takes a string as an argument
func sayMessage(msg string) { // msg parameter is scoped to the function sayMessage only
	fmt.Println(msg)
}

//* Function with multiple different parameters
func sayMessageWithIndex(msg string, idx int) {
	fmt.Printf("%v ", msg)
	fmt.Println("The value of the index is ", idx)
}

//* Function with multiple same parameters
func sayMessage2(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Mohammad" // this will change the value of the name parameter not the local variable
}

//* Function with pointers
func sayMessage3(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Mohammad" // this will change the value of the local variable because the address of the name parameter is copied to the local variable
}

//* Function with variadic parameters
func sum(numbers ...int) { // ... means that the function can take any number of arguments
	fmt.Println(numbers)
	total := 0
	for _, n := range numbers {
		total += n
	}
	fmt.Println("The sum is ", total)
}

//* Function with return value
func sum2(numbers ...int) int { // second int is the return type of the function
	fmt.Println(numbers)
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

//* Function that returns pointers
func sum3(numbers ...int) *int { // the return type is a pointer to an int
	fmt.Println(numbers)
	total := 0
	for _, n := range numbers {
		total += n
	}
	return &total // return the address of the total variable
}

//* Function with named return values
func sum4(numbers ...int) (total int) { // the return type is named total and the value of total is returned
	fmt.Println(numbers)
	for _, n := range numbers {
		total += n
	}
	return // do not need to return the value of total
}

//* Function with multiple return values
func divide(a, b int) (int, error) { // error is one of the return type of the function
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero") // fmt.Errorf is a function that returns an error value
	}
	return a / b, nil // return nil because there is no error
}
