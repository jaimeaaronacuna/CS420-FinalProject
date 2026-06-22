package commands

import (
	"expense-tracker/storage"
	"fmt"
)

// ListExpenses displays all saved expenses.
func ListExpenses() error {
	// Load all expenses
	expenses, err := storage.LoadExpenses()
	if err != nil {
		return err
	}

	// If there are no expenses, let the user know.
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return nil
	}

	// Print each expense
	for _, expense := range expenses {
		fmt.Printf("$%.2f - %s\n", expense.Amount, expense.Category)
	}

	return nil
}
