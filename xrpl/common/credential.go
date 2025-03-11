package common

import (
	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// IsValidCredentialType checks if a credential type meets all the requirements:
// - Not empty
// - Valid hex string
// - Length between MinCredentialTypeLength and MaxCredentialTypeLength
func IsValidCredentialType(credType types.CredentialType) bool {
	if credType == "" {
		return false
	}

	credTypeStr := credType.String()
	if !typecheck.IsHex(credTypeStr) {
		return false
	}

	length := len(credTypeStr)
	return length >= MinCredentialTypeLength && length <= MaxCredentialTypeLength
}
