package ConcurrencyPatterns

import (
	"fmt"
	"sync"
)

func WorkerPool() {
	
	jobChan := make(chan int)
	
	var wg sync.WaitGroup 
	
	const workerCount = 3
	wg.Add(workerCount)
	worker := func(workerID int, jobs <- chan int){
		defer wg.Done()
		for i := range jobs {
			fmt.Printf("Worker %d Started working on %d \n", workerID, i)
			// time.Sleep(500 * time.Millisecond) // Simulate work
		}
	}

	for i:= 1; i<= workerCount; i++ {
		go worker(i, jobChan)
	}

	for i:=0; i < 100; i++ {
		jobChan <- i
	}

	close(jobChan)
	wg.Wait()
}
