package main

import (
	"fmt"
	"sync"
	"time"
)

func FanOut() {
	// Single Producer Multiple Consumers

	arr := make([]int, 10)
	for i := range arr {
		arr[i] = i + 1
	}

	ch := make(chan int)
	var mx sync.Mutex

	var wg sync.WaitGroup
	workerCount := 3
	wg.Add(workerCount)

	worker := func(wID int, jobs <-chan int) {
		defer wg.Done()
		for job := range jobs {
			fmt.Println("worker", wID, "started task", job)
			time.Sleep(time.Second * 2)

			mx.Lock()
			arr[job] *= arr[job]
			mx.Unlock()

		}
	}

	for i := 1; i <= workerCount; i++ {
		go worker(i, ch)
	}

	for i := range arr {
		ch <- i // PRODUCER
	}

	close(ch)
	wg.Wait()
}
