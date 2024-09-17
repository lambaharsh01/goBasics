package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "Hello, World!"
    substr := "World"
    prefix := "Hello"
    suffix := "World"
    chars := "aeiou"
    r := 'o'
    
    // Contains checks if s contains substr
    fmt.Println(strings.Contains(s, substr)) // true
    
    // ContainsAny checks if s contains any of the characters in chars
    fmt.Println(strings.ContainsAny(s, chars)) // true
    
    // ContainsRune checks if s contains the rune r
    fmt.Println(strings.ContainsRune(s, r)) // true
    
    // Count counts the number of non-overlapping instances of substr in s
    fmt.Println(strings.Count(s, substr)) // 1
    
    // EqualFold reports whether s1 and s2 are equal, ignoring case
    fmt.Println(strings.EqualFold(s, "hello, world!")) // true
    
    // Fields splits s into a slice of substrings separated by white space
    fmt.Println(strings.Fields(s)) // [Hello, World!]
    
    // HasPrefix checks if s starts with prefix
    fmt.Println(strings.HasPrefix(s, prefix)) // true
    
    // HasSuffix checks if s ends with suffix
    fmt.Println(strings.HasSuffix(s, suffix)) // true
    
    // Index returns the index of the first occurrence of substr in s, or -1 if not found
    fmt.Println(strings.Index(s, substr)) // 7
    
    // IndexAny returns the index of the first occurrence of any of the characters in chars in s, or -1 if none are present
    fmt.Println(strings.IndexAny(s, chars)) // 4 (index of 'o')
    
    // IndexByte returns the index of the first occurrence of c in s, or -1 if not found
    fmt.Println(strings.IndexByte(s, 'o')) // 4
    
    // IndexRune returns the index of the first occurrence of r in s, or -1 if not found
    fmt.Println(strings.IndexRune(s, r)) // 4
    
    // Join concatenates the elements of words with a separator sep
    words := []string{"Go", "is", "awesome"}
    fmt.Println(strings.Join(words, " ")) // Go is awesome
    
    // Replace returns a copy of s with the first n non-overlapping instances of old replaced by new
    fmt.Println(strings.Replace(s, "World", "Go", -1)) // Hello, Go!
    
    // Split splits s into a slice of substrings separated by sep
    fmt.Println(strings.Split(s, ", ")) // [Hello World!]
    
    // ToLower returns a copy of s with all Unicode letters mapped to lowercase
    fmt.Println(strings.ToLower(s)) // hello, world!
    
    // ToUpper returns a copy of s with all Unicode letters mapped to uppercase
    fmt.Println(strings.ToUpper(s)) // HELLO, WORLD!
    
    // Trim returns a slice of s with all leading and trailing white space removed
    fmt.Println(strings.Trim(s, "H!")) // ello, World
}
