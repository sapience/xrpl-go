package types

const (
	// TfInnerBatchTxn flag that must be set on inner transactions within a batch
	TfInnerBatchTxn uint32 = 0x40000000
)

// RawTransactionWrapper represents the wrapper structure for transactions within a Batch.
type RawTransaction struct {
	RawTransaction map[string]any `json:"RawTransaction"`
}

// Flatten returns the flattened map representation of the RawTransaction.
func (r *RawTransaction) Flatten() map[string]any {
	return r.RawTransaction
}

// Validate validates the RawTransaction and its wrapped transaction.
func (r *RawTransaction) Validate() (bool, error) {
	flattened := r.Flatten()

	// Validate that the flattened structure is a record
	if !IsTransactionArray(flattened) {
		return false, ErrBatchRawTransactionNotObject
	}

	// Validate RawTransaction field exists
	rawTxField, exists := flattened["RawTransaction"]
	if !exists {
		return false, ErrBatchRawTransactionMissing
	}

	if !IsTransactionArray(rawTxField) {
		return false, ErrBatchRawTransactionFieldNotObject
	}

	rawTx, ok := rawTxField.(map[string]any)
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
