package main

import (
	"fmt"
	"sync"
)

//* Wait groups are used to wait for a group of goroutines to finish.
//* It synchronizes the execution of all goroutines in the group.
//* The main goroutine calls Wait() on the wait group.
//* The main goroutine will then wait for all of the goroutines to finish.
//* The main goroutine can then do something else.
//* The main goroutine can then call Wait() again to wait for the goroutines to finish.

//* Variable to hold the wait group.
var wg = sync.WaitGroup{}

func main() {
	var msg string = "Hello"
	//* Add(n) adds an number of  goroutine to the wait group.
	wg.Add(1) // add 1 goroutine to the wait group.

	//* Go routine that will be executed.
	go func() {
		fmt.Println(msg) // prints "Hello"
		//* Done() is used to decrement the wait group.
		wg.Done() // decrement the wait group.
	}() // this will create a new goroutine and run it in the background.

	//* Wait() blocks until all goroutines have finished.
	wg.Wait() // Wait for all goroutines to finish.
}
