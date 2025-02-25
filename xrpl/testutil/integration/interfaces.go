package integration

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type FaucetProvider interface {
	FundWallet(address types.Address) error
}
