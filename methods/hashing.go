package main;

import (
	"fmt"
	"crypto/sha256"
)


func main(){
	sha256Func()
}

func sha256Func(){
	// symilar to bcrypt js
	// hash function that converts data into a fixed-size string of 64 hexadecimal characters, no matter the input size.
	// Verifying data integrity (e.g., file verification)
	// Password hashing (although you might combine it with techniques like salting for security)
	// Digital signatures 

    input := "Hello, Go!"
    hash := sha256.New()
    hash.Write([]byte(input))
    hashedBytes := hash.Sum(nil)

    fmt.Printf("Input: %s\nSHA256 Hash: %x\n", input, hashedBytes);

	// It's a deterministic cryptographic function, meaning that the same input will always produce the same output. This is one of the key properties of cryptographic hashes:

	// Deterministic: The same input always gives the same output.
	// Irreversible: You cannot reverse the hash to get the original input.
	// Avalanche Effect: A small change in the input will lead to a significantly different output.
}

