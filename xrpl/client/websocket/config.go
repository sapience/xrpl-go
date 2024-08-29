package websocket

import "github.com/Peersyst/xrpl-go/xrpl/faucet"

type WebsocketClientConfig struct {
	// Connection config
	host 		string

	// Fee config
	feeCushion float32
	maxFeeXRP  float32

	// Faucet config
	faucetProvider faucet.FaucetProvider
}

func NewWebsocketClientConfig() *WebsocketClientConfig {
	return &WebsocketClientConfig{
		host: "",
		feeCushion: DEFAULT_FEE_CUSHION,
		maxFeeXRP: DEFAULT_MAX_FEE_XRP,
	}
}

func (wc WebsocketClientConfig) WithHost(host string) WebsocketClientConfig {
	wc.host = host
	return wc
}

func (wc WebsocketClientConfig) WithFeeCushion(feeCushion float32) WebsocketClientConfig {
	wc.feeCushion = feeCushion
	return wc
}

func (wc WebsocketClientConfig) WithMaxFeeXRP(maxFeeXrp float32) WebsocketClientConfig {
	wc.maxFeeXRP = maxFeeXrp
	return wc
}

func (wc WebsocketClientConfig) WithFaucetProvider(fp faucet.FaucetProvider) WebsocketClientConfig {
	wc.faucetProvider = fp
	return wc
}