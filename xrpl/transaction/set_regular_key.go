package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// A SetRegularKey transaction assigns, changes, or removes the regular key pair associated with an account.
//
// You can protect your account by assigning a regular key pair to it and using it instead of the master key pair to sign transactions whenever possible.
// If your regular key pair is compromised, but your master key pair is not, you can use a SetRegularKey transaction to regain control of your account.
//
// Example:
//
// ```json
//
//	{
//	    "Flags": 0,
//	    "TransactionType": "SetRegularKey",
//	    "Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
//	    "Fee": "12",
//	    "RegularKey": "rAR8rR8sUkBoCZFawhkWzY4Y5YoyuznwD"
//	}
//
// ```
type SetRegularKey struct {
	BaseTx
	// (Optional) A base-58-encoded Address that indicates the regular key pair to be assigned to the account.
	// If omitted, removes any existing regular key pair from the account. Must not match the master key pair for the address.
	RegularKey types.Address `json:",omitempty"`
}

// TxType returns the transaction type for this transaction (SetRegularKey).
func (*SetRegularKey) TxType() TxType {
	return SetRegularKeyTx
}

// Flatten returns the flattened map of the SetRegularKey transaction.
func (s *SetRegularKey) Flatten() FlatTransaction {
	flattened := s.BaseTx.Flatten()

	flattened["TransactionType"] = "SetRegularKey"

	if s.RegularKey != "" {
		flattened["RegularKey"] = s.RegularKey
	}

	return flattened
}
