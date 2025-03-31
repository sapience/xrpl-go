package types

import (
	"errors"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
)

var (
	// Credential-specific errors
	ErrEmptyCredentials        = errors.New("credentials list cannot be empty")
	ErrInvalidCredentialCount  = errors.New("accepted credentials list must contain at least one and no more than the maximum allowed number of items")
	ErrDuplicateCredentials    = errors.New("credentials list cannot contain duplicate elements")
	ErrInvalidCredentialType   = errors.New("invalid credential type, must be a hexadecimal string between 1 and 64 bytes")
	ErrInvalidCredentialIssuer = errors.New("credential type: missing field Issuer")
)

const (
	// Minimum length of a credential type is 1 byte (1 byte = 2 hex characters).
	MinCredentialTypeLength = 2

	// Maximum length of a credential type is 64 bytes (1 byte = 2 hex characters).
	MaxCredentialTypeLength = 128
)

type CredentialType string

func (c *CredentialType) String() string {
	return string(*c)
}

// IsValidCredentialType checks if a credential type meets all the requirements:
// - Not empty
// - Valid hex string
// - Length between MinCredentialTypeLength and MaxCredentialTypeLength
func (c CredentialType) IsValid() bool {
	hexStr := c.String()
	if hexStr == "" {
		return false
	}
	if !typecheck.IsHex(hexStr) {
		return false
	}
	length := len(hexStr)
	return length >= MinCredentialTypeLength && length <= MaxCredentialTypeLength
}

// AuthorizeCredential represents an accepted credential for PermissionedDomainSet transactions.
type AuthorizeCredential struct {
	Issuer         Address
	CredentialType CredentialType `json:"CredentialType"`
}

// Validate checks if the AuthorizeCredential is valid.
func (a AuthorizeCredential) Validate() error {
	if a.Issuer.String() == "" {
		return ErrInvalidCredentialIssuer
	}
	if !a.CredentialType.IsValid() {
		return ErrInvalidCredentialType
	}
	return nil
}

// Flatten returns a flattened map representation of the AuthorizeCredential.
func (a AuthorizeCredential) Flatten() map[string]interface{} {
	m := make(map[string]interface{})
	if a.Issuer.String() != "" {
		m["Issuer"] = a.Issuer.String()
	}
	if a.CredentialType != "" {
		m["CredentialType"] = a.CredentialType.String()
	}
	return m
}

// validateCredentialsList validates a list of AuthorizeCredential objects.
// It checks that:
// - The credentials slice is not nil.
// - The length is between 1 and maxCredentials.
// - Each credential is valid.
// - There are no duplicate credentials.
func ValidateCredentialsList(credentials []AuthorizeCredential, maxCredentials int) error {

	if len(credentials) == 0 {
		return ErrEmptyCredentials
	}

	if len(credentials) > maxCredentials {
		return ErrInvalidCredentialCount
	}

	seen := make(map[string]bool)
	for _, cred := range credentials {
		key := cred.Issuer.String() + cred.CredentialType.String()
		if seen[key] {
			return ErrDuplicateCredentials
		}
		seen[key] = true

		if err := cred.Validate(); err != nil {
			return err
		}
	}
	return nil

}
