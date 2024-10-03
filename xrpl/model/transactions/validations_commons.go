package transactions

import (
	"fmt"
)

func ValidateRequiredField(tx FlatTransaction, field string, checkValidity func(interface{}) bool) error {
	// Check if the field is present in the transaction map.
	if _, ok := tx[field]; !ok {
		return fmt.Errorf("%s is missing", field)
	}

	// Check if the field is valid.
	if !checkValidity(tx[field]) {
		transactionType, _ := tx["TransactionType"].(string)
		return fmt.Errorf("%s: invalid field %s", transactionType, field)
	}

	return nil
}

// ValidateOptionalField validates an optional field in the transaction map.
func ValidateOptionalField(tx FlatTransaction, paramName string, checkValidity func(interface{}) bool) error {
	// Check if the field is present in the transaction map.
	if value, ok := tx[paramName]; ok {
		// Check if the field is valid.
		if !checkValidity(value) {
			transactionType, _ := tx["TransactionType"].(string)
			return fmt.Errorf("%s: invalid field %s", transactionType, paramName)
		}
	}

	return nil
}

// validateMemos validates the Memos field in the transaction map.
func validateMemos(memoWrapper []MemoWrapper) error {
	// loop through each memo and validate it
	for _, memo := range memoWrapper {
		isMemo, err := IsMemo(memo.Memo)
		if !isMemo {
			return err
		}
	}

	return nil
}

// validateSigners validates the Signers field in the transaction map.
func validateSigners(tx FlatTransaction) error {
	if tx["Signers"] != nil {
		signers, ok := tx["Signers"].([]FlatSigner)
		if !ok {
			return fmt.Errorf("BaseTransaction: invalid Signers")
		}
		for _, signer := range signers {
			if !IsSigner(signer) {
				return fmt.Errorf("BaseTransaction: invalid Signers")
			}
		}
	}

	return nil
}
