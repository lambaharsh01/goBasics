package main

import (
    "fmt"
    "os"
    "log"
)

func main() {
    // Path to the file
    filePath := "example.txt"

    // Content to append
    content := "Appending new content to the file.\n"

    // Open the file in append mode

    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// os.OpenFile: This function is used to open a file with specific options/modes and permissions.

	// filePath: The path to the file you're trying to open.

	// os.O_APPEND: This mode tells Go to append new data to the end of the file. If the file already contains data, new data will be added after the existing content.

	// os.O_CREATE: If the file doesn't exist, this mode will create the file.

	// os.O_WRONLY: This opens the file in write-only mode, meaning you can write to the file but not read from it.

	// 0644: This is the file permission mode. In this case:

	// Owner can read and write (6).
	// Group can read (4).
	// Others can read (4).
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close() // Ensure the file is closed after appending

    // Write content to file
    _, err = file.WriteString(content)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Content appended successfully!")
}
