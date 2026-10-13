package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/ortizdavid/golang-pocs/calculation-breakdown/services"
)

type CalculationHandler struct {
    Service *services.InvoiceService
}

func NewCalculationHandler(service *services.InvoiceService) *CalculationHandler {
    return &CalculationHandler{
        Service: service,
    }
}

type CreateInvoiceRequest struct {
    UnitPrice float64 `json:"unitPrice"`
    Quantity  float64 `json:"quantity"`
    TaxRate   float64 `json:"taxRate"`
}

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
            "POST /invoices":          "Calculate and save breakdown",
            "POST /invoices/simulate": "Simulate calculation without saving",
            "GET /invoices":           "List all saved invoices",
            "GET /invoices/{id}":   "Get calculation details by ID",
        },
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func (h *CalculationHandler) Calculate(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req CreateInvoiceRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    invoice, err := h.Service.CreateInvoice(req.UnitPrice, req.Quantity, req.TaxRate)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(invoice)
}

func (h *CalculationHandler) Simulate(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req CreateInvoiceRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    invoice := h.Service.SimulateInvoice(req.UnitPrice, req.Quantity, req.TaxRate)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(invoice)
}

func (h *CalculationHandler) GetInvoice(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Extacts id from endpoint /invoices/{id}
    id := r.PathValue("id")
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