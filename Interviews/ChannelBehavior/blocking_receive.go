package main

import (
	"fmt"
	"time"
)

func BlockingReceive() {

	c := make(chan int)

	go func() {
		c <- 1
	}()

	go func(){
		fmt.Println(<-c)
		fmt.Println(<-c) // "BLOCKING RECEIVE". if the close(c) function is not called and the value is not sent into the channel the second time
		// Second receive blocks forever because the channel is open but no value is available.
		fmt.Println("Execution Finished") // THE EXECUTION NEVER FINISHES
	}()

	time.Sleep(3 * time.Second)

}
