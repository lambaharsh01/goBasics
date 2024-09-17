package main;

import(
	"fmt"
	"context"
	"time"
)

func main(){
	// requestScopeValues();
	cancleLongRunningOpperations()
	// deadline()
}



// 1. Request-Scoped Values
// The context package allows you to pass values across API boundaries and between different functions in your Go programs. 
//This is particularly useful in web servers and APIs where you need to pass request-specific data (like user authentication details) through different layers of your application.

func requestScopeValues() {
	// In web and API development, especially in microservices or request-response cycles, you often need to share some data across different layers (e.g., middleware, handlers, services). Using a context with key-value pairs allows you to propagate this data in a structured way without passing too many parameters explicitly.
    ctx := context.WithValue(context.Background(), "key", "value")
    value := ctx.Value("key").(string)
    fmt.Println(value) // Output: value
}



// 2. Cancellation
// The context package provides a way to cancel operations that are no longer needed. 
// This is crucial for cleaning up resources and stopping operations when they are no longer relevant. 
// For instance, if a user cancels a request, you can use the context to halt ongoing operations and clean up resources.

func longRunningOperation(ctx context.Context) {
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Operation completed")
    case <-ctx.Done():
        fmt.Println("Operation cancelled")
    }
}

func cancleLongRunningOpperations() {
    ctx, cancel := context.WithCancel(context.Background()) // here the cancle is declared without any time limit which means whenever it'll be called the 
    go longRunningOperation(ctx)

    // Simulate some work and then cancel the operation
    time.Sleep(2 * time.Second)


	// The cancel function (spelled cancel not cancle in your code) is a function returned by context.WithCancel in Go. It's used to signal the cancellation of an operation or task that's tied to the context created by context.WithCancel


    cancel();// Yes, calling cancel() will signal only the goroutines that are using the context to stop their work. It will not stop the main thread or any other goroutines unless they are specifically monitoring the context for cancellation.
	// calling cancel() after 2 seconds is canceling the parallel process (goroutine) that is running the longRunningOperation, but only if that operation takes more than 2 seconds to complete.
	
    time.Sleep(1 * time.Second)
	fmt.Println("234")
}



// 3. Deadlines
// You can set deadlines for operations using contexts, which is useful when you want to enforce time limits on certain tasks.

func operationWithDeadline(ctx context.Context) {
    select {
    case <-time.After(1 * time.Second):
        fmt.Println("Operation completed")
    case <-ctx.Done():
        fmt.Println("Operation timed out or cancelled")
    }
}

func deadline() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    operationWithDeadline(ctx)
}