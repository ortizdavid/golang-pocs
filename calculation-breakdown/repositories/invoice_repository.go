package repositories

import (
	"database/sql"

	"github.com/ortizdavid/golang-pocs/calculation-breakdown/models"
)

type CalculationRepository struct {
	DB *sql.DB
}

func NewCalculationRepository(db *sql.DB) *CalculationRepository {
	return &CalculationRepository{
		DB: db,
	}
}

// Save persists the invoice and all breakdown steps within a single transaction
func (repo *CalculationRepository) Save(invoice *models.Invoice) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Inserts the main invoice
	invoiceQuery := `INSERT INTO invoices (id, total, currency) VALUES (?, ?, ?)`
	_, err = tx.Exec(invoiceQuery, invoice.ID, invoice.Total, invoice.Currency)
	if err != nil {
		return err
	}

	// Inserts each atomic breakdown step
	stepQuery := `
		INSERT INTO invoice_breakdown_steps (id, invoice_id, step_order, label, expression, formula, value, description) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	for _, step := range invoice.Steps {
		_, err = tx.Exec(stepQuery, step.ID, invoice.ID, step.StepOrder, step.Label, step.Expression, step.Formula, step.Value, step.Description)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (repo *CalculationRepository) GetByID(id string) (*models.Invoice, error) {
	// Fetches the main invoice
	invoiceQuery := `SELECT id, total, currency, created_at FROM invoices WHERE id = ?`
	row := repo.DB.QueryRow(invoiceQuery, id)

	var inv models.Invoice
	err := row.Scan(&inv.ID, &inv.Total, &inv.Currency, &inv.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	// Fetches the associated steps ordered chronologically
	stepsQuery := `
		SELECT id, invoice_id, step_order, label, expression, formula, value, description 
		FROM invoice_breakdown_steps 
		WHERE invoice_id = ? 
		ORDER BY step_order ASC
	`
	rows, err := repo.DB.Query(stepsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var steps []models.Step
	for rows.Next() {
		var s models.Step
		if err := rows.Scan(&s.ID, &s.InvoiceID, &s.StepOrder, &s.Label, &s.Expression, &s.Formula, &s.Value, &s.Description); err != nil {
			return nil, err
		}
		steps = append(steps, s)
	}

	inv.Steps = steps
	return &inv, nil
}

func (repo *CalculationRepository) GetAll() ([]models.Invoice, error) {
	query := `SELECT id, total, currency, created_at FROM invoices`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []models.Invoice
	for rows.Next() {
		var inv models.Invoice
		if err := rows.Scan(&inv.ID, &inv.Total, &inv.Currency, &inv.CreatedAt); err != nil {
			return nil, err
		}

		// Loads the steps for each invoice
		stepsQuery := `
			SELECT id, invoice_id, step_order, label, expression, formula, value, description 
			FROM invoice_breakdown_steps 
			WHERE invoice_id = ? 
			ORDER BY step_order ASC
		`
		stepRows, err := repo.DB.Query(stepsQuery, inv.ID)
		if err != nil {
			return nil, err
		}
		
		var steps []models.Step
		for stepRows.Next() {
			var s models.Step
			if err := stepRows.Scan(&s.ID, &s.InvoiceID, &s.StepOrder, &s.Label, &s.Expression, &s.Formula, &s.Value, &s.Description); err != nil {
				stepRows.Close()
				return nil, err
			}
			steps = append(steps, s)
		}
		stepRows.Close()

		inv.Steps = steps
		invoices = append(invoices, inv)
	}
	return invoices, nil
}

func (repo *CalculationRepository) Delete(id string) error {
	query := `DELETE FROM invoices WHERE id = ?`
	_, err := repo.DB.Exec(query, id)
	return err
}

func (repo *CalculationRepository) Exists(id string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM invoices WHERE id = ?)`
	err := repo.DB.QueryRow(query, id).Scan(&exists)
	return exists, err
}