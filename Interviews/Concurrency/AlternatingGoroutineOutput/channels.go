package main

import (
	"fmt"
	"sync"
)

func WithChannels() {

	chanChar := make(chan bool)
	chanInt := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(2)

	go func(){
		wg.Done()
		for i:= 'A'; i<='E'; i++ {
			<-chanChar
			fmt.Print(string(i))
			chanInt <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i:=1; i<=5; i++ {
			<-chanInt
			fmt.Print(i)

			if i < 5 { // IF if condition not implemented the service will create a lock in hopes that the communication sent in the channel will be read somewhere
				chanChar <- true
			}
			
		}
	}()

	chanChar <- true

	wg.Wait()

}
