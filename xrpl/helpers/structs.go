package helpers

import "reflect"

// Function to convert struct to map[string]any
func StructToMap(s interface{}) map[string]any {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("Input is not a struct")
	}

	m := make(map[string]any)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		m[field.Name] = value
	}
	return m
}
