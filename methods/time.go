package main

import (
    "fmt"
    "time"
)

func main() {
    // Getting the current time
    currentTime := time.Now()
    fmt.Println("Current time:", currentTime)

    // Formatting the current time
    formattedTime := currentTime.Format("2006-01-02 15:04:05")
    fmt.Println("Formatted time:", formattedTime)

    // Parsing a time string
    layout := "2006-01-02 15:04:05"
    value := "2024-09-17 13:45:00"
    parsedTime, err := time.Parse(layout, value)
    if err != nil {
        fmt.Println("Error parsing time:", err)
    } else {
        fmt.Println("Parsed time:", parsedTime)
    }

    // Working with time durations
    duration := 5 * time.Minute
    fmt.Println("Duration:", duration)

    // Adding a duration to the current time
    futureTime := currentTime.Add(10 * time.Hour)
    fmt.Println("Future time:", futureTime)

    // Subtracting a duration from the current time
    pastTime := currentTime.Add(-1 * time.Hour)
    elapsed := currentTime.Sub(pastTime)
    fmt.Println("Elapsed time:", elapsed)

    // Time comparisons
    later := currentTime.Add(1 * time.Hour)
    fmt.Println("Is now before later?", currentTime.Before(later))

    earlier := currentTime.Add(-2 * time.Hour)
    fmt.Println("Is now after earlier?", currentTime.After(earlier))

    // Working with Unix timestamps
    timestamp := time.Unix(1694942400, 0) // Example Unix timestamp
    fmt.Println("Timestamp:", timestamp)

    // Getting the Unix timestamp from the current time
    seconds := currentTime.Unix()
    fmt.Println("Unix timestamp:", seconds)

    // Working with time zones
    loc, err := time.LoadLocation("America/New_York")
    if err != nil {
        fmt.Println("Error loading location:", err)
    } else {
        nowInNY := currentTime.In(loc)
        fmt.Println("Current time in New York:", nowInNY)
    }


	// currentTime := time.Now()

    // Formatting the current time as dd-mm-yyyy
    // formattedDate := currentTime.Format("02-01-2006")
    formattedDate := currentTime.Format("2006-01-02")
    fmt.Println("Current date in dd-mm-yyyy format:", formattedDate)
}
