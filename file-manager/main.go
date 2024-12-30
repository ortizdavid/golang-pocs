package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    // Get the current working directory
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Get the base of the current working directory
    folderName := filepath.Base(cwd)
    fmt.Println("Current folder name:", folderName)
}
