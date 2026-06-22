package commands

import (
	"expense-tracker/storage"
	"fmt"
)

func Monthly() error {
	expenses, err := storage.LoadExpenses()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return nil
	}

	total := 0.0
	fmt.Println("Monthly Report")

	for _, expense := range expenses {
		fmt.Printf("$%.2f - %s\n", expense.Amount, expense.Category)
		total += expense.Amount
	}

	fmt.Printf("\nTotal: $%.2f\n", total)
	return nil
}
