# XRPL-GO

[![Go Reference](https://pkg.go.dev/badge/github.com/Peersyst/xrpl-go/v1.svg)](https://pkg.go.dev/github.com/Peersyst/xrpl-go/v1)
[![Go Report Card](https://goreportcard.com/badge/github.com/Peersyst/xrpl-go/v1)](https://goreportcard.com/report/github.com/Peersyst/xrpl-go/v1)
[![Release Card](https://img.shields.io/github/v/release/Peersyst/xrpl-go?include_prereleases)](https://github.com/Peersyst/xrpl-go/v1/releases)


The `xrpl-go` library provides a Go implementation for interacting with the XRP Ledger. From serialization to signing transactions, the library allows users to work with the most
complex elements of the XRP Ledger. A full library of models for all transactions and core server rippled API objects are provided.

## Disclaimer
This library is still in development and not yet intended for production use.

## Requirements

Requiring Go version `1.22.0` and later.
[Download latest Go version](https://go.dev/dl/)

## Packages

| Name | Description |
|---------|-------------|
| addresscodec | Provides functions for encoding and decoding XRP Ledger addresses |
| binarycodec | Implements binary serialization and deserialization of XRP Ledger objects |
| keypairs | Handles generation and management of cryptographic key pairs for XRP Ledger accounts |
| xrpl | Core package containing the main functionality for interacting with the XRP Ledger |
| examples | Contains example code demonstrating usage of the xrpl-go library |

## Usage

Here's a list of examples of how to use the library.

### Create a wallet
This example shows how to create a new wallet with a random seed or from a given seed.

```go
wallet, err := xrpl.NewWallet(addresscodec.ED25519)
walletFromSeed, _ := xrpl.NewWalletFromSeed("sEdSMVV4dJ1JbdBxmakRR4Puu3XVZz2", "")
walletFromSecret, _ := xrpl.NewWalletFromSecret("sEdSMVV4dJ1JbdBxmakRR4Puu3XVZz2")
```

### Fund a wallet
This example shows how to fund a wallet on the testnet. Devnet wallet funding is also supported. For custom ledger funding, you can implement the `FaucetProvider` interface.
```go
testnetFaucet := faucet.NewTestnetFaucetProvider()
testnetClientCfg := websocket.NewWebsocketClientConfig().
    WithHost("wss://s.altnet.rippletest.net:51233").
    WithFaucetProvider(testnetFaucet)
testnetClient := websocket.NewWebsocketClient(testnetClientCfg)

wallet, err := xrpl.NewWallet(addresscodec.ED25519)
if err != nil {
    // ...
}

err = testnetClient.FundWallet(&wallet)
if err != nil {
    // ...
}
``` 

### Sign and submit a transaction

This example shows how to sign and submit a transaction to the XRP Ledger.
```go
wallet, err := xrpl.NewWalletFromSeed("sEdSMVV4dJ1JbdBxmakRR4Puu3XVZz2", "")
if err != nil {
    // ...
}
receiverWallet, err := xrpl.NewWalletFromSeed("sEd7d8Ci9nevdLCeUMctF3uGXp9WQqJ", "")
if err != nil {
    // ...
}

client := websocket.NewWebsocketClient(
    websocket.NewWebsocketClientConfig().
        WithHost("wss://s.altnet.rippletest.net:51233").
        WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
)

payment := transactions.Payment{
    BaseTx: transactions.BaseTx{
        Account: types.Address(wallet.GetAddress()),
    },
    Destination: types.Address(receiverWallet.GetAddress()),
    Amount:      types.XRPCurrencyAmount(1000000),
}
flatTx := payment.Flatten()

err = client.Autofill(&flatTx)
if err != nil {
    // ...
}

txBlob, _, err := wallet.Sign(flatTx)
if err != nil {
    // ...
}

response, err := client.SubmitTransactionBlob(txBlob, true)
if err != nil {
    // ...
}
```

## Report an issue

If you find any issues, please report them to the [XRPL-GO GitHub repository](https://github.com/Peersyst/xrpl-go/v1/issues). 

## License
The `xrpl-go` library is licensed under the MIT License.
