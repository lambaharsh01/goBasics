// used to manage and distribute tasks across a fixed number of worker goroutines, enabling efficient parallel processing of tasks.
//It helps limit the number of goroutines running simultaneously, preventing resource exhaustion

// 1. Workers (Goroutines): Goroutines that perform the actual work
// 2. Task Queue (Channel): A buffered or unbuffered Go channel that holds tasks 
// 3. Dispatcher (Optional): Sometimes, a separate dispatcher goroutine is used to send tasks into the task queue, which workers then pick up.

package main

import (
	"fmt"
	"time"
)

// func worker(id int, jobs <-chan int, results chan<- int): This is the worker function. It takes:

// id: The worker's ID (so you know which worker is doing the job).
// jobs <-chan int: A channel that the worker listens to for jobs.
// results chan<- int: A channel where the worker will send the result after finishing the job.
// for job := range jobs: This is an infinite loop that keeps checking the jobs channel for new work (tasks). If the jobs channel has something, it pulls a job (a number, in this case) from it.

// It's like each worker is sitting around waiting for a job to arrive. Once a job is put into the jobs channel, a worker grabs it and starts working.
func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {
		// fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate processing time
		// fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 5)   // Task queue
	results := make(chan int, 5) // Result queue

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// go worker(w, jobs, results): This is where the "magic" happens. We are calling the worker function, but by using the go keyword, each call runs in the background as a goroutine (a lightweight thread in Go).
	// This means we are creating three separate workers that will run in the background, each waiting to receive tasks from the jobs channel.
	// Each worker gets its own unique ID (w), which is passed to the worker function (so you can know which worker is working on which task).
	// The workers are like employees sitting around waiting for work (which will come from the jobs channel).
	// So, after this loop finishes, you have three background workers ready to take on jobs but doing nothing until work is added to the jobs channel.










	// Send 5 jobs to the workers
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	// The loop runs 5 times because j starts at 1 and increases up to 5.
	// Sending jobs to the jobs channel:

	// Each time the loop runs, it sends a job (an integer) to the jobs channel using jobs <- j.
	// The workers are waiting for jobs to arrive in the jobs channel, so as soon as you send a job, one of the workers picks it up.
	// If all 3 workers are busy, the remaining jobs will wait in the channel (because it's buffered with a capacity of 5).


	// Collect results from the workers
	for r := 1; r <= 5; r++ {
		fmt.Printf("Result: %d\n", <-results)
	}


// WHY THE ERROR OCCURES WHEN I GIVE 100 JOBS TO THE CHANEL OF 5
// Channel Capacity:

// You initially created the jobs and results channels with a capacity of 5: jobs := make(chan int, 5) and results := make(chan int, 5).
// This means the channels can only hold 5 jobs or 5 results at a time.
// Sending 100 Jobs:

// When you increase the number of jobs to 100, you’re trying to send all 100 jobs into the jobs channel.
// However, because the channel can only hold 5 jobs, after the first 5 jobs are sent, it gets full, and the code gets stuck waiting for space to send the remaining 95 jobs.
// Deadlock:

// Since you're trying to send more jobs than the workers can process at a time, the program stops. The workers are either asleep (waiting for more jobs) or the main function is stuck because it can't send more jobs into the jobs channel.
// This creates a deadlock: neither the workers nor the main function can proceed, and Go gives the deadlock error.

	
}
