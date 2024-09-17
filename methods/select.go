// select is a control structure used to handle multiple channel operations simultaneously. It allows a goroutine to wait on multiple communication operations (channel sends or receives) and proceed with the first one that’s ready. This is very useful in concurrent programming when you're dealing with multiple channels.

package main;

import ("fmt"; "time")


func main(){

ch1 := make(chan string)
ch2 := make(chan string)

go func() {
	time.Sleep(3 * time.Second)
	ch1 <- "message from ch1"
}()

go func() {
	time.Sleep(1 * time.Second)
	ch2 <- "message from ch2"
}()

select {
case msg1 := <-ch1:
	fmt.Println("Received:", msg1)
case msg2 := <-ch2:
	fmt.Println("Received:", msg2)
}

// select waits for one of the two channels (ch1 or ch2) to become ready.
// It will print whichever message arrives first (from ch2 in this case, since it has a 1-second delay).




time.Sleep(10 * time.Second)

}