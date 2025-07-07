package transaction

import (
	"errors"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

const (
	// Batch transaction flags
	tfAllOrNothing uint32 = 0x00010000
	tfOnlyOne      uint32 = 0x00020000
	tfUntilFailure uint32 = 0x00040000
	tfIndependent  uint32 = 0x00080000
)

var (
	// General batch validation errors
	ErrBatchRawTransactionsEmpty = errors.New("RawTransactions must be a non-empty array")
	ErrBatchSignersNotArray      = errors.New("BatchSigners must be an array")

	// RawTransactions validation errors
	ErrBatchRawTransactionNotObject      = errors.New("batch RawTransaction element is not an object")
	ErrBatchRawTransactionMissing        = errors.New("batch RawTransaction field is missing")
	ErrBatchRawTransactionFieldNotObject = errors.New("batch RawTransaction field is not an object")
	ErrBatchNestedTransaction            = errors.New("batch cannot contain nested Batch transactions")
	ErrBatchMissingInnerFlag             = errors.New("batch RawTransaction must contain the TfInnerBatchTxn flag")

	// Inner transaction validation errors
	ErrBatchInnerTransactionInvalid = errors.New("batch inner transaction validation failed")

	// BatchSigners validation errors
	ErrBatchSignerNotObject        = errors.New("batch BatchSigner element is not an object")
	ErrBatchSignerMissing          = errors.New("batch BatchSigner field is missing")
	ErrBatchSignerFieldNotObject   = errors.New("batch BatchSigner field is not an object")
	ErrBatchSignerAccountMissing   = errors.New("batch BatchSigner Account is missing")
	ErrBatchSignerAccountNotString = errors.New("batch BatchSigner Account must be a string")
	ErrBatchSignerInvalid          = errors.New("batch signer validation failed")
)

// BatchSigner represents a single batch signer entry.
type BatchSigner struct {
	BatchSigner BatchSignerData `json:"BatchSigner"`
}

// BatchSignerData contains the actual batch signer information.
type BatchSignerData struct {
	Account       types.Address `json:"Account"`
	SigningPubKey string        `json:"SigningPubKey,omitempty"`
	TxnSignature  string        `json:"TxnSignature,omitempty"`
	Signers       []Signer      `json:"Signers,omitempty"`
}

// RawTransactionWrapper represents the wrapper structure for transactions within a Batch.
type RawTransactionWrapper struct {
	RawTransaction FlatTransaction `json:"RawTransaction"`
}

// Batch represents a Batch transaction that can execute multiple transactions atomically.
//
// Example:
//
// ```json
//
//	{
//	    "TransactionType": "Batch",
//	    "Account": "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
//	    "Fee": "100",
//	    "Flags": 65536,
//	    "Sequence": 1,
//	    "RawTransactions": [
//	        {
//	            "RawTransaction": {
//	                "TransactionType": "Payment",
//	                "Account": "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
//	                "Amount": "1000000",
//	                "Destination": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
//	                "Flags": 1073741824,
//	                "Fee": "0",
//	                "SigningPubKey": ""
//	            }
//	        }
//	    ]
//	}
//
// ```
type Batch struct {
	BaseTx
	// Array of transactions to be executed as part of this batch.
	RawTransactions []RawTransactionWrapper `json:"RawTransactions"`
	// Optional array of batch signers for multi-signing the batch.
	BatchSigners []BatchSigner `json:"BatchSigners,omitempty"`
}

// Flatten returns the flattened map of the BatchSigner.
func (bs *BatchSigner) Flatten() map[string]any {
	signer := map[string]any{
		"Account": bs.BatchSigner.Account.String(),
	}

	if bs.BatchSigner.SigningPubKey != "" {
		signer["SigningPubKey"] = bs.BatchSigner.SigningPubKey
	}
	if bs.BatchSigner.TxnSignature != "" {
		signer["TxnSignature"] = bs.BatchSigner.TxnSignature
	}
	if len(bs.BatchSigner.Signers) > 0 {
		innerSigners := make([]map[string]any, len(bs.BatchSigner.Signers))
		for i, innerSigner := range bs.BatchSigner.Signers {
			innerSigners[i] = innerSigner.Flatten()
		}
		signer["Signers"] = innerSigners
	}

	return map[string]any{
		"BatchSigner": signer,
	}
}

// TxType returns the type of the transaction (Batch).
func (*Batch) TxType() TxType {
	return BatchTx
}

// **********************************
// Batch Flags
// **********************************

// SetAllOrNothingFlag sets the AllOrNothing flag.
//
// AllOrNothing: Execute all transactions in the batch or none at all.
// If any transaction fails, the entire batch fails.
func (b *Batch) SetAllOrNothingFlag() {
	b.Flags |= tfAllOrNothing
}

// SetOnlyOneFlag sets the OnlyOne flag.
//
// OnlyOne: Execute only the first transaction that succeeds.
// Stop execution after the first successful transaction.
func (b *Batch) SetOnlyOneFlag() {
	b.Flags |= tfOnlyOne
}

// SetUntilFailureFlag sets the UntilFailure flag.
//
// UntilFailure: Execute transactions until one fails.
// Stop execution at the first failed transaction.
func (b *Batch) SetUntilFailureFlag() {
	b.Flags |= tfUntilFailure
}

// SetIndependentFlag sets the Independent flag.
//
// Independent: Execute all transactions independently.
// The failure of one transaction does not affect others.
func (b *Batch) SetIndependentFlag() {
	b.Flags |= tfIndependent
}

// Flatten returns the flattened map of the Batch transaction.
func (b *Batch) Flatten() FlatTransaction {
	out := b.BaseTx.Flatten()

	out["TransactionType"] = b.TxType().String()

	rawTxs := make([]map[string]any, len(b.RawTransactions))
	for i, rtw := range b.RawTransactions {
		rawTxs[i] = map[string]any{
			"RawTransaction": map[string]any(rtw.RawTransaction),
		}
	}
	out["RawTransactions"] = rawTxs

	if len(b.BatchSigners) > 0 {
		signers := make([]map[string]any, len(b.BatchSigners))
		for i, bs := range b.BatchSigners {
			signers[i] = bs.Flatten()
		}
		out["BatchSigners"] = signers
	}

	return out
}

// Validate validates the Batch transaction.
func (b *Batch) Validate() (bool, error) {
	_, err := b.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	flattenedTx := b.Flatten()

	if err := ValidateRequiredField(flattenedTx, "RawTransactions", isArray); err != nil {
		return false, err
	}

	// Validate RawTransactions array is not empty
	rawTxs, ok := flattenedTx["RawTransactions"].([]map[string]any)
	if !ok || len(rawTxs) == 0 {
		return false, ErrBatchRawTransactionsEmpty
	}

	// Validate each RawTransaction
	for _, rawTxObj := range rawTxs {
		if !isRecord(rawTxObj) {
			return false, ErrBatchRawTransactionNotObject
		}

		// Validate RawTransaction field exists
		rawTxField, exists := rawTxObj["RawTransaction"]
		if !exists {
			return false, ErrBatchRawTransactionMissing
		}

		if !isRecord(rawTxField) {
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

		// Fee must be "0" for inner transactions
		if feeStr, ok := rawTx["Fee"].(string); !ok || feeStr != "0" {
			return false, ErrBatchInnerTransactionInvalid
		}

		// SigningPubKey must be empty for inner transactions
		if signingPubKey, ok := rawTx["SigningPubKey"].(string); !ok || signingPubKey != "" {
			return false, ErrBatchInnerTransactionInvalid
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

	}

	if err := ValidateOptionalField(flattenedTx, "BatchSigners", isArray); err != nil {
		return false, err
	}

	// Validate BatchSigners if present
	if batchSignersField, exists := flattenedTx["BatchSigners"]; exists {
		batchSigners, ok := batchSignersField.([]map[string]any)
		if !ok {
			return false, ErrBatchSignersNotArray
		}

		for _, signerObj := range batchSigners {
			if !isRecord(signerObj) {
				return false, ErrBatchSignerNotObject
			}

			// Validate BatchSigner field exists
			batchSignerField, exists := signerObj["BatchSigner"]
			if !exists {
				return false, ErrBatchSignerMissing
			}

			if !isRecord(batchSignerField) {
				return false, ErrBatchSignerFieldNotObject
			}

			signer, ok := batchSignerField.(map[string]interface{})
			if !ok {
				return false, ErrBatchSignerFieldNotObject
			}

			// Validate required Account field
			if account, exists := signer["Account"]; !exists {
				return false, ErrBatchSignerAccountMissing
			} else if accountStr, ok := account.(string); !ok {
				return false, ErrBatchSignerAccountNotString
			} else if accountStr == "" {
				return false, ErrBatchSignerInvalid
			}

			// Validate optional SigningPubKey field
			if signingPubKey, exists := signer["SigningPubKey"]; exists {
				if _, ok := signingPubKey.(string); !ok {
					return false, ErrBatchSignerInvalid
				}
			}

			// Validate optional TxnSignature field
			if txnSignature, exists := signer["TxnSignature"]; exists {
				if _, ok := txnSignature.(string); !ok {
					return false, ErrBatchSignerInvalid
				}
			}

			// Validate optional Signers field
			if signersField, exists := signer["Signers"]; exists {
				if !isArray(signersField) {
					return false, ErrBatchSignerInvalid
				}
			}
		}
	}

	return true, nil
}
