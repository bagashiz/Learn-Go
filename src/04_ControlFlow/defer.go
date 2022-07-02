package main

import "fmt"

func main() {
	//* defer keyword
	//* defer keyword is used to defer a function call until the surrounding function returns.
	//* syntax:
	// defer func(){function cal}()
	//* defer use LIFO order
	//* defer always run before panic

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	} // prints 10 to 0 because the defer statement is executed after the loop has finished
}
