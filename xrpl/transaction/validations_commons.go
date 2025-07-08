package transaction

import (
	"fmt"
	"reflect"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// isArray verifies the form and type of an Array at runtime.
// Equivalent to xrpl.js isArray function.
func isArray(input interface{}) bool {
	if input == nil {
		return false
	}
	val := reflect.ValueOf(input)
	return val.Kind() == reflect.Slice || val.Kind() == reflect.Array
}

// isRecord verifies the form and type of a Record/Object at runtime.
// Equivalent to xrpl.js isRecord function.
func isRecord(value interface{}) bool {
	if value == nil {
		return false
	}
	val := reflect.ValueOf(value)
	// Check if it's an object (map or struct) but not an array/slice
	return (val.Kind() == reflect.Map || val.Kind() == reflect.Struct) &&
		val.Kind() != reflect.Slice && val.Kind() != reflect.Array
}

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
func validateMemos(memoWrapper []types.MemoWrapper) error {
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
func validateSigners(signers []Signer) error {
	// loop through each signer and validate it
	for _, signer := range signers {
		isSigner, err := IsSigner(signer.SignerData)
		if !isSigner {
			return err
		}
	}

	return nil
}
