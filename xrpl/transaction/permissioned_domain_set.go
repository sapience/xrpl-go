package transaction

import (
	"errors"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

var (
	// Credential-specific errors
	ErrEmptyCredentials       = errors.New("credentials list cannot be empty")
	ErrInvalidCredentialCount = errors.New("accepted credentials list must contain at least one and no more than the maximum allowed number of items")
	ErrDuplicateCredentials   = errors.New("credentials list cannot contain duplicate elements")
)

// Maximum number of accepted credentials.
const MaxAcceptedCredentials int = 10

// PermissionedDomainSet represents a PermissionedDomainSet transaction.
type PermissionedDomainSet struct {
	BaseTx
	// DomainID is the ledger entry ID of an existing permissioned domain to modify.
	// When omitted, it creates a new permissioned domain.
	DomainID string `json:"DomainID,omitempty"`
	// AcceptedCredentials is a list of credentials that grant access to the domain.
	// An empty array indicates deletion of the field.
	AcceptedCredentials []types.AuthorizeCredential
}

// TxType returns the type of the transaction.
func (p *PermissionedDomainSet) TxType() TxType {
	return PermissionedDomainSetTx
}

// Flatten returns a flattened map representation of the PermissionedDomainSet transaction.
func (p *PermissionedDomainSet) Flatten() FlatTransaction {
	flattened := p.BaseTx.Flatten()
	flattened["TransactionType"] = p.TxType().String()

	if p.DomainID != "" {
		flattened["DomainID"] = p.DomainID
	}

	if len(p.AcceptedCredentials) > 0 {
		credentials := make([]interface{}, len(p.AcceptedCredentials))
		for i, cred := range p.AcceptedCredentials {
			// Inline flattening for each credential.
			credMap := make(map[string]interface{})
			if cred.Issuer != "" {
				credMap["Issuer"] = cred.Issuer
			}
			if cred.CredentialType != "" {
				credMap["CredentialType"] = cred.CredentialType.String()
			}
			entry := map[string]interface{}{
				"Credential": credMap,
			}
			credentials[i] = entry
		}
		flattened["AcceptedCredentials"] = credentials
	}

	return flattened
}

func (p *PermissionedDomainSet) Validate() (bool, error) {
	if ok, err := p.BaseTx.Validate(); !ok {
		return false, err
	}

	// Check that the credentials list is not empty.
	if len(p.AcceptedCredentials) == 0 {
		return false, ErrEmptyCredentials
	}
	// Check that the number of credentials does not exceed the maximum allowed.
	if len(p.AcceptedCredentials) > MaxAcceptedCredentials {
		return false, ErrInvalidCredentialCount
	}

	// Validate each credential and check for duplicates.
	seen := make(map[string]bool)
	for _, cred := range p.AcceptedCredentials {
		// Create a unique key based on Issuer and CredentialType.
		key := cred.Issuer + cred.CredentialType.String()
		if seen[key] {
			return false, ErrDuplicateCredentials
		}
		seen[key] = true

		// Inline validation for each credential.
		if cred.Issuer == "" {
			return false, ErrInvalidIssuer
		}
		if !cred.CredentialType.IsValid() {
			return false, ErrInvalidCredentialType
		}
	}

	return true, nil
}
