package models

// Expense represents a single expense record.
type Expense struct {
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
}
