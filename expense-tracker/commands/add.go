package commands

import (
	"expense-tracker/models"
	"expense-tracker/storage"
)

// AddExpense creates a new expense and saves it.
func AddExpense(amount float64, category string) error {
	// Load existing expenses
	expenses, err := storage.LoadExpenses()
	if err != nil {
		return err
	}

	// Create the new expense
	expense := models.Expense{
		Amount:   amount,
		Category: category,
	}

	// Add it to the list
	expenses = append(expenses, expense)

	// Save the updated list
	return storage.SaveExpenses(expenses)
}
