package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/elzidanecodes/DOMPETKU/internal/transaction"
)

func TestSaveAndLoad(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "transactions.json")

	transactions := []transaction.Transaction{
		{
			ID:       "TRX-001",
			Date:     "2026-08-24",
			Type:     "Expense",
			Amount:   25000,
			Category: "Food",
			Note:     "Lunch",
		},
		{
			ID:       "TRX-002",
			Date:     "2026-08-24",
			Type:     "Income",
			Amount:   5000000,
			Category: "Salary",
			Note:     "Monthly salary",
		},
	}

	err := Save(filePath, transactions)
	if err != nil {
		t.Fatalf("failed to save transactions: %v", err)
	}

	result, err := Load(filePath)
	if err != nil {
		t.Fatalf("failed to load transactions: %v", err)
	}

	if len(result) != len(transactions) {
		t.Errorf(
			"expected %d transactions, got %d",
			len(transactions),
			len(result),
		)
	}

	if result[0].ID != transactions[0].ID {
		t.Errorf(
			"expected ID %s, got %s",
			transactions[0].ID,
			result[0].ID,
		)
	}

	if result[1].Amount != transactions[1].Amount {
		t.Errorf(
			"expected amount %v, got %v",
			transactions[1].Amount,
			result[1].Amount,
		)
	}
}

func TestLoadFileNotFound(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "transactions.json")

	result, err := Load(filePath)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 0 {
		t.Errorf("expected empty transactions, got %d", len(result))
	}
}

func TestLoadInvalidJSON(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "transactions.json")

	invalidJSON := []byte(`{"invalid json`)

	err := os.WriteFile(filePath, invalidJSON, 0644)
	if err != nil {
		t.Fatalf("failed to write invalid JSON: %v", err)
	}

	_, err = Load(filePath)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
