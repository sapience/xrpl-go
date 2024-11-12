package transaction

import "errors"

var (
	// ErrDestinationAccountConflict is returned when the Destination matches the Account.
	ErrDestinationAccountConflict = errors.New("Destination cannot be the same as the Account")
	// ErrInvalidAccount is returned when the Account field does not meet XRPL address standards.
	ErrInvalidAccount = errors.New("invalid xrpl address for Account")
	// ErrInvalidDestination is returned when the Destination field does not meet XRPL address standards.
	ErrInvalidDestination = errors.New("invalid xrpl address for Destination")
	// ErrInvalidOwner is returned when the Owner field does not meet XRPL address standards.
	ErrInvalidOwner = errors.New("invalid xrpl address for Owner")
	// ErrInvalidTransactionType is returned when the TransactionType field is invalid or missing.
	ErrInvalidTransactionType = errors.New("invalid or missing TransactionType")
)
