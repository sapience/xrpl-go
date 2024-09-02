package utils

import (
	"testing"
)

func TestIsMemo(t *testing.T) {
	t.Run("Valid Memo object", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"Memo": map[string]interface{}{
				"MemoData":   "Hello World",
				"MemoFormat": "text/plain",
				"MemoType":   "general",
			},
		}
		if !IsMemo(obj1) {
			t.Errorf("Expected IsMemo to return true, but got false")
		}
	})

	t.Run("Memo object with missing fields", func(t *testing.T) {
		obj2 := map[string]interface{}{
			"Memo": map[string]interface{}{
				"MemoData": "Hello World",
			},
		}
		if !IsMemo(obj2) {
			t.Errorf("Expected IsMemo to return true, but got false")
		}
	})

	t.Run("Memo object with invalid field types", func(t *testing.T) {
		obj3 := map[string]interface{}{
			"Memo": map[string]interface{}{
				"MemoData":   12345,
				"MemoFormat": 12345,
				"MemoType":   12345,
			},
		}
		if IsMemo(obj3) {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})

	t.Run("Memo object with extra fields", func(t *testing.T) {
		obj4 := map[string]interface{}{
			"Memo": map[string]interface{}{
				"MemoData":   "Hello World",
				"MemoFormat": "text/plain",
				"MemoType":   "general",
				"ExtraField": "Extra Value",
			},
		}
		if IsMemo(obj4) {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})

	t.Run("Nil object", func(t *testing.T) {
		obj5 := map[string]interface{}{}
		if IsMemo(obj5) {
			t.Errorf("Expected IsMemo to return false, but got true")
		}
	})
}
