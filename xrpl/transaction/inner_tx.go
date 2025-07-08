package transaction

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

// RawTransactionWrapper represents the wrapper structure for transactions within a Batch.
type InnerTransaction struct {
	RawTransaction FlatTransaction `json:"RawTransaction"`
}

// Flatten returns the flattened map representation of the InnerTransaction.
func (i *InnerTransaction) Flatten() FlatTransaction {
	return map[string]any{
		"RawTransaction": map[string]any(i.RawTransaction),
	}
}

// Validate validates the InnerTransaction and its wrapped transaction.
func (i *InnerTransaction) Validate() (bool, error) {
	flattened := i.Flatten()

	// Validate that the flattened structure is a record
	if !types.IsRecord(flattened) {
		return false, ErrBatchRawTransactionNotObject
	}

	// Validate RawTransaction field exists
	rawTxField, exists := flattened["RawTransaction"]
	if !exists {
		return false, ErrBatchRawTransactionMissing
	}

	if !types.IsRecord(rawTxField) {
		return false, ErrBatchRawTransactionFieldNotObject
	}

	rawTx, ok := rawTxField.(map[string]interface{})
	if !ok {
		return false, ErrBatchRawTransactionFieldNotObject
	}

	// Check that TransactionType is not "Batch" (no nesting)
	if txType, ok := rawTx["TransactionType"].(string); ok && txType == "Batch" {
		return false, ErrBatchNestedTransaction
	}

	// Check for the TfInnerBatchTxn flag in the inner transactions
	if flags, ok := rawTx["Flags"].(uint32); !ok || !IsFlagEnabled(flags, TfInnerBatchTxn) {
		return false, ErrBatchMissingInnerFlag
	}

	// Fee must be "0" for inner transactions (or missing, which means 0)
	if feeField, exists := rawTx["Fee"]; exists {
		if feeStr, ok := feeField.(string); !ok || feeStr != "0" {
			return false, ErrBatchInnerTransactionInvalid
		}
	}

	// SigningPubKey must be empty for inner transactions (or missing, which means empty)
	if signingPubKeyField, exists := rawTx["SigningPubKey"]; exists {
		if signingPubKey, ok := signingPubKeyField.(string); !ok || signingPubKey != "" {
			return false, ErrBatchInnerTransactionInvalid
		}
	}

	// Check for disallowed fields in inner transactions
	if _, exists := rawTx["LastLedgerSequence"]; exists {
		return false, ErrBatchInnerTransactionInvalid
	}
	if _, exists := rawTx["Signers"]; exists {
		return false, ErrBatchInnerTransactionInvalid
	}
	if _, exists := rawTx["TxnSignature"]; exists {
		return false, ErrBatchInnerTransactionInvalid
	}

	return true, nil
}
