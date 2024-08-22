package transactions

import (
	"testing"
)

func TestValidateBaseTx(t *testing.T) {
	// Test case 1: Valid BaseTx
	tx1 := &BaseTx{
		Account:         "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
		TransactionType: "Payment",
	}
	// Expect no panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test case 1: Unexpected panic: %v", r)
		}
	}()
	ValidateBaseTx(tx1)

	// Test case 2: Missing Account
	tx2 := &BaseTx{
		TransactionType: "Payment",
	}
	// Expect panic with error message
	defer func() {
		if r := recover(); r != nil {
			errMsg := "base transaction: missing Account"
			if r != errMsg {
				t.Errorf("Test case 2: Expected panic with error message '%s', got '%v'", errMsg, r)
			}
		} else {
			t.Errorf("Test case 2: Expected panic with error message, but no panic occurred")
		}
	}()
	ValidateBaseTx(tx2)

	// Test case 3: Missing TransactionType
	tx3 := &BaseTx{
		Account: "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
	}
	// Expect panic with error message
	defer func() {
		if r := recover(); r != nil {
			errMsg := "base transaction: missing TransactionType"
			if r != errMsg {
				t.Errorf("Test case 3: Expected panic with error message '%s', got '%v'", errMsg, r)
			}
		} else {
			t.Errorf("Test case 3: Expected panic with error message, but no panic occurred")
		}
	}()
	ValidateBaseTx(tx3)

	// Add more test cases as needed
}
