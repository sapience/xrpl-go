package types

import (
	"errors"
	"fmt"
	"slices"

	"github.com/Peersyst/xrpl-go/xrpl/hash"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
)

// ErrBatchSignableInvalid is returned when the batch signable is invalid.
var ErrBatchSignableInvalid = errors.New("batch signable is invalid")

// BatchSignable contains the fields needed to perform a Batch transactions signature.
// It contains the Flags of all transactions in the batch and the IDs of the transactions.
type BatchSignable struct {
	Flags uint32
	TxIDs []string
}

// FromBatchTransaction creates a BatchSignable from a Batch transaction.
// It returns an error if the transaction is invalid.
func FromBatchTransaction(transaction *transaction.Batch) (*BatchSignable, error) {
	batchSignable := &BatchSignable{
		Flags: transaction.Flags,
		TxIDs: make([]string, len(transaction.RawTransactions)),
	}

	for i, rawTx := range transaction.RawTransactions {
		txID, err := hash.SignTx(rawTx.RawTransaction)
		if err != nil {
			return nil, fmt.Errorf("failed to get txID from raw transaction: %w", ErrBatchSignableInvalid)
		}
		batchSignable.TxIDs[i] = txID
	}

	return batchSignable, nil
}

// Equals checks if the BatchSignable is equal to another BatchSignable.
// It returns true if the flags and txIDs are equal, false otherwise.
func (b *BatchSignable) Equals(other *BatchSignable) bool {
	return b.Flags == other.Flags && slices.Equal(b.TxIDs, other.TxIDs)
}

// Flatten returns the BatchSignable as a map[string]interface{} for encoding.
func (b *BatchSignable) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})

	flattened["flags"] = b.Flags

	if len(b.TxIDs) > 0 {
		flattened["txIDs"] = b.TxIDs
	}

	return flattened
}
