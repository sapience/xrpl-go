package types

import "errors"

var (
	// General batch validation errors

	// ErrBatchRawTransactionsEmpty is returned when the RawTransactions array is empty or nil.
	ErrBatchRawTransactionsEmpty = errors.New("RawTransactions must be a non-empty array")

	// ErrBatchSignersNotArray is returned when BatchSigners field is present but not an array type.
	ErrBatchSignersNotArray = errors.New("BatchSigners must be an array")

	// RawTransactions validation errors

	// ErrBatchRawTransactionNotObject is returned when a RawTransaction array element is not an object.
	ErrBatchRawTransactionNotObject = errors.New("batch RawTransaction element is not an object")

	// ErrBatchRawTransactionMissing is returned when the RawTransaction field is missing from an array element.
	ErrBatchRawTransactionMissing = errors.New("batch RawTransaction field is missing")

	// ErrBatchRawTransactionFieldNotObject is returned when the RawTransaction field is not an object.
	ErrBatchRawTransactionFieldNotObject = errors.New("batch RawTransaction field is not an object")

	// ErrBatchNestedTransaction is returned when trying to include a Batch transaction within another Batch.
	ErrBatchNestedTransaction = errors.New("batch cannot contain nested Batch transactions")

	// ErrBatchMissingInnerFlag is returned when an inner transaction lacks the TfInnerBatchTxn flag.
	ErrBatchMissingInnerFlag = errors.New("batch RawTransaction must contain the TfInnerBatchTxn flag")

	// Inner transaction validation errors

	// ErrBatchInnerTransactionInvalid is returned when an inner transaction fails its own validation.
	ErrBatchInnerTransactionInvalid = errors.New("batch inner transaction validation failed")

	// BatchSigners validation errors

	// ErrBatchSignerNotObject is returned when a BatchSigner array element is not an object.
	ErrBatchSignerNotObject = errors.New("batch BatchSigner element is not an object")

	// ErrBatchSignerMissing is returned when the BatchSigner field is missing from an array element.
	ErrBatchSignerMissing = errors.New("batch BatchSigner field is missing")

	// ErrBatchSignerFieldNotObject is returned when the BatchSigner field is not an object.
	ErrBatchSignerFieldNotObject = errors.New("batch BatchSigner field is not an object")

	// ErrBatchSignerAccountMissing is returned when a BatchSigner lacks the required Account field.
	ErrBatchSignerAccountMissing = errors.New("batch BatchSigner Account is missing")

	// ErrBatchSignerAccountNotString is returned when a BatchSigner Account field is not a string.
	ErrBatchSignerAccountNotString = errors.New("batch BatchSigner Account must be a string")

	// ErrBatchSignerInvalid is returned when a BatchSigner fails its validation rules.
	ErrBatchSignerInvalid = errors.New("batch signer validation failed")
)

// BatchSigner represents a single batch signer entry.
type BatchSigner struct {
	BatchSigner BatchSignerData `json:"BatchSigner"`
}

// BatchSignerData contains the actual batch signer information.
type BatchSignerData struct {
	Account       Address  `json:"Account"`
	SigningPubKey string   `json:"SigningPubKey,omitempty"`
	TxnSignature  string   `json:"TxnSignature,omitempty"`
	Signers       []Signer `json:"Signers,omitempty"`
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

func (bs *BatchSigner) Validate() error {
	// Validate that BatchSigner exists
	if bs == nil {
		return ErrBatchSignerInvalid
	}

	// Convert the single BatchSigner to the same format that was being validated
	flattened := bs.Flatten()

	// Validate BatchSigner field exists
	batchSignerField, exists := flattened["BatchSigner"]
	if !exists {
		return ErrBatchSignerMissing
	}

	if !IsRecord(batchSignerField) {
		return ErrBatchSignerFieldNotObject
	}

	signer, ok := batchSignerField.(map[string]interface{})
	if !ok {
		return ErrBatchSignerFieldNotObject
	}

	// Validate required Account field
	if account, exists := signer["Account"]; !exists {
		return ErrBatchSignerAccountMissing
	} else if accountStr, ok := account.(string); !ok {
		return ErrBatchSignerAccountNotString
	} else if accountStr == "" {
		return ErrBatchSignerInvalid
	}

	// Validate optional SigningPubKey field
	if signingPubKey, exists := signer["SigningPubKey"]; exists {
		if _, ok := signingPubKey.(string); !ok {
			return ErrBatchSignerInvalid
		}
	}

	// Validate optional TxnSignature field
	if txnSignature, exists := signer["TxnSignature"]; exists {
		if _, ok := txnSignature.(string); !ok {
			return ErrBatchSignerInvalid
		}
	}

	// Validate optional Signers field
	if signersField, exists := signer["Signers"]; exists {
		if !IsArray(signersField) {
			return ErrBatchSignerInvalid
		}
	}

	return nil
}
