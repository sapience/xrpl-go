package types

import "reflect"

// IsArray verifies the form and type of an Array at runtime.
func IsArray(input interface{}) bool {
	if input == nil {
		return false
	}
	val := reflect.ValueOf(input)
	return val.Kind() == reflect.Slice || val.Kind() == reflect.Array
}

// IsRecord verifies the form and type of a Record/Object at runtime.
func IsRecord(value interface{}) bool {
	if value == nil {
		return false
	}
	val := reflect.ValueOf(value)
	// Check if it's an object (map or struct) but not an array/slice
	return (val.Kind() == reflect.Map || val.Kind() == reflect.Struct) &&
		val.Kind() != reflect.Slice && val.Kind() != reflect.Array
}
