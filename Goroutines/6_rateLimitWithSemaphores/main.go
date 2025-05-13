// CREATE A SECTION WHERE 1-10 gets printed in an interval of 3 seconds like
// 1-10 in an instant
// 2 second sleep
// 11-20 in an instant
// 2 second sleep
// 21-30 in an instant
// 2 second sleep

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main(){

// 	const total int = 10
// 	const buffer int = 2

// 	var wg sync.WaitGroup

// 	for i:=0; i<total; i+=buffer {

// 		for j:=i; j< i+buffer; j++ {

// 			wg.Add(1)

// 			go func(wg *sync.WaitGroup, num int){
// 				defer wg.Done()

// 				time.Sleep(3 * time.Second)

// 				fmt.Println("done --", num)
// 				}(&wg, j)

// 			}

// 		fmt.Println("--")

// 		wg.Wait()

// 	}
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
		
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	for w := 0; w < 2; w++ {
		go func(id int) {

			fmt.Println("HERE activates the worker-", id)

			for job := range jobs {// The range jobs loop automatically performs the job := <-jobs operation.
				fmt.Printf("Worker %d processing job %d\n", id, job)
				time.Sleep(3 * time.Second)
				wg.Done()
			}
		}(w)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		jobs <- i
	}

	wg.Wait()
	close(jobs)

}

// EXPLANATION

// Creating the Channel:
// jobs := make(chan int, 10)
// This creates a buffered channel jobs that can hold up to 10 integer values. The channel is used to send tasks (jobs) to the workers.


// Setting Up Workers:
// for w := 0; w < 2; w++ {
//     go func(id int) {
//     }(w)
// }
// This launches 2 goroutines (workers). Each worker will execute the anonymous function concurrently. The id parameter is used to differentiate between the two workers.


// Worker Loop (Receiving Jobs):
// for job := range jobs {
//     fmt.Printf("Worker %d processing job %d\n", id, job)
//     time.Sleep(3 * time.Second)
//     wg.Done()
// }
// The range jobs loop will automatically receive values from the jobs channel. Once a worker receives a job, it processes it (prints a message and sleeps for 3 seconds to simulate work). When the worker finishes the job, wg.Done() is called to signal that the worker is done with this task.


// Sending Jobs to Workers:
// for i := 0; i < 10; i++ {
//     wg.Add(1)
//     jobs <- i
// }
// This loop adds 10 jobs to the jobs channel. For each job, wg.Add(1) increases the sync.WaitGroup counter to indicate that a worker will process one more job. The job itself (the integer i) is sent into the channel.


// Waiting for All Jobs to Be Processed:
// wg.Wait()
// This makes the main goroutine wait until all the workers are done with their tasks (until the sync.WaitGroup counter goes back to 0).


// Closing the Channel:
// close(jobs)
// After all jobs have been sent, the close(jobs) is called. This signals to the workers that no more jobs will be sent. The workers will stop processing when they encounter the closed channel.


// How the Range Loop Works:
// The range jobs loop on the workers will automatically stop when the channel is closed and empty. It’s a convenient way to process all items in a channel until it's closed and all data has been received.

// Key Points:
// The sync.WaitGroup (wg) ensures that the main goroutine waits for all workers to finish processing their jobs.

// The workers pick up tasks from the jobs channel and process them one by one.

// The close(jobs) signals to the workers that no more jobs will be added to the channel, allowing them to terminate once the channel is empty.

// The wg.Wait() ensures the program waits until all tasks are complete before exiting.

// Your understanding is correct! You've grasped how channels and goroutines work together to execute tasks concurrently.