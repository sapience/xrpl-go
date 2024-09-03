package transactions

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/utils/typecheck"
)

func ValidateBaseTransaction(tx FlatTransaction) error {
	ValidateRequiredField(tx, "TransactionType", typecheck.IsString)
	ValidateRequiredField(tx, "Account", typecheck.IsString)

	// optional fields
	ValidateOptionalField(tx, "Fee", typecheck.IsString)
	ValidateOptionalField(tx, "Sequence", typecheck.IsInt)
	ValidateOptionalField(tx, "AccountTxnID", typecheck.IsString)
	ValidateOptionalField(tx, "LastLedgerSequence", typecheck.IsInt)
	ValidateOptionalField(tx, "SourceTag", typecheck.IsInt)
	ValidateOptionalField(tx, "SigningPubKey", typecheck.IsString)
	ValidateOptionalField(tx, "TicketSequence", typecheck.IsInt)
	ValidateOptionalField(tx, "TxnSignature", typecheck.IsString)
	ValidateOptionalField(tx, "NetworkID", typecheck.IsInt)

	// memos
	validateMemos(tx)

	// signers
	validateSigners(tx)

	return nil
}

func ValidateRequiredField(tx FlatTransaction, field string, checkValidity func(interface{}) bool) {
	// Check if the field is present in the transaction map.
	if _, ok := tx[field]; !ok {
		panic(field + " is missing")
	}

	// Check if the field is valid.
	if !checkValidity(tx[field]) {
		transactionType, _ := tx["TransactionType"].(string)
		panic(fmt.Errorf("%s: invalid field %s", transactionType, field))
	}
}

// ValidateOptionalField validates an optional field in the transaction map.
func ValidateOptionalField(tx FlatTransaction, paramName string, checkValidity func(interface{}) bool) {
	// Check if the field is present in the transaction map.
	if value, ok := tx[paramName]; ok {
		// Check if the field is valid.
		if !checkValidity(value) {
			transactionType, _ := tx["TransactionType"].(string)
			panic(fmt.Errorf("%s: invalid field %s", transactionType, paramName))
		}
	}
}

// validateMemos validates the Memos field in the transaction map.
func validateMemos(tx FlatTransaction) {
	// Check if the field Memos is set
	if tx["Memos"] != nil {
		flatMemoWrappers, ok := tx["Memos"].([]FlatMemoWrapper)
		if !ok {
			panic("BaseTransaction: invalid Memos conversion to []FlatMemoWrapper")
		}
		// loop through each memo and validate it
		for _, memo := range flatMemoWrappers {
			if !IsMemo(memo) {
				panic("BaseTransaction: invalid Memos. A memo can only have hexadecimals values. See https://xrpl.org/docs/references/protocol/transactions/common-fields#memos-field")
			}
		}
	}
}

func validateSigners(tx FlatTransaction) {
	if tx["Signers"] != nil {
		signers, ok := tx["Signers"].([]FlatTransaction)
		if !ok {
			panic("BaseTransaction: invalid Signers")
		}
		for _, signer := range signers {
			if !IsSigner(signer) {
				panic("BaseTransaction: invalid Signers")
			}
		}
	}
}
