package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// Withdraw assets from an Automated Market Maker (AMM) instance by returning the AMM's liquidity provider tokens (LP Tokens).
//
// # Example
//
// ```json
//
//	{
//	    "Account" : "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
//	    "Amount" : {
//	        "currency" : "TST",
//	        "issuer" : "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd",
//	        "value" : "5"
//	    },
//	    "Amount2" : "50000000",
//	    "Asset" : {
//	        "currency" : "TST",
//	        "issuer" : "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd"
//	    },
//	    "Asset2" : {
//	        "currency" : "XRP"
//	    },
//	    "Fee" : "10",
//	    "Flags" : 1048576,
//	    "Sequence" : 10,
//	    "TransactionType" : "AMMWithdraw"
//	}
//
// / ```
type AMMWithdraw struct {
	BaseTx
	// The definition for one of the assets in the AMM's pool. In JSON, this is an object with currency and issuer fields (omit issuer for XRP).
	Asset ledger.Asset
	// The definition for the other asset in the AMM's pool. In JSON, this is an object with currency and issuer fields (omit issuer for XRP).
	Asset2 ledger.Asset
	// The amount of one asset to withdraw from the AMM. This must match the type of one of the assets (tokens or XRP) in the AMM's pool.
	Amount types.CurrencyAmount `json:",omitempty"`
	// The amount of another asset to withdraw from the AMM. If present, this must match the type of the other asset in the AMM's pool and cannot be the same type as Amount.
	Amount2 types.CurrencyAmount `json:",omitempty"`
	// The minimum effective price, in LP Token returned, to pay per unit of the asset to withdraw.
	EPrice types.CurrencyAmount `json:",omitempty"`
	// How many of the AMM's LP Tokens to redeem.
	LPTokenIn types.CurrencyAmount `json:",omitempty"`
}

// ****************************
// AMMWithdraw Flags
// ****************************

const (
	// Perform a double-asset withdrawal returning all your LP Tokens.
	tfWithdrawAll uint = 131072
	// Perform a single-asset withdrawal returning all of your LP Tokens.
	tfOneAssetWithdrawAll uint = 262144
)

// Perform a double-asset withdrawal and receive the specified amount of LP Tokens.
func (a *AMMWithdraw) SetLPTokentFlag() {
	a.Flags |= tfLPToken
}

// Perform a double-asset withdrawal returning all your LP Tokens.
func (a *AMMWithdraw) SetWithdrawAllFlag() {
	a.Flags |= tfWithdrawAll
}

// Perform a single-asset withdrawal returning all of your LP Tokens.
func (a *AMMWithdraw) SetOneAssetWithdrawAllFlag() {
	a.Flags |= tfOneAssetWithdrawAll
}

// Perform a single-asset withdrawal with a specified amount of the asset to withdrawal.
func (a *AMMWithdraw) SetSingleAssetFlag() {
	a.Flags |= tfSingleAsset
}

// Perform a double-asset withdrawal with specified amounts of both assets.
func (a *AMMWithdraw) SetTwoAssetFlag() {
	a.Flags |= tfTwoAsset
}

// Perform a single-asset withdrawal and receive the specified amount of LP Tokens.
func (a *AMMWithdraw) SetOneAssetLPTokenFlag() {
	a.Flags |= tfOneAssetLPToken
}

// Perform a single-asset withdrawal with a specified effective price.
func (a *AMMWithdraw) SetLimitLPTokenFlag() {
	a.Flags |= tfLimitLPToken
}

func (*AMMWithdraw) TxType() TxType {
	return AMMWithdrawTx
}

// TODO: Implement flatten
func (s *AMMWithdraw) Flatten() FlatTransaction {
	return nil
}
