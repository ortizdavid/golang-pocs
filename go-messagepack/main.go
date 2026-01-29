package main

import (
	"log"
	"net/http"

	"github.com/ortizdavid/golang-pocs/go-messagepack/database"
	"github.com/ortizdavid/golang-pocs/go-messagepack/handlers"
	"github.com/ortizdavid/golang-pocs/go-messagepack/repositories"
	"github.com/ortizdavid/golang-pocs/go-messagepack/services"
)


func main() {

	db := database.InitDB("go_messagepack.db")
	defer db.Close()

	mux := http.NewServeMux()

	// service and repo
	repo := repositories.NewProductRepository(db)
	service := services.NewProductService(repo)

	// routes
	handler := handlers.NewProductHandler(service)
	mux.HandleFunc("POST /products", handler.Create)
	mux.HandleFunc("PUT /products/[id}", handler.Update)
	mux.HandleFunc("DELETE /products/[id}", handler.Delete)
	mux.HandleFunc("GET /products/[id}", handler.GetAll)
	mux.HandleFunc("GET /products/[id}", handler.GetByID)

	// Run Server
	log.Println("Product server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}