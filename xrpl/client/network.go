package client

const (
	// Sidechains are expected to have network IDs above this.
	// Networks with ID above this restricted number are expected specify an accurate NetworkID field
	// in every transaction to that chain to prevent replay attacks.
	// Mainnet and testnet are exceptions. More context: https://github.com/XRPLF/rippled/pull/4370
	RESTRICTED_NETWORKID_VERSION uint = 1024
)
