package transaction

import (
	"errors"
)

var (
	// Credential-specific errors
	ErrMissingDomainID = errors.New("missing required field: DomainID")
)

// PermissionedDomainDelete represents a PermissionedDomainDelete transaction.
type PermissionedDomainDelete struct {
	BaseTx
	// DomainID is the ledger entry ID of the Permissioned Domain entry to delete.
	DomainID string
}

// TxType returns the transaction type.
func (p *PermissionedDomainDelete) TxType() TxType {
	return PermissionedDomainDeleteTx
}

// Flatten returns a flattened map representation of the PermissionedDomainDelete transaction.
func (p *PermissionedDomainDelete) Flatten() FlatTransaction {
	flattened := p.BaseTx.Flatten()
	flattened["TransactionType"] = p.TxType().String()
	flattened["DomainID"] = p.DomainID
	return flattened
}

// Validate validates the PermissionedDomainDelete transaction.
// It ensures that the base transaction is valid and that the required DomainID field is present.
func (p *PermissionedDomainDelete) Validate() (bool, error) {
	// Validate common transaction fields.
	if ok, err := p.BaseTx.Validate(); !ok {
		return false, err
	}
	// Ensure DomainID is provided.
	if p.DomainID == "" {
		return false, errors.New("missing required field: DomainID")
	}
	return true, nil
}
