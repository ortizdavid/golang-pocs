package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ortizdavid/golang-pocs/calculation-breakdown/models"
)

// calculateBreakdown is a pure function that only processes steps without touching the DB
func (s *InvoiceService) calculateBreakdown(base float64, qty float64, taxRate float64) models.Invoice {
	invoiceID := uuid.New().String()

	// Step 1: Base Calculation (Quantity x Unit Price)
	baseValue := calculateBase(qty, base)
	step1 := models.Step{
		ID:          uuid.New().String(),
		InvoiceID:   invoiceID,
		StepOrder:   1,
		Label:       "Base Calculation",
		Expression:  "qty * unitPrice",
		Formula:     fmt.Sprintf("%.0f * %.2f = %.2f", qty, base, baseValue),
		Value:       baseValue,
		Description: "Quantity multiplied by unit price",
	}

	// Step 2: Tax Calculation
	taxValue := calculateTax(baseValue, taxRate)
	step2 := models.Step{
		ID:          uuid.New().String(),
		InvoiceID:   invoiceID,
		StepOrder:   2,
		Label:       "Tax",
		Expression:  "baseValue * taxRate",
		Formula:     fmt.Sprintf("%.2f * %.2f = %.2f", baseValue, taxRate, taxValue),
		Value:       taxValue,
		Description: fmt.Sprintf("Tax rate incidence of %.0f%%", taxRate*100),
	}

	total := baseValue + taxValue

	return models.Invoice{
		ID:       invoiceID,
		Total:    total,
		Currency: "AOA",
		Steps:    []models.Step{step1, step2},
	}
}

func calculateBase(qty float64, base float64) float64 {
	return qty * base
}

func calculateTax(baseValue float64, taxRate float64) float64 {
	return baseValue * taxRate
}
