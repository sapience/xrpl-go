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

// String returns the string representation of a CredentialType.
func (c *CredentialType) String() string {
	return string(*c)
}

// IsValidCredentialType checks if a credential type meets all the requirements:
// - Not empty
// - Valid hex string
// - Length between MinCredentialTypeLength and MaxCredentialTypeLength
func (c *CredentialType) IsValid() bool {
	if c.String() == "" {
		return false
	}

	credTypeStr := c.String()
	if !typecheck.IsHex(credTypeStr) {
		return false
	}

	length := len(credTypeStr)
	return length >= MinCredentialTypeLength && length <= MaxCredentialTypeLength
}

// AuthorizeCredential represents an accepted credential for PermissionedDomainSet transactions.
type AuthorizeCredential struct {
	Issuer         string         `json:"Issuer"`
	CredentialType CredentialType `json:"CredentialType"`
}

// Validate checks if the AuthorizeCredential is valid.
func (ac *AuthorizeCredential) Validate() error {
	if ac.Issuer == "" {
		return ErrInvalidCredentialIssuer
	}
	if !ac.CredentialType.IsValid() {
		return ErrInvalidCredentialType
	}
	return nil
}

// Flatten returns a flattened map representation of the AuthorizeCredential.
func (ac *AuthorizeCredential) Flatten() interface{} {
	json := make(map[string]interface{})
	if ac.Issuer != "" {
		json["Issuer"] = ac.Issuer
	}
	if ac.CredentialType != "" {
		json["CredentialType"] = ac.CredentialType.String()
	}
	return json
}

// validateCredentialsList validates a list of AuthorizeCredential objects.
// It checks that:
// - The credentials slice is not nil.
// - The length is between 1 and maxCredentials.
// - Each credential is valid.
// - There are no duplicate credentials.
func ValidateCredentialsList(credentials []AuthorizeCredential, transactionType string, maxCredentials int) error {

	if len(credentials) == 0 {
		return ErrEmptyCredentials
	}

	if len(credentials) > maxCredentials {
		return ErrInvalidCredentialCount
	}

	seen := make(map[string]bool)
	for _, cred := range credentials {
		key := cred.Issuer + cred.CredentialType.String()
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
