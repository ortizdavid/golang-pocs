package main

import (
    "flag"
    "fmt"
)

func main() {
    // Define flags
    var name string
    var age int
    var married bool

    // Parse flags
    flag.StringVar(&name, "name", "", "The name of the person")
    flag.IntVar(&age, "age", 0, "The age of the person")
    flag.BoolVar(&married, "married", false, "Whether the person is married or not")

    flag.Parse()

    // Output the values of the flags
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Married:", married)
}
