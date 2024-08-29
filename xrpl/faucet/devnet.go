package faucet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DEVNET_FAUCET_HOST = "faucet.altnet.rippletest.net"
	DEVNET_FAUCET_PATH = "/accounts"
)

var _ FaucetProvider = (*DevnetFaucetProvider)(nil)

type DevnetFaucetProvider struct {
	host        string
	accountPath string
}

func NewDevnetFaucetProvider() *DevnetFaucetProvider {
	return &DevnetFaucetProvider{
		host:        DEVNET_FAUCET_HOST,
		accountPath: DEVNET_FAUCET_PATH,
	}
}

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
