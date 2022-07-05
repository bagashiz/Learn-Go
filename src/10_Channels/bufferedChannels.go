package main

import (
	"fmt"
	"sync"
)

//* Use buffer to getting around from a case which a goroutine sends 2 values to the channel,
//* but the receiver is not ready to receive the 2nd value.

var wg = sync.WaitGroup{}

func main() {
	//* Add a second parameter to the channel to specify the buffer size.
	ch := make(chan int, 50) // a variable that is a channel of ints with  buffer size of 50 ints.

	wg.Add(2)

	//* This goroutine will receive a value from the channel.
	go func(ch <-chan int) {
		for { // receives data from the channel using for loop.
			if i, ok := <-ch; ok { // checks if the channel is closed.
				fmt.Println("Received:", i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	//* This goroutine will send 2 values to the channel.
	go func(ch chan<- int) {
		ch <- 69  // send data to the channel.
		ch <- 420 // send data to the channel.
		close(ch) // close the channel to avoid deadlock.
		wg.Done()
	}(ch)

	wg.Wait()
}
