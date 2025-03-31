package types

import (
	"errors"
)

// Maximum number of accepted credentials.
const MaxAcceptedCredentials int = 10

var (
	// Credential-specific errors
	ErrEmptyCredentials        = errors.New("credentials list cannot be empty")
	ErrInvalidCredentialCount  = errors.New("accepted credentials list must contain at least one and no more than the maximum allowed number of items")
	ErrDuplicateCredentials    = errors.New("credentials list cannot contain duplicate elements")
	ErrInvalidCredentialType   = errors.New("invalid credential type, must be a hexadecimal string between 1 and 64 bytes")
	ErrInvalidCredentialIssuer = errors.New("credential type: missing field Issuer")
)

// AuthorizeCredential represents an accepted credential for PermissionedDomainSet transactions.
type Credential struct {
	Issuer         Address
	CredentialType CredentialType
}

type AuthorizeCredential struct {
	Credential Credential
}

// Validate checks if the AuthorizeCredential is valid.
func (a AuthorizeCredential) Validate() error {
	if a.Credential.Issuer.String() == "" {
		return ErrInvalidCredentialIssuer
	}
	if !a.Credential.CredentialType.IsValid() {
		return ErrInvalidCredentialType
	}
	return nil
}

// Flatten returns a flattened map representation of the AuthorizeCredential.
func (a AuthorizeCredential) Flatten() map[string]interface{} {
	m := make(map[string]interface{})
	if a.Credential.Issuer.String() != "" {
		m["Issuer"] = a.Credential.Issuer.String()
	}
	if a.Credential.CredentialType != "" {
		m["CredentialType"] = a.Credential.CredentialType.String()
	}
	return m
}

// AuthorizeCredentialList represents a list of AuthorizeCredential.
type AuthorizeCredentialList []AuthorizeCredential

// Validate checks that the credential list:
// - is not empty,
// - does not exceed maxCredentials,
// - contains no duplicate credentials (based on Issuer + CredentialType),
// - and that each individual credential is valid.
func (list AuthorizeCredentialList) Validate() error {
	if len(list) == 0 {
		return ErrEmptyCredentials
	}
	if len(list) > MaxAcceptedCredentials {
		return ErrInvalidCredentialCount
	}
	seen := make(map[string]bool)
	for _, cred := range list {
		key := cred.Credential.Issuer.String() + cred.Credential.CredentialType.String()
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
