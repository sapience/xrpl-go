package types

// IsArray verifies the form and type of an Array at runtime.
func IsArray(input interface{}) bool {
	if input == nil {
		return false
	}
	// Check for []map[string]any - the type used for all our array fields
	_, ok := input.([]map[string]any)
	return ok
}

// IsRecord verifies the form and type of a Record/Object at runtime.
func IsRecord(value interface{}) bool {
	if value == nil {
		return false
	}
	// Check for map[string]any (which is the same as map[string]interface{})
	_, ok := value.(map[string]any)
	return ok
}
