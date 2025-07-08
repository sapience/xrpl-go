package types

// IsTransactionArray verifies that the input is a valid transaction array format.
// This checks for []map[string]any - the type used for serialized transaction arrays
// like RawTransactions, Memos, Signers, etc.
func IsTransactionArray(input interface{}) bool {
	if input == nil {
		return false
	}
	// Check for []map[string]any - the type used for all our array fields
	_, ok := input.([]map[string]any)
	return ok
}

// IsTransactionObject verifies that the input is a valid transaction object format.
// This checks for map[string]any - the type used for serialized transaction objects
// and ensures it's not an array.
func IsTransactionObject(value interface{}) bool {
	if value == nil {
		return false
	}
	// Explicitly check that it's not an array (matching xrpl.js behavior)
	if IsTransactionArray(value) {
		return false
	}
	// Check for map[string]any (which is the same as map[string]interface{})
	_, ok := value.(map[string]any)
	return ok
}
