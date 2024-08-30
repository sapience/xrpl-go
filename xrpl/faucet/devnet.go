package faucet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DEVNET_FAUCET_HOST = "faucet.devnet.rippletest.net"
	DEVNET_FAUCET_PATH = "/accounts"
)

var _ FaucetProvider = (*DevnetFaucetProvider)(nil)

// DevnetFaucetProvider implements the FaucetProvider interface for the XRPL Devnet.
// It provides functionality to interact with the Devnet faucet for funding wallets.
type DevnetFaucetProvider struct {
	host        string // The hostname of the Devnet faucet
	accountPath string // The API path for account-related operations
}

// NewDevnetFaucetProvider creates and returns a new instance of DevnetFaucetProvider
// with predefined Devnet faucet host and account path.
func NewDevnetFaucetProvider() *DevnetFaucetProvider {
	return &DevnetFaucetProvider{
		host:        DEVNET_FAUCET_HOST,
		accountPath: DEVNET_FAUCET_PATH,
	}
}

// FundWallet sends a request to the Devnet faucet to fund the specified wallet address.
// It returns an error if the funding request fails.
func (fp *DevnetFaucetProvider) FundWallet(address string) error {
	url := fmt.Sprintf("https://%s%s", fp.host, fp.accountPath)
	payload := map[string]string{"destination": address, "userAgent": USER_AGENT}
	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
