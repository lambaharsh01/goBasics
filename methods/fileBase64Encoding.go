
// 	method of encoding binary data (such as images or files) into an ASCII string format. Often used when transmitting data over media designed to handle text
// 	Encoding data for transfer via HTTP (URLs, APIs)
// 	Encoding data in JWT tokens or cookies
// 	Embedding images or other binary content directly into text (HTML, JSON)

package main

import (
    "encoding/base64" // similar to fs
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    // Specify the file path (example: "example.txt")
    filePath := "read.txt"

    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Read the file contents into a byte slice
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

    // Encode the file content to Base64
    encoded := base64.StdEncoding.EncodeToString(fileBytes)

    // Print the encoded data
    fmt.Println("Base64 Encoded File Content:")
    fmt.Println(encoded)

    // Optionally, decode the Base64 back to the original content
    decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        log.Fatal(err)
    }

    // Print the decoded content as a string
    fmt.Println("\nDecoded File Content:")
    fmt.Println(string(decodedBytes))
}
