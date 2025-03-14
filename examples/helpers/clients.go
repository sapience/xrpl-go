package helpers

import (
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/rpc"
)

// GetDevnetRpcClient returns a new rpc client for the devnet
func GetDevnetRpcClient() *rpc.Client {
	cfg, err := rpc.NewClientConfig(
		"https://s.devnet.rippletest.net:51234",
		rpc.WithFaucetProvider(faucet.NewDevnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	return rpc.NewClient(cfg)
}

// GetTestnetRpcClient returns a new rpc client for the testnet
func GetTestnetRpcClient() *rpc.Client {
	cfg, err := rpc.NewClientConfig(
		"https://s.altnet.rippletest.net:51234",
		rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	return rpc.NewClient(cfg)
}
