package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World" // okay so basically here the channel sleeps after holding the value assigned in buffer argument

	// form here you can only insert if you have extracted the previous inserted value form the channel

	fmt.Println(<-ch)
	
	// ch<-"Harsh" // will insert after the extraction over
	// ch<-"Here" // lock will be applied  until the channel is not full
	
	fmt.Println("Code Finished")
	fmt.Println(<-ch)

}