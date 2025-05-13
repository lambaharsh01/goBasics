package main

import (
	"fmt"
	"sync"
)

// WaitGroup is used to synchronize the goroutine execution

func main(){
	
	var waitGroup sync.WaitGroup

	for i:=0; i<10; i++ {
		waitGroup.Add(1)
		go concurrencyPrint(&waitGroup, i)
	}
	
	waitGroup.Wait()
	
}

func concurrencyPrint(waitGroup *sync.WaitGroup, i int){
	defer waitGroup.Done()

	fmt.Println("Printing Integer --", i)
}




// Mistake	Why it Happens
// Main exits before goroutines finish	No sync/wait
// Shared variable access	No locking — leads to race conditions
// Misuse of closures in loops	All goroutines capture the same variable