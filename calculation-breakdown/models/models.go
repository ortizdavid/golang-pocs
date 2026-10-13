package models

import "time"

type Step struct {
	ID          string
	InvoiceID   string
	StepOrder   int
	Label       string
	Expression  string
	Formula     string
	Value       float64
	Description string
}

type Invoice struct {
	ID        string
	Total     float64
	Currency  string
	CreatedAt time.Time
	Steps     []Step
}