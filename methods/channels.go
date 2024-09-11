// Used for communication between goroutines. They allow goroutines to send and receive values of a specified type, enabling synchronization and data exchange
//ch := make(chan int) // Creates a channel of type int


package main

import (
    "fmt"
    "time"
)

func main() {

	bufferedChannel()
}

func unbuffered() {
    ch := make(chan int, 2) // Create a channel

    // Start a goroutine that sends data to the channel
    go func() {
		
        time.Sleep(time.Second) // Simulate work
        time.Sleep(time.Second) // Simulate work
        time.Sleep(time.Second) // Simulate work
        time.Sleep(time.Second) // Simulate work
        time.Sleep(time.Second) // Simulate work
        time.Sleep(time.Second) // Simulate work

        ch <- 1 // Send data to the channel // data is fead in the channel so it'll tell the groutine under to proceed // similar to a trigger.
		//  this will block until another goroutine receives the value
    }()

	fmt.Println(111111111111)


    data := <-ch // output 1 // activates/ runs when the date is fed into the channel
	// main goroutine will block here until a value is sent
    fmt.Println(data) // Output: hello from goroutine


    // Receive data from the channel
}

func bufferedChannel() {
    
	ch := make(chan string, 1) // Buffered channel with capacity of 2

    go func() {

        time.Sleep(1 * time.Second)

        ch <- "task 1 completed"

        fmt.Println("Sent task 1")
        
        ch <- "task 2 completed"

        fmt.Println("Sent task 2")
    }()

    time.Sleep(2 * time.Second)
    
	// since the task one is recived, a new value can be pushed into the channel of buffer 1  
    fmt.Println(<-ch) // Receive task 1 
	
    time.Sleep(2 * time.Second)

    fmt.Println(<-ch) // Receive task 2
	
	time.Sleep(10 * time.Second)
}