package transaction

import "testing"

func TestAddSuccess(t *testing.T) {
	transactions := []Transaction{}

	newTransaction := Transaction{
		ID:       "TRX-001",
		Date:     "2026-08-24",
		Type:     "Expense",
		Amount:   25000,
		Category: "Food",
		Note:     "Lunch",
	}

	result, err := Add(transactions, newTransaction)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 transaction, got %d", len(result))
	}

	if result[0].ID != "TRX-001" {
		t.Errorf("expected ID TRX-001, got %s", result[0].ID)
	}
}

func TestAddDuplicateID(t *testing.T) {
	transactions := []Transaction{
		{
			ID:       "TRX-001",
			Date:     "2026-08-24",
			Type:     "Expense",
			Amount:   25000,
			Category: "Food",
		},
	}

	newTransaction := Transaction{
		ID:       "TRX-001",
		Date:     "2026-08-24",
		Type:     "Income",
		Amount:   5000000,
		Category: "Salary",
	}

	_, err := Add(transactions, newTransaction)

	if err != ErrDuplicateTransactionID {
		t.Errorf(
			"expected ErrDuplicateTransactionID, got %v",
			err,
		)
	}
}

func TestValidateTransaction(t *testing.T) {
	tests := []struct {
		name        string
		transaction Transaction
		expectedErr string
	}{
		{
			name: "valid transaction",
			transaction: Transaction{
				ID:       "TRX-001",
				Date:     "2026-08-24",
				Type:     "Expense",
				Amount:   25000,
				Category: "Food",
			},
			expectedErr: "",
		},
		{
			name: "empty ID",
			transaction: Transaction{
				Date:     "2026-08-24",
				Type:     "Expense",
				Amount:   25000,
				Category: "Food",
			},
			expectedErr: "transaction ID cannot be empty",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateTransaction(test.transaction)

			if test.expectedErr == "" && err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			if test.expectedErr != "" {
				if err == nil {
					t.Errorf("expected error %q, got nil", test.expectedErr)
				}

				if err != nil && err.Error() != test.expectedErr {
					t.Errorf("expected error %q, got %q", test.expectedErr, err.Error())
				}
			}
		})
	}
}

func TestUpdateSuccess(t *testing.T) {
	transactions := []Transaction{
		{
			ID:       "TRX-001",
			Date:     "2026-08-24",
			Type:     "Expense",
			Amount:   25000,
			Category: "Food",
			Note:     "Lunch",
		},
	}

	updatedTransaction := Transaction{
		Date:     "2026-08-25",
		Type:     "Expense",
		Amount:   50000,
		Category: "Transport",
		Note:     "Taxi",
	}

	result, err := Update(
		transactions,
		"TRX-001",
		updatedTransaction,
	)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if result[0].ID != "TRX-001" {
		t.Errorf(
			"expected ID TRX-001, got %s",
			result[0].ID,
		)
	}

	if result[0].Amount != 50000 {
		t.Errorf(
			"expected amount 50000, got %v",
			result[0].Amount,
		)
	}

	if result[0].Category != "Transport" {
		t.Errorf(
			"expected category Transport, got %s",
			result[0].Category,
		)
	}
}

func TestUpdateTransactionNotFound(t *testing.T) {
	transactions := []Transaction{
		{
			ID:       "TRX-001",
			Date:     "2026-08-24",
			Type:     "Expense",
			Amount:   25000,
			Category: "Food",
		},
	}

	updatedTransaction := Transaction{
		Date:     "2026-08-25",
		Type:     "Expense",
		Amount:   50000,
		Category: "Transport",
	}

	_, err := Update(
		transactions,
		"TRX-999",
		updatedTransaction,
	)

	if err != ErrTransactionNotFound {
		t.Errorf(
			"expected ErrTransactionNotFound, got %v",
			err,
		)
	}
}

func TestUpdateInvalidTransaction(t *testing.T) {
	transactions := []Transaction{
		{
			ID:       "TRX-001",
			Date:     "2026-08-24",
			Type:     "Expense",
			Amount:   25000,
			Category: "Food",
		},
	}

	updatedTransaction := Transaction{
		Date:     "2026-08-25",
		Type:     "Expense",
		Amount:   0,
		Category: "Transport",
	}

	_, err := Update(
		transactions,
		"TRX-001",
		updatedTransaction,
	)

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestDeleteSuccess(t *testing.T) {
	transactions := []Transaction{
		{
			ID:       "TRX-001",
			Date:     "2026-08-24",
			Type:     "Expense",
			Amount:   25000,
			Category: "Food",
		},
		{
			ID:       "TRX-002",
			Date:     "2026-08-24",
			Type:     "Income",
			Amount:   5000000,
			Category: "Salary",
		},
	}

	result, err := Delete(transactions, "TRX-001")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 transaction, got %d", len(result))
	}

	if result[0].ID != "TRX-002" {
		t.Errorf(
			"expected remaining transaction ID TRX-002, got %s",
			result[0].ID,
		)
	}
}

func TestDeleteTransactionNotFound(t *testing.T) {
	transactions := []Transaction{
		{
			ID:       "TRX-001",
			Date:     "2026-08-24",
			Type:     "Expense",
			Amount:   25000,
			Category: "Food",
		},
	}

	_, err := Delete(transactions, "TRX-999")

	if err != ErrTransactionNotFound {
		t.Errorf(
			"expected ErrTransactionNotFound, got %v",
			err,
		)
	}
}

func TestSearchSuccess(t *testing.T) {
	transactions := []Transaction{
		{
			ID:       "TRX-001",
			Date:     "2026-08-24",
			Type:     "Expense",
			Amount:   25000,
			Category: "Food",
		},
		{
			ID:       "TRX-002",
			Date:     "2026-08-24",
			Type:     "Income",
			Amount:   5000000,
			Category: "Salary",
		},
	}

	result, err := Search(transactions, "TRX-002")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if result.ID != "TRX-002" {
		t.Errorf(
			"expected transaction ID TRX-002, got %s",
			result.ID,
		)
	}

	if result.Category != "Salary" {
		t.Errorf(
			"expected category Salary, got %s",
			result.Category,
		)
	}
}

func TestSearchTransactionNotFound(t *testing.T) {
	transactions := []Transaction{
		{
			ID:       "TRX-001",
			Date:     "2026-08-24",
			Type:     "Expense",
			Amount:   25000,
			Category: "Food",
		},
	}

	_, err := Search(transactions, "TRX-999")

	if err != ErrTransactionNotFound {
		t.Errorf(
			"expected ErrTransactionNotFound, got %v",
			err,
		)
	}
}

func TestGetBalance(t *testing.T) {
	transactions := []Transaction{
		{
			ID:       "TRX-001",
			Type:     "Income",
			Amount:   5000000,
			Category: "Salary",
		},
		{
			ID:       "TRX-002",
			Type:     "Expense",
			Amount:   25000,
			Category: "Food",
		},
		{
			ID:       "TRX-003",
			Type:     "Expense",
			Amount:   50000,
			Category: "Transport",
		},
	}

	result := GetBalance(transactions)

	expected := 4925000.0

	if result != expected {
		t.Errorf(
			"expected balance %v, got %v",
			expected,
			result,
		)
	}
}

func TestGetBalanceEmpty(t *testing.T) {
	transactions := []Transaction{}

	result := GetBalance(transactions)

	expected := 0.0

	if result != expected {
		t.Errorf(
			"expected balance %v, got %v",
			expected,
			result,
		)
	}
}

func TestGetSummary(t *testing.T) {
	transactions := []Transaction{
		{
			ID:       "TRX-001",
			Type:     "Income",
			Amount:   5000000,
			Category: "Salary",
		},
		{
			ID:       "TRX-002",
			Type:     "Income",
			Amount:   1000000,
			Category: "Freelance",
		},
		{
			ID:       "TRX-003",
			Type:     "Expense",
			Amount:   25000,
			Category: "Food",
		},
		{
			ID:       "TRX-004",
			Type:     "Expense",
			Amount:   30000,
			Category: "Food",
		},
		{
			ID:       "TRX-005",
			Type:     "Expense",
			Amount:   50000,
			Category: "Transport",
		},
	}

	incomeSummary, expenseSummary := GetSummary(transactions)

	expectedIncome := map[string]float64{
		"Salary":    5000000,
		"Freelance": 1000000,
	}

	expectedExpense := map[string]float64{
		"Food":      55000,
		"Transport": 50000,
	}

	for category, expected := range expectedIncome {
		if incomeSummary[category] != expected {
			t.Errorf(
				"expected income %s to be %v, got %v",
				category,
				expected,
				incomeSummary[category],
			)
		}
	}

	for category, expected := range expectedExpense {
		if expenseSummary[category] != expected {
			t.Errorf(
				"expected expense %s to be %v, got %v",
				category,
				expected,
				expenseSummary[category],
			)
		}
	}
}
