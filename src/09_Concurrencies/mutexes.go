package main

import (
	"fmt"
	"sync"
)

//* Mutexes are used to synchronize access to shared data.
//* It is a mutual exclusion lock.
//* One goroutine will acquire the lock and the other will wait.
//* If the mutex is locked, and another goroutine tries to acquire it, it will block until the mutex is unlocked.
//* The mutex will unlock when the goroutine that locked it is finished.
//* RWMutex means other goroutines can read the data, but only one goroutine can write at a time.

var wg = sync.WaitGroup{}
var counter int
var m = sync.RWMutex{} // variable to hold the RWMutex group.

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2) // add 2 goroutines to the wait group.
		m.RLock() // acquire a read lock.
		go sayHello()
		m.Lock() // acquire a write lock.
		go increment()
	}
	wg.Wait() // Wait for all goroutines to finish.
} //* The loop prints the number sequentially because mutex is used to synchronize access to the variable.

func sayHello() {
	fmt.Printf("Hello #%v\n", counter) // prints "Hello #(counter)"
	m.RUnlock()                        // release the read lock.
	wg.Done()                          // decrement the wait group.
}

func increment() {
	counter++
	m.Unlock() // release the write lock.
	wg.Done()  // decrement the wait group.
}
