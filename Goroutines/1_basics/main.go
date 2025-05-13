package main

import (
	"fmt"
	"time"
)

// 1. GoRoutines are light weight threads managed by Go Runtime 2kb stack
// 2. Invoked with go keyword could be a function of IIFE go func(){}()

// Race Conditions: Two or more goroutines (or threads) access the same shared data at the same time, and at least one of them is writing to it.

func main(){ // main() runs on a special goroutine called the main goroutine.
	
	for i:=0; i<10; i++ {
		go concurrencyPrint(i)

		go func(i int){
			fmt.Println("Printing Integer --", i)
		}(i)

	}
	time.Sleep(3* time.Second)
	
}

func concurrencyPrint(i int){
	fmt.Println("Printing Integer --", i)
}



// Main exits before goroutines finish === No sync/wait
// Shared variable access === No locking — leads to race conditions
// Misuse of closures in loops === All goroutines capture the same variable