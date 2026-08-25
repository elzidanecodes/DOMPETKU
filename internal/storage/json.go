package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/elzidanecodes/DOMPETKU/internal/transaction"
)

func Load(filePath string) ([]transaction.Transaction, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []transaction.Transaction{}, nil
		}
		return nil, err
	}

	var transactions []transaction.Transaction

	err = json.Unmarshal(data, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func Save(filePath string, transactions []transaction.Transaction) error {
	data, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
