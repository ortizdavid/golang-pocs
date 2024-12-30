package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /items", getAllItemsHandler)
	mux.HandleFunc("GET /items/{id}", getItemHandler)
	mux.HandleFunc("POST /items", createItemHandler)
	mux.HandleFunc("PUT /items/{id}", updateItemHandler)
	mux.HandleFunc("DELETE /items/{id}", deleteItemHandler)
	mux.HandleFunc("GET /request-info/{id1}/other/{id2}", requestInfoHandler)

	fmt.Println("Listen at http://127.0.0.1:8000")
	http.ListenAndServe(":8000", mux)
}

func getAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Retrieving all items")
}

func getItemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintln(w, "Item ID: ", id)
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating item: ", r.Body)
}

func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintln(w, "Updated Item with id: ", id)
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintln(w, "Deleted Item with id: ", id)
}

func requestInfoHandler(w http.ResponseWriter, r *http.Request) {
	id1 := r.PathValue("id1")
	id2 := r.PathValue("id2")
	fmt.Fprintf(w, "Id1: %s\nId2: %s\nURL: %v\nPath: %v\nHost: %v", id1, id2, r.URL, r.URL.Path, r.URL.Host)
}