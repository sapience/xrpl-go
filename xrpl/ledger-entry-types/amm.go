package ledger

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// An AMM ledger entry describes a single Automated Market Maker (AMM) instance.
// This is always paired with a special AccountRoot entry. https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/amm#amm
//
// Example:
//
//	{
//	    "Account" : "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
//	    "Asset" : {
//	      "currency" : "XRP"
//	    },
//	    "Asset2" : {
//	      "currency" : "TST",
//	      "issuer" : "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd"
//	    },
//	    "AuctionSlot" : {
//	      "Account" : "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
//	      "AuthAccounts" : [
//	          {
//	            "AuthAccount" : {
//	                "Account" : "rMKXGCbJ5d8LbrqthdG46q3f969MVK2Qeg"
//	            }
//	          },
//	          {
//	            "AuthAccount" : {
//	                "Account" : "rBepJuTLFJt3WmtLXYAxSjtBWAeQxVbncv"
//	            }
//	          }
//	      ],
//	      "DiscountedFee" : 60,
//	      "Expiration" : 721870180,
//	      "Price" : {
//	          "currency" : "039C99CD9AB0B70B32ECDA51EAAE471625608EA2",
//	          "issuer" : "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
//	          "value" : "0.8696263565463045"
//	      }
//	    },
//	    "Flags" : 0,
//	    "LPTokenBalance" : {
//	      "currency" : "039C99CD9AB0B70B32ECDA51EAAE471625608EA2",
//	      "issuer" : "rE54zDvgnghAoPopCgvtiqWNq3dU5y836S",
//	      "value" : "71150.53584131501"
//	    },
//	    "TradingFee" : 600,
//	    "VoteSlots" : [
//	      {
//	          "VoteEntry" : {
//	            "Account" : "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
//	            "TradingFee" : 600,
//	            "VoteWeight" : 100000
//	          }
//	      }
//	    ]
//	}
type AMM struct {
	LedgerEntryCommonFields
	// The address of the special account that holds this AMM's assets.
	Account types.Address
	// The definition for one of the two assets this AMM holds. In JSON, this is an object with currency and issuer fields.
	Asset Asset
	// The definition for the other asset this AMM holds. In JSON, this is an object with currency and issuer fields.
	Asset2 Asset
	// Details of the current owner of the auction slot, as an Auction Slot object.
	AuctionSlot AuctionSlot `json:",omitempty"`
	// The total outstanding balance of liquidity provider tokens from this AMM instance.
	// The holders of these tokens can vote on the AMM's trading fee in proportion to their holdings, or redeem the tokens for a share of the AMM's assets which grows with the trading fees collected.
	LPTokenBalance types.CurrencyAmount
	// The percentage fee to be charged for trades against this AMM instance, in units of 1/100,000. The maximum value is 1000, for a 1% fee.
	TradingFee uint16
	// A list of vote objects, representing votes on the pool's trading fee.
	VoteSlots []VoteSlots `json:",omitempty"`
}

// ---------------------------------------------
// Asset Object
// ---------------------------------------------

// The definition for one of the two assets the AMM holds.
// In JSON, this is an object with currency and issuer fields.
type Asset struct {
	Currency string        `json:"currency"`
	Issuer   types.Address `json:"issuer,omitempty"`
}

func (a *Asset) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})

	if a.Issuer != "" {
		flattened["issuer"] = a.Issuer
	}

	if a.Currency != "" {
		flattened["currency"] = a.Currency
	}

	return flattened
}

// ---------------------------------------------
// Auction Slot Object
// ---------------------------------------------

// A liquidity provider can bid LP Tokens to claim the auction slot to receive a discount on the trading fee for a 24-hour period.
// The LP tokens that were bid are returned to the AMM.
type AuctionSlot struct {
	// The current owner of this auction slot.
	Account types.Address
	// A list of at most 4 additional accounts that are authorized to trade at the discounted fee for this AMM instance.
	AuthAccounts []AuthAccounts `json:",omitempty"`
	// The trading fee to be charged to the auction owner, in the same format as TradingFee. Normally, this is 1/10 of the normal fee for this AMM.
	DiscountedFee uint32
	// The amount the auction owner paid to win this slot, in LP Tokens.
	Price types.CurrencyAmount
	// The time when this slot expires, in seconds since the Ripple Epoch. https://xrpl.org/docs/references/protocol/data-types/basic-data-types#specifying-time.
	Expiration uint32
}

// ---------------------------------------------
// AuthAccounts Object
// ---------------------------------------------

// A list of up to 4 additional accounts that you allow to trade at the discounted fee.
// This cannot include the address of the transaction sender.
type AuthAccounts struct {
	AuthAccount AuthAccount
}

func (a *AuthAccounts) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})
	flattened["AuthAccount"] = a.AuthAccount.Flatten()
	return flattened
}

// An additional account that you allow to trade at the discounted fee.
type AuthAccount struct {
	// Authorized account to trade at the discounted fee for this AMM instance.
	Account types.Address
}

func (a *AuthAccount) Flatten() map[string]interface{} {
	flattened := make(map[string]interface{})
	flattened["Account"] = a.Account
	return flattened
}

// ---------------------------------------------
// VoteSlots / Vote Entry Objects
// ---------------------------------------------

// Each entry in the vote_slots array represents one liquidity provider's vote to set the trading fee.
type VoteSlots struct {
	VoteEntry VoteEntry
}

// Represents one liquidity provider's vote to set the trading fee.
type VoteEntry struct {
	Account    types.Address
	TradingFee uint
	VoteWeight uint
}

func (*AMM) EntryType() LedgerEntryType {
	return AMMEntry
}
