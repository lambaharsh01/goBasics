// ✅ Goroutines → for running code concurrently.
// ✅ Mutexes / RWMutex → for managing shared memory safely.
// 🔄 Channels → for communicating between goroutines without sharing memory.

// MUTEXES are for protecting data when multiple goroutines are trying to access the same data (read or write).

// CHANNELS are for communicating data between goroutines. They don’t need a lock because they’re a safe way to pass data from one goroutine to another. They do the synchronization automatically

package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string) {
	time.Sleep(3 * time.Second) // a very time consuming concurrent activity which can take n(unknown) number of time
    ch <- "Hello from goroutine!" // Sending message to channel
}

func main() {

	fmt.Println("Program Started")

    unbufferedChannel := make(chan string) // Create a channel to pass strings
    
    go sendMessage(unbufferedChannel)

    msg := <-unbufferedChannel // makes the code synchronous , will only after the concurrent activity is completed // conceptually similar to await in javascript
    fmt.Println(msg) // Prints: Hello from goroutine!
}


// EFFICIENCY: this process can be me more efficient when a signal is sent instead of the data itself 

	// done := make(chan struct{})

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	done <- struct{}{} // Send empty signal
	// }()

	// <-done // Wait for signal