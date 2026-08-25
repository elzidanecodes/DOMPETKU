package transaction

import (
	"errors"
)

type Transaction struct {
	ID       string
	Date     string
	Type     string
	Amount   float64
	Category string
	Note     string
}

var (
	ErrTransactionNotFound    = errors.New("transaction not found")
	ErrDuplicateTransactionID = errors.New("transaction ID already exists")
)

func Add(transactions []Transaction, newTransaction Transaction) ([]Transaction, error) {
	if err := ValidateTransaction(newTransaction); err != nil {
		return nil, err
	}

	for _, transaction := range transactions {
		if transaction.ID == newTransaction.ID {
			return nil, ErrDuplicateTransactionID
		}
	}

	return append(transactions, newTransaction), nil
}

func List(transactions []Transaction) []Transaction {
	return transactions
}

func Search(transactions []Transaction, id string) (Transaction, error) {
	index := findTransactionIndex(transactions, id)
	if index == -1 {
		return Transaction{}, ErrTransactionNotFound
	}
	return transactions[index], nil
}

func Update(transactions []Transaction, id string, updatedTransaction Transaction) ([]Transaction, error) {
	index := findTransactionIndex(transactions, id)
	if index == -1 {
		return nil, ErrTransactionNotFound
	}

	updatedTransaction.ID = id

	if err := ValidateTransaction(updatedTransaction); err != nil {
		return nil, err
	}
	transactions[index] = updatedTransaction
	return transactions, nil

}

func Delete(transactions []Transaction, id string) ([]Transaction, error) {
	index := findTransactionIndex(transactions, id)
	if index == -1 {
		return nil, ErrTransactionNotFound
	}

	return append(transactions[:index], transactions[index+1:]...), nil
}

func ValidateTransaction(transaction Transaction) error {
	if transaction.ID == "" {
		return errors.New("transaction ID cannot be empty")
	}
	if transaction.Date == "" {
		return errors.New("transaction date cannot be empty")
	}
	if transaction.Type != "Income" && transaction.Type != "Expense" {
		return errors.New("transaction type must be either 'Income' or 'Expense'")
	}
	if transaction.Amount <= 0 {
		return errors.New("transaction amount must be greater than zero")
	}
	if transaction.Category == "" {
		return errors.New("transaction category cannot be empty")
	}
	return nil
}

func findTransactionIndex(transactions []Transaction, id string) int {
	for index, transaction := range transactions {
		if transaction.ID == id {
			return index
		}
	}
	return -1
}

func GetBalance(transactions []Transaction) float64 {
	var balance float64
	for _, transaction := range transactions {
		switch transaction.Type {
		case "Income":
			balance += transaction.Amount
		case "Expense":
			balance -= transaction.Amount
		}
	}
	return balance
}

func GetSummary(transactions []Transaction) (map[string]float64, map[string]float64) {

	incomeSummary := make(map[string]float64)
	expenseSummary := make(map[string]float64)

	for _, transaction := range transactions {
		switch transaction.Type {
		case "Income":
			incomeSummary[transaction.Category] += transaction.Amount
		case "Expense":
			expenseSummary[transaction.Category] += transaction.Amount
		}
	}

	return incomeSummary, expenseSummary

}
