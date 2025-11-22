package main

import (
	"fmt"
	"time"
)

func ClosingTheChannel() {


	c := make(chan int)

	go func() {
		c <- 1
		close(c)
	}()

	go func(){
		// if the channel recived the second boolean argument would be true because it executed after receiving 
		value1, ok1 := <-c
		fmt.Println(value1, ok1)
		// ok1 is true because you received a real value from an open channel.


		// If you receive from a channel after it has been closed, you get the zero value of the channel's data type, and the second boolean (ok) will be false indicating the channel is closed.
		value2, ok2 := <-c
		fmt.Println(value2, ok2)
		// No more values will ever come && Receive does NOT block

		fmt.Println("Execution Finished")
	}()

	time.Sleep(3 * time.Second)

}
