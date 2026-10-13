package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ortizdavid/golang-pocs/calculation-breakdown/services"
)

type CalculationHandler struct {
	Service *services.CalculationService
}

func NewCalculationHandler(service *services.CalculationService) *CalculationHandler {
	return &CalculationHandler{
		Service: service,
	}
}

// CalculationRequest defines the input structure for calculations and simulations
type CalculationRequest struct {
	UnitPrice float64 `json:"unitPrice"`
	Quantity  float64 `json:"quantity"`
	TaxRate   float64 `json:"taxRate"`
}

// Index provides basic info about the API
func (h *CalculationHandler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]interface{}{
		"app":         "Calculation Breakdown PoC",
		"version":     "1.0.0",
		"description": "API for audit-ready business rules storing calculation steps and totals",
		"endpoints": map[string]string{
			"POST /calculations":          "Calculate and save breakdown",
			"POST /calculations/simulate": "Simulate calculation without saving",
			"GET /calculations":           "List all saved calculations",
			"GET /calculations/{id}":      "Get calculation details by ID",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Calculate processes, persists, and returns the invoice with the full breakdown
func (h *CalculationHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	invoice, err := h.Service.CalculateAndSave(req.UnitPrice, req.Quantity, req.TaxRate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

// Simulate executes pure calculation and returns the breakdown without saving to the database
func (h *CalculationHandler) Simulate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Directly calls the service's pure calculation function
	invoice := h.Service.SimulateCalculation(req.UnitPrice, req.Quantity, req.TaxRate)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoice)
}

// GetInvoice fetches details of a specific invoice by ID
func (h *CalculationHandler) GetInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing invoice id", http.StatusBadRequest)
		return
	}

	invoice, err := h.Service.GetInvoiceDetails(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

// ListInvoices lists all persisted invoices
func (h *CalculationHandler) ListInvoices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	invoices, err := h.Service.ListAllInvoices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}