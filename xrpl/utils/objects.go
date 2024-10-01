package utils

// OnlyHasFields checks if the given object has only the specified fields or a subset of them.
func ObjectOnlyHasFields(obj map[string]interface{}, fields []string) bool {
	// if obj is nil or obj is empty but not fields, return false
	if obj == nil || (len(obj) == 0 && len(fields) > 0) {
		return false
	}

	// if fields is empty, return true
	if len(fields) == 0 {
		return true
	}

	fieldSet := make(map[string]struct{}, len(fields))
	for _, field := range fields {
		fieldSet[field] = struct{}{}
	}

	for key := range obj {
		if _, ok := fieldSet[key]; !ok {
			return false
		}
	}
	return true
}
