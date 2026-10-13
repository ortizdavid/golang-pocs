package main

import (
	"log"
	"net/http"

	"github.com/ortizdavid/golang-pocs/calculation-breakdown/database"
	"github.com/ortizdavid/golang-pocs/calculation-breakdown/handlers"
	"github.com/ortizdavid/golang-pocs/calculation-breakdown/repositories"
	"github.com/ortizdavid/golang-pocs/calculation-breakdown/services"
)

func main() {
	// Initialize the SQLite database and load tables.sql
	db := database.InitDB("calculation_breakdown.db")
	defer db.Close()

	mux := http.NewServeMux()

	// Repository, Service, and Handler
	repo := repositories.NewCalculationRepository(db)
	service := services.NewCalculationService(repo)
	handler := handlers.NewCalculationHandler(service)

	// HTTP Routes
	mux.HandleFunc("GET /", handler.Index)
	mux.HandleFunc("POST /calculations", handler.Calculate)
	mux.HandleFunc("POST /calculations/simulate", handler.Simulate)
	mux.HandleFunc("GET /calculations", handler.ListInvoices)
	mux.HandleFunc("GET /calculations/{id}", handler.GetInvoice)

	// Run the Server
	log.Println("Calculation Breakdown server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}