package faucet

const (
	USER_AGENT = "xrpl.go"
)

// FaucetProvider defines an interface for interacting with XRPL faucets.
// Implementations of this interface can be used to fund wallets on different
// XRPL networks (e.g., Devnet, Testnet) by requesting XRP from their respective faucets.
type FaucetProvider interface {
	// FundWallet sends a request to the faucet to fund the specified wallet address.
	// It returns an error if the funding request fails.
	FundWallet(address string) error
}
