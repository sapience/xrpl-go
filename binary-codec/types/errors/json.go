package errors

// ErrNotValidJSON is an error that occurs when the json is not valid.
type ErrNotValidJSON struct{}

// Error returns the error message for the ErrNotValidJson error.
func (e *ErrNotValidJSON) Error() string {
	return "not a valid json"
}
