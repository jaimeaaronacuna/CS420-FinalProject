package commands

import (
	"fmt"
	"strings"

	"expense-tracker/storage"
)

func Search(category string) error {

	expenses, err := storage.LoadExpenses()
	if err != nil {
		return err
	}

	found := false

	for i, expense := range expenses {

		if strings.EqualFold(expense.Category, category) {
			fmt.Printf(
				"%d. %s - $%.2f\n",
				i+1,
				expense.Category,
				expense.Amount,
			)

			found = true
		}
	}

	if !found {
		fmt.Println("No expenses found.")
	}

	return nil
}