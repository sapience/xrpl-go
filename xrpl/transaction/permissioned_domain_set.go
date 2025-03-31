package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// Maximum number of accepted credentials.
const MaxAcceptedCredentials int = 10

// PermissionedDomainSet represents a PermissionedDomainSet transaction.
type PermissionedDomainSet struct {
	BaseTx
	// DomainID is the ledger entry ID of an existing permissioned domain to modify.
	// When omitted, it creates a new permissioned domain.
	DomainID string `json:",omitempty"`
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
			entry := make(map[string]interface{})
			entry["Credential"] = cred.Flatten()
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

	// Use the custom credentials validation function.
	if err := types.ValidateCredentialsList(p.AcceptedCredentials, p.TxType().String(), MaxAcceptedCredentials); err != nil {
		return false, err
	}

	return true, nil
}
