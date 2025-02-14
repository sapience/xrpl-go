---
sidebar_position: 4
---

xrpl is the main package of the `xrpl-go` library. It provides every necessary component to interact with a XRP Ledger.


## Introduction

## Subpackages

### currency

Currency is a package that provides utility functions to handle with XRPL ledger currency types. For **native currency**, it provides XRP and drops conversions. For **IOUs**, it provides utility functions to convert non-standard currency codes (you can learn more about it in the [official documentation](https://xrpl.org/docs/references/protocol/data-types/currency-formats#nonstandard-currency-codes)).

#### API

```go
// XRP <-> Drops conversions
func XrpToDrops(value string) (string, error)
func DropsToXrp(value string) (string, error)

// Non-standard currency codes conversions
func ConvertStringToHex(input string) string
func ConvertHexToString(input string) (string, error)
```

### faucet

Faucet is a package that allows the user to get XRP for testing purposes on `testnet` and `devnet` ledgers and even from custom chains. To be able to fund your accounts programmatically, you can initialize the desired `FaucetProvider` for the ledger you want to use.

The package already exposes the `TestnetFaucetProvider` and `DevnetFaucetProvider` providers. If you want to use a custom chain, you can implement the `FaucetProvider` interface and use your own provider.

#### Usage

For devnet, you can use the following code:

```go
devnetFaucet := faucet.NewDevnetFaucetProvider()

err := devnetFaucet.FundWallet("rJ96831v5JXxna35JYvsW9VRmENwq23ib9")
if err != nil {
    // ...
}
```

for testnet, you can use the following code:

```go
testnetFaucet := faucet.NewTestnetFaucetProvider()

err := testnetFaucet.FundWallet("rJ96831v5JXxna35JYvsW9VRmENwq23ib9")
if err != nil {
    // ...
}
```


### hash

The `hash` package contains functions and types related to the XRPL hash types. Currently, it only contains the function `SignTxBlob` that hashes a signed transaction blob, which is mainly used for multisigning.

#### API

```go
func SignTxBlob(blob []byte, secret string) ([]byte, error)
```

### ledger-entry-types

The `ledger-entry-types` package contains types and functions to handle with XRPL ledger entry types.

#### API

https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types

### queries

### rpc

### time

This package contains functions to handle with XRPL time conversions. It enables conversions between RippleTime and UnixTime. To learn more about RippleTime and UnixTime, you can read the [official documentation](https://xrpl.org/docs/references/protocol/data-types/basic-data-types#specifying-time).

#### API

The following functions are available:

```go
func RippleTimeToUnixTime(rpepoch int64) int64
func UnixTimeToRippleTime(timestamp int64) int64
func RippleTimeToISOTime(rippleTime int64) string
func IsoTimeToRippleTime(isoTime string) (int64, error)
```

### transaction

### wallet

### websocket


## Usage

## Guides





