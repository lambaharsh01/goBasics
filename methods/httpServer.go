package main

import (
    "fmt"
    "net/http"
)

// Handler function for the root path
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
}
// Handler function for the root path
func baseHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world! FROM base")
}

func main() {
    // Register the handler function for the root path
    http.HandleFunc("/", handler)
    http.HandleFunc("/base", baseHandler)

    // Start the server on port 8080
    fmt.Println("Server starting at http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server failed:", err)
    }
}
