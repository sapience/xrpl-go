package transaction

import "errors"

var (
	// errInvalidOwnerAddress is returned when the owner address is invalid.
	errInvalidOwnerAddress = errors.New("invalid xrpl address for the Owner field")
	// errInvalidDestinationAddress is returned when the destination address is invalid.
	errInvalidDestinationAddress = errors.New("invalid xrpl address for the Destination field")
	// ErrDestinationAccountConflict is returned when the destination is the same as the account field.
	errDestinationAccountConflict = errors.New("destination must be different from the account")
)
