package main

import (
	"fmt"
	"os"
	"strconv"

	"expense-tracker/commands"
)

func main() {

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {

	case "add":

		if len(os.Args) < 4 {
			fmt.Println("Usage: add <amount> <category>")
			return
		}

		amount, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Println("Invalid amount")
			return
		}

		category := os.Args[3]

		err = commands.AddExpense(amount, category)
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "list":

		err := commands.ListExpenses()
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "summary":

		err := commands.Summary()
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "search":

		if len(os.Args) < 3 {
			fmt.Println("Usage: search <category>")
			return
		}

		err := commands.Search(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "budget":

		if len(os.Args) < 3 {
			fmt.Println("Usage:")
			fmt.Println("  budget set <amount>")
			fmt.Println("  budget status")
			return
		}

		switch os.Args[2] {

		case "set":

			if len(os.Args) < 4 {
				fmt.Println("Usage: budget set <amount>")
				return
			}

			amount, err := strconv.ParseFloat(os.Args[3], 64)
			if err != nil {
				fmt.Println("Invalid amount")
				return
			}

			err = commands.SetBudget(amount)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "status":

			err := commands.BudgetStatus()
			if err != nil {
				fmt.Println("Error:", err)
			}

		default:
			fmt.Println("Unknown budget command")
		}

	case "monthly":

		err := commands.Monthly()
		if err != nil {
			fmt.Println("Error:", err)
		}

	default:
		fmt.Println("Unknown command")
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Expense Tracker")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  add <amount> <category>")
	fmt.Println("  list")
	fmt.Println("  summary")
	fmt.Println("  search <category>")
	fmt.Println("  budget set <amount>")
	fmt.Println("  budget status")
	fmt.Println("  monthly")
}
