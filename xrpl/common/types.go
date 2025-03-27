package common

import (
	transactions "github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

type SubmitOptions struct {
	Autofill bool
	Wallet   *wallet.Wallet
	FailHard bool
}

type SubmittableTransaction interface {
	TxType() transactions.TxType
	Flatten() transactions.FlatTransaction
}
