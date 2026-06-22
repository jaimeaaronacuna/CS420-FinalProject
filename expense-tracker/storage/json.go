package storage

import (
	"encoding/json"
	"os"

	"your-module-name/models"
)

const fileName = "expenses.json"

// LoadExpenses reads all expenses from the JSON file.
func LoadExpenses() ([]models.Expense, error) {
	var expenses []models.Expense

	// If the file doesn't exist yet, return an empty slice.
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return expenses, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

// SaveExpenses writes all expenses to the JSON file.
func SaveExpenses(expenses []models.Expense) error {
	data, err := json.MarshalIndent(expenses, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
