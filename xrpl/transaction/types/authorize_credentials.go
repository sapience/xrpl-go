package types

import (
	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
)

type AuthorizeCredentials struct {
	// The issuer of the credential.
	Issuer Address
	// The credential type of the credential.
	CredentialType CredentialType
}

// IsValid returns true if the authorize credentials are valid.
func (a *AuthorizeCredentials) IsValid() bool {
	return addresscodec.IsValidAddress(a.Issuer.String()) && a.CredentialType.IsValid()
}

// Flatten returns a map of the authorize credentials.
func (a *AuthorizeCredentials) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})

	flattened["Issuer"] = a.Issuer.String()
	flattened["CredentialType"] = a.CredentialType.String()

	return flattened
}
