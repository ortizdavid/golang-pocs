package repositories

import (
	"database/sql"

	"github.com/ortizdavid/golang-pocs/calculation-breakdown/models"
)

type InvoiceRepository struct {
	DB *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{
		DB: db,
	}
}

// Save persists the invoice and all breakdown steps within a single transaction
func (repo *InvoiceRepository) Save(invoice *models.Invoice) error {
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

func (repo *InvoiceRepository) GetByID(id string) (*models.Invoice, error) {
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

func (repo *InvoiceRepository) GetAll() ([]models.Invoice, error) {
	query := `SELECT id, total, currency, created_at FROM invoices`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var invoices []models.Invoice

	// 1. Fetch all invoices and close the main result set immediately to release the connection
	for rows.Next() {
		var inv models.Invoice
		if err := rows.Scan(&inv.ID, &inv.Total, &inv.Currency, &inv.CreatedAt); err != nil {
			rows.Close()
			return nil, err
		}
		invoices = append(invoices, inv)
	}
	rows.Close()

	// 2. Safely query steps for each invoice using the freed connection pool
	for i := range invoices {
		stepsQuery := `
            SELECT id, invoice_id, step_order, label, expression, formula, value, description 
            FROM invoice_breakdown_steps 
            WHERE invoice_id = ? 
            ORDER BY step_order ASC
        `
		stepRows, err := repo.DB.Query(stepsQuery, invoices[i].ID)
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

		invoices[i].Steps = steps
	}

	return invoices, nil
}
