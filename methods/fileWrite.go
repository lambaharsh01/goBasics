package main

import (
    "fmt"
    "os"
    "log"
)

func main() {
    // Path to the file
    filePath := "example.txt"

    // Content to write
    content := []byte("Hello, Go!\nWelcome to file handling.")

    // Write content to file
    err := os.WriteFile(filePath, content, 0644)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("File written successfully!")
}
