package main

import (
	"fmt"

	"github.com/elzidanecodes/DOMPETKU/internal/storage"
	"github.com/elzidanecodes/DOMPETKU/internal/transaction"
)

func main() {

	filePath := "data/transactions.json"

	transactions, err := storage.Load(filePath)
	if err != nil {
		fmt.Println("Error loading transactions:", err)
		return
	}

	for {
		fmt.Println("\n========== DOMPETKU ==========")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. List Transactions")
		fmt.Println("3. Search Transaction")
		fmt.Println("4. Update Transaction")
		fmt.Println("5. Delete Transaction")
		fmt.Println("6. Show Balance")
		fmt.Println("7. Show Summary")
		fmt.Println("8. Exit")

		fmt.Print("Choose menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\n--- Add Transaction ---")

			var id string
			var date string
			var transactionType string
			var amount float64
			var category string
			var note string

			fmt.Print("ID: ")
			fmt.Scanln(&id)

			fmt.Print("Date: ")
			fmt.Scanln(&date)

			fmt.Print("Type (Income/Expense): ")
			fmt.Scanln(&transactionType)

			fmt.Print("Amount: ")
			fmt.Scanln(&amount)

			fmt.Print("Category: ")
			fmt.Scanln(&category)

			fmt.Print("Note: ")
			fmt.Scanln(&note)

			newTransaction := transaction.Transaction{
				ID:       id,
				Date:     date,
				Type:     transactionType,
				Amount:   amount,
				Category: category,
				Note:     note,
			}

			transactions, err = transaction.Add(transactions, newTransaction)

			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			err = storage.Save(filePath, transactions)
			if err != nil {
				fmt.Println("Error saving transactions:", err)
				break
			}

			fmt.Println("Transaction added successfully")

		case 2:
			fmt.Println("\n--- List Transactions ---")
			transactionList := transaction.List(transactions)
			if len(transactionList) == 0 {
				fmt.Println("No transactions found")
				break
			}

			for _, item := range transactionList {
				fmt.Println("--------------------")
				fmt.Println("ID:", item.ID)
				fmt.Println("Date:", item.Date)
				fmt.Println("Type:", item.Type)
				fmt.Println("Amount:", item.Amount)
				fmt.Println("Category:", item.Category)
				fmt.Println("Note:", item.Note)
			}

			fmt.Println("--------------------")

		case 3:
			fmt.Println("\n--- Search Transaction ---")

			var searchID string
			fmt.Print("Enter Transaction ID to search: ")
			fmt.Scanln(&searchID)

			foundTransaction, err := transaction.Search(transactions, searchID)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			fmt.Println("Transaction found:")
			fmt.Println("--------------------")
			fmt.Println("ID:", foundTransaction.ID)
			fmt.Println("Date:", foundTransaction.Date)
			fmt.Println("Type:", foundTransaction.Type)
			fmt.Println("Amount:", foundTransaction.Amount)
			fmt.Println("Category:", foundTransaction.Category)
			fmt.Println("Note:", foundTransaction.Note)
			fmt.Println("--------------------")

		case 4:
			fmt.Println("\n--- Update Transaction ---")

			var id string

			fmt.Print("Enter transaction ID: ")
			fmt.Scanln(&id)

			var date string
			var transactionType string
			var amount float64
			var category string
			var note string

			fmt.Print("New Date: ")
			fmt.Scanln(&date)

			fmt.Print("New Type (Income/Expense): ")
			fmt.Scanln(&transactionType)

			fmt.Print("New Amount: ")
			fmt.Scanln(&amount)

			fmt.Print("New Category: ")
			fmt.Scanln(&category)

			fmt.Print("New Note: ")
			fmt.Scanln(&note)

			updatedTransaction := transaction.Transaction{
				Date:     date,
				Type:     transactionType,
				Amount:   amount,
				Category: category,
				Note:     note,
			}

			transactions, err = transaction.Update(transactions, id, updatedTransaction)

			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			err = storage.Save(filePath, transactions)
			if err != nil {
				fmt.Println("Error saving transactions:", err)
				break
			}

			fmt.Println("Transaction updated successfully")

		case 5:
			fmt.Println("\n--- Delete Transaction ---")

			var id string

			fmt.Print("Enter transaction ID: ")
			fmt.Scanln(&id)

			transactions, err = transaction.Delete(transactions, id)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			err = storage.Save(filePath, transactions)
			if err != nil {
				fmt.Println("Error saving transactions:", err)
				break
			}

			fmt.Println("Transaction deleted successfully")

		case 6:
			fmt.Println("\n--- Show Balance ---")

			balance := transaction.GetBalance(transactions)

			fmt.Println("Current Balance:", balance)

		case 7:
			fmt.Println("\n--- Show Summary ---")

			incomeSummary, expenseSummary := transaction.GetSummary(transactions)

			fmt.Println("\nIncome Summary:")
			for category, total := range incomeSummary {
				fmt.Println(category+":", total)
			}

			fmt.Println("\nExpense Summary:")
			for category, total := range expenseSummary {
				fmt.Println(category+":", total)
			}

		case 8:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid menu")
		}
	}
}
