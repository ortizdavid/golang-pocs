package services

import (
	"fmt"

	"github.com/ortizdavid/golang-pocs/calculation-breakdown/models"
	"github.com/ortizdavid/golang-pocs/calculation-breakdown/repositories"
)

type CalculationService struct {
	Repo *repositories.CalculationRepository
}

func NewCalculationService(repo *repositories.CalculationRepository) *CalculationService {
	return &CalculationService{
		Repo: repo,
	}
}

// CalculateAndSave executes pure calculation and persists it into the database
func (s *CalculationService) CalculateAndSave(base float64, qty float64, taxRate float64) (*models.Invoice, error) {
	// Delegates to the pure function that generates the formula history
	invoice := s.calculateBreakdown(base, qty, taxRate)

	// Persists everything via transactional repository
	err := s.Repo.Save(&invoice)
	if err != nil {
		return nil, fmt.Errorf("failed to save calculation breakdown: %w", err)
	}

	return &invoice, nil
}

// SimulateCalculation executes pure calculation without touching the database
func (s *CalculationService) SimulateCalculation(base float64, qty float64, taxRate float64) models.Invoice {
	return s.calculateBreakdown(base, qty, taxRate)
}

func (s *CalculationService) GetInvoiceDetails(id string) (*models.Invoice, error) {
	return s.Repo.GetByID(id)
}

func (s *CalculationService) ListAllInvoices() ([]models.Invoice, error) {
	return s.Repo.GetAll()
}