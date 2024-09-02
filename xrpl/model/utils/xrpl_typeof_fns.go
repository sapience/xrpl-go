package utils

import (
	objectfns "github.com/Peersyst/xrpl-go/xrpl/utils/object-fns"
	typeoffns "github.com/Peersyst/xrpl-go/xrpl/utils/typeof-fns"
)

const MEMO_SIZE = 3

// IsMemo checks if the given object is a valid Memo object.
func IsMemo(obj map[string]interface{}) bool {
	if obj == nil {
		return false
	}

	memo, ok := obj["Memo"].(map[string]interface{})
	if !ok {
		return false
	}

	size := len(objectfns.GetKeys(memo))
	validData := memo["MemoData"] == nil || typeoffns.IsString(memo["MemoData"])
	validFormat := memo["MemoFormat"] == nil || typeoffns.IsString(memo["MemoFormat"])
	validType := memo["MemoType"] == nil || typeoffns.IsString(memo["MemoType"])

	return size >= 1 && size <= MEMO_SIZE && validData && validFormat && validType && onlyHasFields(memo, []string{"MemoFormat", "MemoData", "MemoType"})
}

// onlyHasFields checks if the given object has only the specified fields or a subset of them.
func onlyHasFields(obj map[string]interface{}, fields []string) bool {
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
