package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Index int
	Value int
}

func WorkerPoolPipeline() {
	// Data flows through multiple sequential stages, where the output of one stage becomes the input of the next.

	// Producer → Workers → Collector

	// It’s a Fan-Out / Fan-In pattern using channels + WaitGroup to process tasks in parallel and collect results safely.

	arr := make([]int, 20)
	for i := range arr {
		arr[i] = i + 1
	}

	producer := make(chan int)
	consumer := make(chan Result)

	var wg sync.WaitGroup
	workerCount := 3
	wg.Add(workerCount)

	worker := func(_ int, pCh <-chan int, cCh chan<- Result) {
		defer wg.Done()

		for job := range pCh {
			// fmt.Println("worker", wID, "took task", job)
			time.Sleep(time.Second * 2)

			cCh <- Result{Index: job, Value: arr[job] * arr[job]}
		}
	}

	for i := 1; i <= workerCount; i++ {
		go worker(i, producer, consumer)
	}

	go func() {
		for i := range arr {
			producer <- i
		}
		close(producer)
	}()

	go func() {
		wg.Wait()
		close(consumer)
	}()

	for jobDone := range consumer {
		fmt.Println("Job Done at index", jobDone.Index, "result value is", jobDone.Value)
	}

}
