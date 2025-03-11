package transaction

import (
	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl/common"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

type CredentialDelete struct {
	BaseTx

	// Arbitrary data defining the type of credential to delete. The minimum length is 1 byte and the maximum length is 256 bytes.
	CredentialType types.CredentialType

	// The subject of the credential to delete. If omitted, use the Account (sender of the transaction) as the subject of the credential.
	Subject types.Address `json:",omitempty"`

	// The issuer of the credential to delete. If omitted, use the Account (sender of the transaction) as the issuer of the credential.
	Issuer types.Address `json:",omitempty"`
}

func (*CredentialDelete) TxType() TxType {
	return CredentialDeleteTx
}

func (c *CredentialDelete) Flatten() FlatTransaction {
	flattened := c.BaseTx.Flatten()

	flattened["TransactionType"] = c.TxType().String()
	flattened["CredentialType"] = c.CredentialType.String()

	if c.Subject != "" {
		flattened["Subject"] = c.Subject.String()
	}

	if c.Issuer != "" {
		flattened["Issuer"] = c.Issuer.String()
	}

	return flattened
}

func (c *CredentialDelete) Validate() (bool, error) {
	// validate the base transaction
	_, err := c.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if !common.IsValidCredentialType(c.CredentialType) {
		return false, ErrInvalidCredentialType
	}

	if c.Subject != "" && !addresscodec.IsValidAddress(c.Subject.String()) {
		return false, ErrInvalidSubject
	}

	if c.Issuer != "" && !addresscodec.IsValidAddress(c.Issuer.String()) {
		return false, ErrInvalidIssuer
	}

	return true, nil
}
