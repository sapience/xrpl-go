package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
)

// The SignerListSet transaction creates, replaces, or removes a list of signers that can be used to multi-sign a transaction. This transaction type was introduced by the MultiSign amendment.
//
// Example:
//
// ```json
//
//	{
//	    "Flags": 0,
//	    "TransactionType": "SignerListSet",
//	    "Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
//	    "Fee": "12",
//	    "SignerQuorum": 3,
//	    "SignerEntries": [
//	        {
//	            "SignerEntry": {
//	                "Account": "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
//	                "SignerWeight": 2
//	            }
//	        },
//	        {
//	            "SignerEntry": {
//	                "Account": "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
//	                "SignerWeight": 1
//	            }
//	        },
//	        {
//	            "SignerEntry": {
//	                "Account": "raKEEVSGnKSD9Zyvxu4z6Pqpm4ABH8FS6n",
//	                "SignerWeight": 1
//	            }
//	        }
//	    ]
//	}
//
// `
type SignerListSet struct {
	BaseTx
	// A target number for the signer weights. A multi-signature from this list is valid only if the sum weights of the signatures provided is greater than or equal to this value.
	// To delete a signer list, use the value 0.
	SignerQuorum uint
	// (Omitted when deleting) Array of SignerEntry objects, indicating the addresses and weights of signers in this list.
	// This signer list must have at least 1 member and no more than 32 members.
	// No address may appear more than once in the list, nor may the Account submitting the transaction appear in the list.
	SignerEntries []ledger.SignerEntryWrapper
}

// TxType returns the transaction type for this transaction (SignerListSet).
func (*SignerListSet) TxType() TxType {
	return SignerListSetTx
}

// Flatten returns the flattened map of the SignerListSet transaction.
func (s *SignerListSet) Flatten() FlatTransaction {
	flattened := s.BaseTx.Flatten()

	flattened["TransactionType"] = "SignerListSet"
	flattened["SignerQuorum"] = s.SignerQuorum

	if len(s.SignerEntries) > 0 {
		flattedSignerListEntries := make([]map[string]interface{}, 0)
		for _, signerEntry := range s.SignerEntries {
			flattenedEntry := signerEntry.Flatten()
			if flattenedEntry != nil {
				flattedSignerListEntries = append(flattedSignerListEntries, flattenedEntry)
			}
		}
		flattened["SignerEntries"] = flattedSignerListEntries
	}

	return flattened
}
