package faucet

const (
	// Replace with
	USER_AGENT = "xrpl.go"
)

type FaucetProvider interface {
	FundWallet(address string) error
}
