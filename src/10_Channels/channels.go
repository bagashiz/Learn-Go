package main

import (
	"fmt"
	"sync"
)

//* Channels are a go's way for a goroutine to communicate with other goroutines.
//* Channels are a type of data structure that can be used to communicate with other goroutines.
//* Channels use <- syntax to send and receive values.

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // a variable that is a channel of ints.

	//* Create 5 sets of goroutines and each one of them is going to have a sender and a receiver.
	//* And all of them is going to be using a single channel to communicate.
	for j := 0; j < 5; j++ {
		wg.Add(2)

		//* This goroutine will receive a value from the channel.
		go func(ch <-chan int) {
			i := <-ch // variable i receives data from the channel.
			fmt.Println("Received:", i)
			wg.Done()
		}(ch)

		//* This goroutine will send a value to the channel.
		go func(ch chan<- int) {
			ch <- 69 // send data to the channel.
			wg.Done()
		}(ch)
	}

	wg.Wait()
}
