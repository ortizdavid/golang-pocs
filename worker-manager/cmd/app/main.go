
package main

import (
	"fmt"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main application %s", time.Now())
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}