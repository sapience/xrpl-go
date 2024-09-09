package utils

import "testing"

func TestOnlyHasFields(t *testing.T) {
	t.Run("Object has only specified fields", func(t *testing.T) {
		obj := map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
			"field3": "value3",
		}
		fields := []string{"field1", "field2", "field3"}
		if !OnlyHasFields(obj, fields) {
			t.Errorf("Expected OnlyHasFields to return true, but got false")
		}
	})

	t.Run("Object has a subset of specified fields", func(t *testing.T) {
		obj := map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		}
		fields := []string{"field1", "field2", "field3"}
		if !OnlyHasFields(obj, fields) {
			t.Errorf("Expected OnlyHasFields to return true, but got false")
		}
	})

	t.Run("Object has extra fields", func(t *testing.T) {
		obj := map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
			"field3": "value3",
			"field4": "value4",
		}
		fields := []string{"field1", "field2", "field3"}
		if OnlyHasFields(obj, fields) {
			t.Errorf("Expected OnlyHasFields to return false, but got true")
		}
	})

	t.Run("Object has no fields", func(t *testing.T) {
		obj := map[string]interface{}{}
		fields := []string{"field1", "field2", "field3"}
		if OnlyHasFields(obj, fields) {
			t.Errorf("Expected OnlyHasFields to return false, but got true")
		}
	})

	t.Run("Empty object and empty fields", func(t *testing.T) {
		obj := map[string]interface{}{}
		fields := []string{}
		if !OnlyHasFields(obj, fields) {
			t.Errorf("Expected OnlyHasFields to return true, but got false")
		}
	})
}
