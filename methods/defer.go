// defer in Go as a way to tell your program, "Hey, after I'm done with everything else, make sure to do this last thing."
// In Go, you use defer to ensure that certain tasks, like closing a file or releasing a lock, are done at the very end, no matter what happens in the meantime. You don't have to worry about remembering to do it later; defer takes care of it for you.


// Database Connections: Always close your database connections when you're done using them. For example, if you're using a SQL database with a library like gorm or sqlx, make sure to call db.Close() when your application shuts down or when the connection is no longer needed.

// File Handles: If you're reading or writing files, use defer to close file handles right after opening them. This ensures files are properly closed even if an error occurs or the function returns early.

// Network Connections: If you're working with raw network connections or any kind of network I/O, ensure you close these connections to avoid resource leaks.

// HTTP Clients: If you create HTTP clients or servers, be mindful of their lifecycle. For example, if you create a custom HTTP server, ensure you handle graceful shutdowns by calling methods to stop the server and clean up resources.

// Custom Resources: If you have other types of resources (like caches, external services, etc.), make sure to follow their respective documentation on how to properly close or shut them down.

// For Gin itself, while the framework handles many aspects for you, focus on closing any resources your handlers or middleware might use:

// Database connections in your handlers or middleware.
// Files that you open during request handling.
// External services or APIs that you interact with.

package main

import (
	"fmt"
	"io"
	"os"
	"log"
)

func main() {
	// Open the file
	file, err := os.Open("read.txt");
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when done

	// Read the file content
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Print the file content
	fmt.Println(string(data))
}
