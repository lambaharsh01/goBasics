// closing a channel is a way to signal that no more values will be sent on it. This is useful because receivers can then know when to stop waiting for values. Closing a channel is done with the close() function.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

    go func() {
        for i := 1; i <= 10; i++ {
            ch <- i
        }
		fmt.Println("Channel closed, no more values.");

		close(ch); // once a grouting is closed 
		
		// // no more values can be inseted inside of that if tried to be done it goes in the panic state (panic: send on closed channel);
        // for i := 1; i <= 10; i++ {
			//     ch <- i
			// }

		}();

		time.Sleep(3 * time.Second)
		
		
		
		for val := range ch { // this is called range over chanels
			//This range iterates over each element as it’s received from queue. Because we closed the channel above,
			
			fmt.Println(val) // receive values until channel is closed
		}
		
		// fmt.Println("Channel closed, no more values.")


}