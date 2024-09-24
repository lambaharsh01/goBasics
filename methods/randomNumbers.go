package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // Generate random numbers
    fmt.Println("Random Integer:", rand.Intn(100))
    fmt.Println("Random Float:", rand.Float64())
    fmt.Println("Random Integer in Range (10-50):", rand.Intn(41)+10)

	// min := 10
	// max := 50
	// randomInRange := rand.Intn(max-min) + min // Random integer between 10 and 49
	// fmt.Println("Random Integer in Range:", randomInRange)

    // Using a custom source
    source := rand.NewSource(42)
    r := rand.New(source)
    fmt.Println("Random Integer from Custom Source:", r.Intn(100))
}
