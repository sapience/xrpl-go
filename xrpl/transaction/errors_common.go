package transaction

import "errors"

var (
	// ErrInvalidAccount is returned when the Account field does not meet XRPL address standards.
	ErrInvalidAccount = errors.New("invalid xrpl address for Account")
	// ErrInvalidOwner is returned when the Owner field does not meet XRPL address standards.
	ErrInvalidOwner = errors.New("invalid xrpl address for Owner")
	// ErrInvalidDestination is returned when the Destination field does not meet XRPL address standards.
	ErrInvalidDestination = errors.New("invalid xrpl address for Destination")
	// ErrInvalidTransactionType is returned when the TransactionType field is invalid or missing.
	ErrInvalidTransactionType = errors.New("invalid or missing TransactionType")
	// ErrDestinationAccountConflict is returned when the Destination matches the Account.
	ErrDestinationAccountConflict = errors.New("Destination cannot be the same as the Account")
)
