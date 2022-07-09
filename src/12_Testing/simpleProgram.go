package main

import (
	"errors"
	"log"
)

func main() {
	//* Call divide function
	result, err := divide(10, 0)

	//* Check if error is nil
	if err != nil {
		//* Print error message
		log.Println(err.Error())
		return // exit function
	}

	//* Print result
	log.Println("Result of division is ", result) // will not print if error is not nil
}

//* Create function to divide two numbers
func divide(x, y float32) (float32, error) {
	//* Check if y is zero
	if y == 0 {
		//* errors.New() creates a new error with the given error message
		return 0, errors.New("cannot divide by zero") // return error message
	}

	result := x / y
	return result, nil
}
