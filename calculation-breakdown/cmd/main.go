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
	repo := repositories.NewInvoiceRepository(db)
	service := services.NewInvoiceService(repo)
	handler := handlers.NewCalculationHandler(service)

	// HTTP Routes
	mux.HandleFunc("GET /", handler.Index)
	mux.HandleFunc("POST /invoices", handler.Calculate)
	mux.HandleFunc("POST /invoices/simulate", handler.Simulate)
	mux.HandleFunc("GET /invoices", handler.ListInvoices)
	mux.HandleFunc("GET /invoices/{id}", handler.GetInvoice)

	// Run the Server
	log.Println("Calculation Breakdown server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
