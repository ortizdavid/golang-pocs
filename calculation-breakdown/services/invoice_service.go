package services

import (
	"fmt"

	"github.com/ortizdavid/golang-pocs/calculation-breakdown/models"
	"github.com/ortizdavid/golang-pocs/calculation-breakdown/repositories"
)

type InvoiceService struct {
	Repo *repositories.InvoiceRepository
}

func NewInvoiceService(repo *repositories.InvoiceRepository) *InvoiceService {
	return &InvoiceService{
		Repo: repo,
	}
}

// CreateInvoice executes pure calculation and persists it into the database
func (s *InvoiceService) CreateInvoice(base float64, qty float64, taxRate float64) (*models.Invoice, error) {
	// Delegates to the pure function that generates the formula history
	invoice := s.calculateBreakdown(base, qty, taxRate)

	err := s.Repo.Save(&invoice)
	if err != nil {
		return nil, fmt.Errorf("failed to save calculation breakdown: %w", err)
	}

	return &invoice, nil
}

// SimulateInvoice executes pure calculation without touching the database
func (s *InvoiceService) SimulateInvoice(base float64, qty float64, taxRate float64) models.Invoice {
	return s.calculateBreakdown(base, qty, taxRate)
}

func (s *InvoiceService) GetInvoiceDetails(id string) (*models.Invoice, error) {
	return s.Repo.GetByID(id)
}

func (s *InvoiceService) ListAllInvoices() ([]models.Invoice, error) {
	return s.Repo.GetAll()
}
