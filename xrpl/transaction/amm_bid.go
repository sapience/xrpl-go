package transaction

import (
	"encoding/json"
	"errors"

	ledger "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// Bid on an Automated Market Maker's (AMM's) auction slot. If you win, you can trade against the AMM at a discounted fee until you are outbid or 24 hours have passed.
// If you are outbid before 24 hours have passed, you are refunded part of the cost of your bid based on how much time remains.
// If the AMM's trading fee is zero, you can still bid, but the auction slot provides no benefit unless the trading fee changes.
// You bid using the AMM's LP Tokens; the amount of a winning bid is returned to the AMM, decreasing the outstanding balance of LP Tokens.
type AMMBid struct {
	BaseTx
	// The definition for one of the assets in the AMM's pool. In JSON, this is an object with currency and issuer fields (omit issuer for XRP).
	Asset ledger.Asset
	// The definition for the other asset in the AMM's pool. In JSON, this is an object with currency and issuer fields (omit issuer for XRP).
	Asset2 ledger.Asset
	// Pay at least this amount for the slot. Setting this value higher makes it harder for others to outbid you. If omitted, pay the minimum necessary to win the bid.
	BidMin types.CurrencyAmount `json:",omitempty"`
	// Pay at most this amount for the slot. If the cost to win the bid is higher than this amount, the transaction fails. If omitted, pay as much as necessary to win the bid.
	BidMax types.CurrencyAmount `json:",omitempty"`
	// A list of up to 4 additional accounts that you allow to trade at the discounted fee. This cannot include the address of the transaction sender. Each of these objects should be an Auth Account object.
	AuthAccounts []ledger.AuthAccounts `json:",omitempty"`
}

func (*AMMBid) TxType() TxType {
	return AMMBidTx
}

// UnmarshalJSON unmarshals the Payment transaction from JSON.
func (p *AMMBid) UnmarshalJSON(data []byte) error {
	type pHelper struct {
		BaseTx
		Asset        json.RawMessage
		Asset2       json.RawMessage
		BidMin       json.RawMessage       `json:",omitempty"`
		BidMax       json.RawMessage       `json:",omitempty"`
		AuthAccounts []ledger.AuthAccounts `json:",omitempty"`
	}
	var h pHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*p = AMMBid{
		BaseTx: h.BaseTx,
	}
	var asset, asset2 ledger.Asset
	var bidMin, bidMax types.CurrencyAmount
	var authAccounts []ledger.AuthAccounts
	var err error

	asset, err = ledger.UnmarshalAsset(h.Asset)
	if err != nil {
		return err
	}
	asset2, err = ledger.UnmarshalAsset(h.Asset2)
	if err != nil {
		return err
	}

	p.Asset = asset
	p.Asset2 = asset2

	bidMin, err = types.UnmarshalCurrencyAmount(h.BidMin)
	if err != nil {
		return err
	}

	bidMax, err = types.UnmarshalCurrencyAmount(h.BidMax)
	if err != nil {
		return err
	}

	p.BidMin = bidMin
	p.BidMax = bidMax

	authAccounts = append(authAccounts, h.AuthAccounts...)

	p.AuthAccounts = authAccounts

	return nil
}

func (a *AMMBid) Flatten() FlatTransaction {
	// Add BaseTx fields
	flattened := a.BaseTx.Flatten()

	// Add AMMBid-specific fields
	flattened["TransactionType"] = "AMMBid"

	flattened["Asset"] = a.Asset.Flatten()

	flattened["Asset2"] = a.Asset2.Flatten()

	if a.BidMin != nil {
		flattened["BidMin"] = a.BidMin.Flatten()
	}

	if a.BidMax != nil {
		flattened["BidMax"] = a.BidMax.Flatten()
	}

	if len(a.AuthAccounts) > 0 {
		authAccountsFlattened := make([]map[string]interface{}, 0)

		for _, authAccount := range a.AuthAccounts {
			authAccountsFlattened = append(authAccountsFlattened, authAccount.Flatten())
		}

		flattened["AuthAccounts"] = authAccountsFlattened
	}

	return flattened
}

func (a *AMMBid) Validate() (bool, error) {
	_, err := a.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if ok, err := IsAsset(a.Asset); !ok {
		return false, err
	}

	if ok, err := IsAsset(a.Asset2); !ok {
		return false, err
	}

	if a.Asset.Currency == "XRP" && a.Asset2.Currency == "XRP" {
		return false, errors.New("at least one of the assets must be non-XRP")
	}

	if ok, err := IsAmount(IsAmountArgs{field: a.BidMin, fieldName: "BidMin", isFieldRequired: false}); !ok {
		return false, err
	}

	if ok, err := IsAmount(IsAmountArgs{field: a.BidMax, fieldName: "BidMax", isFieldRequired: false}); !ok {
		return false, err
	}

	if ok, err := validateAuthAccounts(a.AuthAccounts); !ok {
		return false, err
	}

	return true, nil
}

// Validate the AuthAccounts field.
func validateAuthAccounts(authAccounts []ledger.AuthAccounts) (bool, error) {
	if len(authAccounts) > 4 {
		return false, errors.New("authAccounts: AuthAccounts should have at most 4 fields")
	}

	// TODO: check that the AuthAccount 'Account' field is a valid XRP address when this function is available

	return true, nil
}
