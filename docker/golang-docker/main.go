package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %v\n", time.Now())
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Docker Image\n")
}

func main() {
	
	http.HandleFunc("/", greet)
	http.HandleFunc("/about", about)
	fmt.Println("Listening on: http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}