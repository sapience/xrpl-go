package faucet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	TESTNET_FAUCET_HOST = "faucet.altnet.rippletest.net"
	TESTNET_FAUCET_PATH = "/accounts"
)

var _ FaucetProvider = (*TestnetFaucetProvider)(nil)

type TestnetFaucetProvider struct {
	host        string
	accountPath string
}

func NewTestnetFaucetProvider() *TestnetFaucetProvider {
	return &TestnetFaucetProvider{
		host:        TESTNET_FAUCET_HOST,
		accountPath: TESTNET_FAUCET_PATH,
	}
}

func (fp *TestnetFaucetProvider) FundWallet(address string) error {
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
