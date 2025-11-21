package main

import (
	"fmt"
	"sync"
)


func WithMutexes() {
	var intPrint bool
	var mx sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	var iterations int

	go func(){
		defer wg.Done()
		for i:= 'A'; i<='E'; {
			iterations++
			mx.Lock()
				if !intPrint {	
					fmt.Print(string(i))
					intPrint = true
					i++
				}
			mx.Unlock()
		}
	}()

	go func(){
		defer wg.Done()
		for i:=1; i <=5; {
			iterations++
			mx.Lock()
				if intPrint {
					fmt.Print(i)
					intPrint = false
					i++
				} 
			mx.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println()
	fmt.Println(iterations, "--iterations")


	// more than 5000 iteration for a scenario where n = 10 
	// this condition is called BUSY WAITING 
	// Busy waiting (also called spin waiting or a spinlock) is when a goroutine or thread keeps actively checking for a condition to become true — instead of sleeping or blocking — and in the process, uses CPU cycles doing nothing productive.
}
