package main

import (
    "fmt"
    "os"
)

func main() {
    // Basic string formatting
    name := "Alice"
    age := 30
    
    // Print formatted output to standard output
    fmt.Printf("Name: %s, Age: %d\n", name, age)
    // Output: Name: Alice, Age: 30

    // Format a string and return it
    formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)
    fmt.Println(formattedString)
    // Output: Name: Alice, Age: 30

    // Write formatted output to a specified writer (e.g., a file or buffer)
    fmt.Fprintf(os.Stdout, "Name: %s, Age: %d\n", name, age)
    // Output: Name: Alice, Age: 30

    // Formatting Verbs
    num := 123456
    pi := 3.14159
    
    // Format integer in decimal and hexadecimal
    fmt.Printf("Integer: %d\n", num)         // Integer: 123456
    fmt.Printf("Hexadecimal: %x\n", num)     // Hexadecimal: 1e240
    
    // Format a floating-point number
    fmt.Printf("Float: %.2f\n", pi)          // Float: 3.14

    // Formatting with width and precision
    amount := 12345.6789
    fmt.Printf("Name: %-10s Amount: %8.2f\n", name, amount)
    // Output: Name: Alice     Amount: 12345.68

    // Common formatting verbs
    // %s: Formats a string
    // %d: Formats an integer in decimal
    // %f: Formats a floating-point number
    // %x: Formats an integer in hexadecimal
    // %t: Formats a boolean
    // %v: Formats a value in default format
    // %+v: Formats a value with field names
    // %#v: Formats a value in Go-syntax representation

    // Example with boolean and default format
    isTrue := true
    fmt.Printf("Boolean: %t\n", isTrue)      // Boolean: true
    fmt.Printf("Default format: %v\n", pi)   // Default format: 3.14159
    fmt.Printf("Go-syntax representation: %#v\n", pi) // Go-syntax representation: 3.14159
}
