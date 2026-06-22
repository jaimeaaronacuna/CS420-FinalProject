package commands

import (
	"fmt"

	"expense-tracker/storage"
)

func Summary() error {

	expenses, err := storage.LoadExpenses()
	if err != nil {
		return err
	}

	total := 0.0

	for _, expense := range expenses {
		total += expense.Amount
	}

	fmt.Printf("Total Spent: $%.2f\n", total)

	return nil
}
