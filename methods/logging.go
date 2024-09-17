// log will basically console the activities with the timeStamp

// DEBUG: Detailed information for debugging purposes.
// INFO: General information about the flow of the application.
// WARN: Information about potentially harmful situations.
// ERROR: Error events that might still allow the application to continue running.
// FATAL: Very severe errors that will likely lead the application to abort.

package main

import (
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println("Received a request:", r.URL.Path) //The log.Println statements are used to log the flow of the application (starting the server, receiving a request).
    _, err := w.Write([]byte("Hello, World!"))
    if err != nil {
        log.Printf("Error writing response: %v\n", err); //log.Printf is used for formatted logging of errors.
    }
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Server failed: %v\n", err); //log.Fatalf is used to log a fatal error and terminate the program.
    }
}
