package main

import (
    "fmt"
    "time"
	"context"
	"net/http"
)

func main() {
    // timeAfter();
	// contextWithTimeout()

	// networkTimeoutOperations();
}


func networkTimeoutOperations(){
	client := http.Client{  //Creates an instance of Go's http.Client, which is used to make HTTP requests.
        Timeout: 2 * time.Second,  // Set a 2-second timeout for the request
    }

    resp, err := client.Get("https://example.com"); //makes an http get request to the given URL
    if err != nil { // if err is not nil then it means some error has occured
        fmt.Println("Request failed:", err);
        return
    }
    defer resp.Body.Close() // defer resp.Body.Close() call schedules the body to be closed after the function finishes executing. This means the Close() function won't be executed immediately when it's called, but rather when the current function returns.

	// The body will still be accessable in the function but it will release the resource the Body is occupying in the memory when i networkTimeoutOperations function exits?

    fmt.Println("Request succeeded:", resp.Status, resp.Body);
}

func contextWithTimeout(){
	// context package provides more advanced control over timeouts and cancellations, especially when working with network requests or background tasks.

	// The context automatically cancels itself after the specified duration or if the parent context is canceled. This is particularly useful for controlling the execution time of long-running or potentially blocking operations like HTTP requests, database queries, or other external services.

	// When a context is canceled (either due to the timeout being reached or manual cancellation), all operations relying on that context should stop or handle the cancellation appropriately. This enables the program to avoid waiting indefinitely for slow or non-responsive operations.

	// USE CASE
	// to ensure that a task doesn’t run indefinitely, especially when dealing with external resources like databases, APIs, or file systems that could take longer than expected or become unresponsive.

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    ch := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch <- "Task completed"
    }()

    select {
    case result := <-ch:
        fmt.Println(result)
    case <-ctx.Done():
        fmt.Println("Timeout:", ctx.Err())  // Will output context deadline exceeded
    }

}

func timeAfter(){
	// time.After waits for a specific amount of time before it starts executing the default code //when it takes too long to execute a task
	
	ch := make(chan string) //here we are creating a chanel 

    go func() { //creating groutine to similate parallel work

        // Simulate a task
        time.Sleep(5 * time.Second)
        ch <- "Task completed"
    }()



	
	// HERE IS THE EXAMPLE OF NON BLOCKING CHANNELS.
	
	//the select case will automatically take the default value if it does not recive the channels response the the activity reached cases section.

	// Here’s a non-blocking receive. If a value is available on messages then select will take the <-messages case with that value. If not it will immediately take the default case

    select { // here a case when our code is executed before time.After(2 * time.Second) then our code runs else our default section executes when it took more than 2 secs and we want to concide the code that it took more than estimated time to run.
    case result := <-ch:
        fmt.Println(result, "Result")
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout: Task took too long")
    }

}


