package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"expense-tracker/storage"
)

const BudgetFile = "data/budget.json"

type Budget struct {
	Amount float64 `json:"amount"`
}

func SetBudget(amount float64) error {
	budget := Budget{
		Amount: amount,
	}

	data, err := json.MarshalIndent(budget, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(BudgetFile, data, 0644)
}

func BudgetStatus() error {
	data, err := os.ReadFile(BudgetFile)
	if err != nil {
		return err
	}

	var budget Budget

	err = json.Unmarshal(data, &budget)
	if err != nil {
		return err
	}

	expenses, err := storage.LoadExpenses()
	if err != nil {
		return err
	}

	totalSpent := 0.0

	for _, expense := range expenses {
		totalSpent += expense.Amount
	}

	remaining := budget.Amount - totalSpent

	fmt.Printf("Budget:    $%.2f\n", budget.Amount)
	fmt.Printf("Spent:     $%.2f\n", totalSpent)
	fmt.Printf("Remaining: $%.2f\n", remaining)

	return nil
}