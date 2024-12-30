package main

import (
    "fmt"
    "time"
)

func main() {
    // Create a time object for "2022-01-01 12:00:09"
    date := time.Date(2022, time.January, 1, 12, 0, 9, 0, time.UTC)

    // Format the time object using the layout string
    formattedDate := date.Format("2006-01-02 15:04:05")

    // Print the formatted date
    fmt.Println(formattedDate)
}
