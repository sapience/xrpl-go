package common

import "time"

const (
	// Ledger constants
	LedgerOffset uint32 = 20

	// Config constants
	DefaultHost                  = "localhost"
	DefaultMaxRetries            = 10
	DefaultMaxReconnects         = 3
	DefaultRetryDelay            = 1 * time.Second
	DefaultFeeCushion    float32 = 1.2
	DefaultMaxFeeXRP     float32 = 2

	// 5 seconds default timeout
	DefaultTimeout = 5 * time.Second

	// Transaction constants
	// Minimum length of a credential type is 1 byte (1 byte = 2 hex characters).
	MinCredentialTypeLength = 2
	// Maximum length of a credential type is 64 bytes (1 byte = 2 hex characters).
	MaxCredentialTypeLength = 128
)
