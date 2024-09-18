package main

import (
    "fmt"
    "os" // reads file
    "log"
)

func main() {
    // Path to the file
    filePath := "read.txt"

    // Read the file content
    content, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }

    // no need to close like this
    // defer content.Close()
    // os.ReadFile function is a convenience function that reads the entire file content into memory and returns a byte slice containing the file's data. It handles opening and closing the file internally, so you don’t have to explicitly close the file.


	// fmt.Println(content) // byte slice
	fmt.Println("File Content:")
    fmt.Println(string(content))
}
