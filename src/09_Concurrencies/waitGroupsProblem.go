package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter int

func main() {
	for i := 0; i < 10; i++ {
		//* Add(n) is used to increment the wait group by n.
		wg.Add(2)      // increment the wait group by 2.
		go sayHello()  // this will create a new goroutine and run it in the background.
		go increment() // this will create a new goroutine and run it in the background.
	}
	//! The loop prints the number randomly because there is no synchronization between goroutines.

	//* Wait() blocks until all goroutines have finished.
	wg.Wait() // Wait for all goroutines to finish.
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter) // prints "Hello #(counter)"
	//* Done() is used to decrement the wait group.
	wg.Done() // decrement the wait group.
}

func increment() {
	counter++
	//* Done() is used to decrement the wait group.
	wg.Done() // decrement the wait group.
}
