package transaction

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

const (
	tfClawTwoAssets uint32 = 0x00000001
)

var (
	ErrInvalidHolder       = errors.New("invalid holder")
	ErrInvalidAmountIssuer = errors.New("invalid amount issuer")
)

type AMMClawback struct {
	BaseTx
	Holder string
	Asset  types.IssuedCurrency
	Asset2 types.CurrencyAmount       `json:",omitempty"`
	Amount types.IssuedCurrencyAmount `json:",omitempty"`
}

func (a *AMMClawback) TxType() TxType {
	return AMMClawbackTx
}

func (a *AMMClawback) Flatten() FlatTransaction {
	flattened := a.BaseTx.Flatten()

	if a.Holder != "" {
		flattened["Holder"] = a.Holder
	}

	if a.Asset != (types.IssuedCurrency{}) {
		flattened["Asset"] = a.Asset.Flatten()
	}

	if a.Asset2 != nil {
		flattened["Asset2"] = a.Asset2.Flatten()
	}

	if a.Amount != (types.IssuedCurrencyAmount{}) {
		flattened["Amount"] = a.Amount.Flatten()
	}

	return flattened
}

func (a *AMMClawback) Validate() (bool, error) {
	_, err := a.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if !addresscodec.IsValidAddress(a.Holder) {
		return false, ErrInvalidHolder
	}

	if a.Asset != (types.IssuedCurrency{}) {
		if !addresscodec.IsValidAddress(a.Asset.Issuer) {
			return false, ErrInvalidAssetIssuer
		}
	}

	if a.Amount != (types.IssuedCurrencyAmount{}) {
		if !addresscodec.IsValidAddress(a.Amount.Issuer.String()) {
			return false, ErrInvalidAmountIssuer
		}
	}

	return true, nil
}

func (a *AMMClawback) SetClawTwoAssets() {
	a.Flags |= tfClawTwoAssets
}
