package types

// BatchSignable contains the fields needed to perform a Batch transactions signature.
// It contains the Flags of all transactions in the batch and the IDs of the transactions.
type BatchSignable struct {
	Flags uint32
	TxIDs []string
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
