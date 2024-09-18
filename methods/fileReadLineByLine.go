package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
)

func main() {
    // Path to the file
    filePath := "read.txt"

    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close() // Ensure the file is closed after reading

    // READING FILE LINE BY LINE

    // METHOD 1: read the whole file
    // scanner := bufio.NewScanner(file)
    // for scanner.Scan() {
    //     fmt.Println(scanner.Text()) // give me test of the read protion
    // }


    // METHOD 2 : stop reading if i have found what i was looking for 
    scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            line := scanner.Text()  // Get the current line

            fmt.Println(line)  // Process the line

            if line == "Useful but Not Always Critical:" {
                fmt.Println("Found what I was looking for, stopping.")
                break  // Exit the loop once the target string is found
            }
        }



	//returns any error that may have occurred during the scanning process. This method is called after you finish reading the file with the bufio.Scanner to check if there were any errors while scanning lines.
    if err := scanner.Err(); err != nil { //you are both declaring a variable and checking its value in a single if statement.
        log.Fatal(err)
    }
}
