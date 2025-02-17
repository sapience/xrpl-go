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

#### Usage

To import the package, you can use the following code:

```go
import "github.com/Peersyst/xrpl-go/xrpl/hash"
```

#### API

```go
func SignTxBlob(blob []byte, secret string) ([]byte, error)
```

### ledger-entry-types

The `ledger-entry-types` package contains types and functions to handle ledger objects. They are used by other packages like `transactions` to type the transaction's fields.

- [`AccountRoot`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/accountroot)
- [`Amendments`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/amendments)
- [`AMM`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/amm)
- [`Bridge`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/bridge)
- [`Check`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/check)
- [`DepositPreauth`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/depositpreauth)
- [`Did`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/did)
- [`DirectoryNode`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/directorynode)
- [`Escrow`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/escrow)
- [`FeeSettings`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/feesettings)
- [`Ledger`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/ledger)
- [`Hashes`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/ledgerhashes)
- [`NegativeUNL`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/negativeunl)
- [`NFTokenOffer`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/nftokenoffer)
- [`NFTokenPage`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/nftokenpage)
- [`Offer`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/offer)
- [`Oracle`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/oracle)
- [`PayChannel`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/paychannel)
- [`RippleState`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/ripplestate)
- [`SignerList`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/signerlist)
- [`Ticket`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/ticket)
- [`XChainOwnedClaimID`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/xchainownedclaimid)
- [`XChainOwnedCreateAccountClaimID`](https://xrpl.org/docs/references/protocol/ledger-data/ledger-entry-types/xchainownedcreateaccountclaimid)

#### Usage

To import the package, you can use the following code:

```go
import "github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
```

### queries

The `queries` package contains mainly requests and responses types for the [XRPL methods](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods). This package is used by the package clients `rpc` and `websocket` to send client queries to the ledger.

:::info

As a developer, you may be interested in calling the queries using the `websocket` or `rpc` clients. Both clients expose methods to call each query exposed by the `queries` package.    

:::

Queries are grouped by different categories or packages:

- `account`: Methods to work with account info.
- `channel`: Methods to work with channels.
- `ledger`: Methods to retrieve ledger info.
- `transaction`: Submit and query ledger transactions.
- `path`: Methods to use paths and order books.
- `nft`: Methods to work with NFTs.
- `oracle`: Methods to work with oracles.
- `clio`: Methods to use the Clio API, not `rippled`.
- `server`: Methods to retrieve information about the current state of the rippled server.
- `utility`: Perform convenient tasks, such as ping and random number generation.


#### API version

By default, all queries are meant to be used with the latest XRPL API version (currently `v2`). If you want to use a specific version, you will need to import the specific version queries package from each subpackage.

For example, if you want to use the XRPL API version `v1` queries of the `account` subpackage, you will need to import it this way:
```go
import accountv1 "github.com/Peersyst/xrpl-go/xrpl/queries/account/v1"
```

#### account

The `account` package contains methods to interact with XRPL accounts. These methods allow you to:

- Retrieve account information like balances, settings, and objects
- Get account transaction history
- Query account channels and escrows
- Check account offers and payment channels

The available methods correspond to the [Account Methods](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods#account-methods) in the XRPL API.

The account subpackage provides the following queries requests:

| Request | Method name | V1 support |
|---------|------------|------------|
| `ChannelRequest` | [account_channels](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/account-methods/account_channels) | ✅ |
| `CurrenciesRequest` | [account_currencies](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/account-methods/account_currencies) | ✅ |
| `GatewayBalancesRequest` | [gateway_balances](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/account-methods/gateway_balances) | ✅ |
| `InfoRequest` | [account_info](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/account-methods/account_info) | ✅ |
| `LinesRequest` | [account_lines](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/account-methods/account_lines) | ✅ |
| `NFTsRequest` | [account_nfts](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/account-methods/account_nfts) | ✅ |
| `NoRippleCheckRequest` | [noripple_check](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/account-methods/noripple_check) | ✅ |
| `ObjectsRequest` | [account_objects](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/account-methods/account_objects) | ✅ |
| `OffersRequest` | [account_offers](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/account-methods/account_offers) | ✅ |
| `TransactionsRequest` | [account_tx](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/account-methods/account_tx) | ✅ |

#### channel

The `channel` package contains methods to interact with XRPL channels. These methods allow you to:

- Verify the channel's state.

The available methods correspond to the [Payment Channel Methods](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/payment-channel-methods) in the XRPL API.

The `channel` subpackage provides the following queries requests:

| Request | Method name | V1 support |
|---------|------------|------------|
| `VerifyRequest` | [channel_verify](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/payment-channel-methods/channel_verify) | ✅ |

#### ledger

The `ledger` package contains methods to interact with XRPL ledgers. These methods allow you to:

- Retrieve specific, current or closed ledger information.

The available methods correspond to the [Ledger Methods](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/ledger-methods) in the XRPL API.

The `ledger` subpackage provides the following queries requests:

| Request | Method name | V1 support |
|---------|------------|------------|
| `Request` | [ledger](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/ledger-methods/ledger) | ✅ |
| `ClosedRequest` | [ledger_closed](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/ledger-methods/ledger_closed) | ✅ |
| `CurrentRequest` | [ledger_current](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/ledger-methods/ledger_current) | ✅ |
| `DataRequest` | [ledger_data](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/ledger-methods/ledger_data) | ✅ |

#### transaction

The `transaction` package contains methods to interact with XRPL transactions. These methods allow you to:

- Submit ledger transactions.
- Query ledger transactions.

The available methods correspond to the [Transaction Methods](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/transaction-methods) in the XRPL API.

The `transaction` subpackage provides the following queries requests:

| Request | Method name | V1 support |
|---------|------------|------------|
| `SubmitRequest` | [submit](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/transaction-methods/submit) | ✅ |
| `SubmitMultisignedRequest` | [submit_multisigned](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/transaction-methods/submit_multisigned) | ✅ |
| `EntryRequest` | [transaction_entry](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/transaction-methods/transaction_entry) | ✅ |
| `TxRequest` | [tx](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/transaction-methods/tx) | ✅ |


#### path, nft and oracle

The `path`, `nft` and `oracle` packages contain methods to interact with XRPL paths, NFTs and oracles. These methods allow you to:

- Retrieve paths and order books.
- Retrieve AMMs information.
- Get NFTs buy and sell offers.

The available methods correspond to the [Path and Order Book Methods](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/path-and-order-book-methods) in the XRPL API.

The `path` subpackage provides the following queries requests:

| Request | Method name | V1 support |
|---------|------------|------------|
| `BookOffersRequest` | [book_offers](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/path-and-order-book-methods/book_offers) | ✅ |
| `DepositAuthorizedRequest` | [deposit_authorized](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/path-and-order-book-methods/deposit_authorized) | ✅ |
| `FindCreateRequest`, `FindCloseRequest`, `FindStatusRequest` | [path_find](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/path-and-order-book-methods/path_find) | ✅ |
| `RipplePathFindRequest` | [ripple_path_find](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/path-and-order-book-methods/ripple_path_find) | ✅ |


The `nft` subpackage provides the following queries requests:

| Request | Method name | V1 support |
|---------|------------|------------|
| `NFTokenBuyOffersRequest` | [nft_buy_offers](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/path-and-order-book-methods/nft_buy_offers) | ✅ |
| `NFTokenSellOffersRequest` | [nft_sell_offers](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/path-and-order-book-methods/nft_sell_offers) | ✅ |

The `oracle` subpackage provides the following queries requests:

| Request | Method name | V1 support |
|---------|------------|------------|
| `GetAggregatePriceRequest` | [get_aggregate_price](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/oracle-methods/get_aggregate_price) | ✅ |

#### clio

The `clio` package contains methods to interact with the Clio API, not `rippled`. These methods allow you to:

- Retrieve NFT history.
- Retrieve NFts information.

The available methods correspond to the [Clio Methods](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/clio-methods) in the XRPL API.

The `clio` subpackage provides the following queries requests:

| Request | Method name | V1 support |
|---------|------------|------------|
| `NFTHistoryRequest` | [nft_history](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/clio-methods/nft_history) | ✅ |
| `NFTInfoRequest` | [nft_info](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/clio-methods/nft_info) | ✅ |
| `NFTsByIssuerRequest` | [nfts_by_issuer](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/clio-methods/nfts_by_issuer) | ✅ |

#### server

The `server` package contains methods to interact with the `rippled` server. These methods allow you to:

- Retrieve server information.
- Get fee information.
- Get the manifest.

The available methods correspond to the [Server Info Methods](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/server-info-methods) in the XRPL API.

The `server` subpackage provides the following queries requests:

| Request | Method name | V1 support |
|---------|------------|------------|
| `FeatureAllRequest` | [feature_all](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/server-info-methods/feature_all) | ✅ |
| `FeatureOneRequest` | [feature](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/server-info-methods/feature) | ✅ |
| `FeeRequest` | [fee](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/server-info-methods/fee) | ✅ |
| `ManifestRequest` | [manifest](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/server-info-methods/manifest) | ✅ |
| `InfoRequest` | [server_info](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/server-info-methods/server_info) | ✅ |
| `StateRequest` | [server_state](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/server-info-methods/server_state) | ✅ |


#### utility

The `utility` package contains methods to interact with the XRPL utility. These methods allow you to:

- Retrieve a random number.
- Ping the server.

The available methods correspond to the [Utility Methods](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/utility-methods) in the XRPL API.

The `utility` subpackage provides the following queries requests:

| Request | Method name | V1 support |
|---------|------------|------------|
| `RandomRequest` | [random](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/utility-methods/random) | ✅ |
| `PingRequest` | [ping](https://xrpl.org/docs/references/http-websocket-apis/public-api-methods/utility-methods/ping) | ✅ |

### rpc

The `rpc` package provides the RPC client for interacting with the XRPL network via its RPC API. This client handles the communication with XRPL nodes, allowing you to:

- Send requests to query the ledger state
- Submit transactions to the network 
- Receive responses and handle errors
- Manage the connections configuration 

#### Client

The `rpc` package provides a `Client` type which and communication with XRPL nodes. This client is configurable and let the user submit transactions and make queries. 

In order to create a new `Client`, you can use the `NewClient` function:

```go
cfg, err := rpc.NewClientConfig("<url>")
if err != nil {
	// ...
}
client := rpc.NewClient(cfg)
```

Every time you create a new `Client`, you need to pass a `Config` struct as argument. You can initialize a `Config` struct using the `NewClientConfig` function.

`Config` struct follows the options pattern, so you can pass different options to the `NewClientConfig` function:

```go
// RPC Client config options
func WithHTTPClient(cl HTTPClient) ConfigOpt
func WithMaxFeeXRP(maxFeeXRP float32) ConfigOpt
func WithFeeCushion(feeCushion float32) ConfigOpt
func WithFaucetProvider(fp common.FaucetProvider) ConfigOpt
```

So, for example, if you want to set a custom `FaucetProvider` and `FeeCushion`, you can do it this way:

```go
cfg, err := rpc.NewClientConfig("https://s.altnet.rippletest.net:51234/",
	rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	rpc.WithFeeCushion(1.5),
)
if err != nil {
	// ...
}
client := rpc.NewClient(cfg)
```

#### Methods

`Client` offers different methods to interact with the XRPL network.

##### Request

The `Request` method is used to make queries to the XRPL network. It returns a `XRPLResponse` interface. This method is used in the client's queries requests.

```go
// Client methods
func (c *Client) Request(reqParams XRPLRequest) (XRPLResponse, error)
```

##### Submit/SubmitMultisigned

The `Submit` method is used to submit a transaction to the XRPL network. It returns a `TxResponse` struct containing the transaction result for the blob submitted. `txBlob` must be signed. There's also a `SubmitMultisigned` method that works the same way but for multisigned transactions.

```go
func (c *Client) Submit(txBlob string, failHard bool) (*requests.TxResponse, error)
func (c *Client) SubmitMultisigned(txBlob string, failHard bool) (*requests.SubmitMultisignedResponse, error)
```

##### Autofill/AutofillMultisigned

The `Autofill` method is used to autofill some fields in a flat transaction. This method is useful for adding dynamic fields like `LastLedgerSequence` or `Fee`. It returns an error if the transaction is not valid or some internall call fails. There's also a `AutofillMultisigned` method that works the same way but for multisigned transactions.

```go
func (c *Client) Autofill(tx *transaction.FlatTransaction) error
func (c *Client) AutofillMultisigned(tx *transaction.FlatTransaction, nSigners uint64) error
```

##### SubmitAndWait

The `SubmitAndWait` method is used to submit a transaction to the XRPL network and wait for it to be included in a ledger. It returns a `TxResponse` struct containing the transaction result for the blob submitted. 

```go
func (c *Client) SubmitAndWait(txBlob string, failHard bool) (*requests.TxResponse, error)
```

#### Queries

`Client` also exposes methods to make queries to the XRPL network. These methods are wrappers of the queries requests exposed by the [`queries`](/xrpl/queries/) package.

#### Usage

To use the `rpc` package, you need to import it in your project:

```go
import "github.com/Peersyst/xrpl-go/xrpl/rpc"
```

#### Examples

##### How to send a payment transaction

This example shows how to send a payment transaction to the XRPL testnet with the `rpc` package.

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/Peersyst/xrpl-go/xrpl/currency"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

func main() {

	// Create a new rpc client config with a testnet faucet provider
	cfg, err := rpc.NewClientConfig(
		"https://s.altnet.rippletest.net:51234/",
		rpc.WithMaxFeeXRP(5.0),
		rpc.WithFeeCushion(1.5),
		rpc.WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	// Create a new rpc client with the config
	client := rpc.NewClient(cfg)

	// Create a new wallet with the ed25519 algorithm
	w, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fund the wallet with the testnet faucet
	if err := client.FundWallet(&w); err != nil {
		fmt.Println(err)
		return
	}

	// Convert the amount to drops
	xrpAmount, err := currency.XrpToDrops("1")
	if err != nil {
		fmt.Println(err)
		return
	}

	xrpAmountInt, err := strconv.ParseInt(xrpAmount, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account: types.Address(w.GetAddress()),
		},
		Destination: "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		Amount:      types.XRPCurrencyAmount(xrpAmountInt),
		DeliverMax:  types.XRPCurrencyAmount(xrpAmountInt),
	}

	flattenedTx := p.Flatten()

	// Autofill the transaction with the client's config
	if err := client.Autofill(&flattenedTx); err != nil {
		fmt.Println(err)
		return
	}

	// Sign the transaction with the wallet
	txBlob, _, err := w.Sign(flattenedTx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Submit the transaction to the network and wait for it to be included in a ledger
	res, err := client.SubmitAndWait(txBlob, false)
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

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

[Introduction]

[Explain Flatten()] 

[List transaction types]

#### API

[List API]

### wallet

The `wallet` package contains the types and functions to work and manage your XRPL accounts. Either you want to create a new account, or you want to sign transactions, this package has you covered.

This package enables you to do the following actions:

- Generate new wallets using a seed, mnemonic or random.
- Sign and multisign transactions.
- Access to wallet's public and private keys and address.

#### Generating a wallet

In order to generate a new wallet, you can either use a **seed**, a **mnemonic** or generate a **random** one. Here are the constructors available:

```go
// Wallet constructors
func New(alg interfaces.CryptoImplementation) (Wallet, error)
func FromSeed(seed string, masterAddress string) (Wallet, error)
func FromSecret(seed string) (Wallet, error)
func FromMnemonic(mnemonic string) (*Wallet, error)
```

:::info

When generating a random wallet, you will need to specify the algorithm you want to use.
`xrpl-go` library provides the package `crypto` that exports `ed25519` and `secp256k1` algorithms which satisfy the `CryptoImplementation` interface.
You can use the `crypto` package by importing it in your project:

```go
import "github.com/xrpl-go/pkg/crypto"
```

:::

:::warning

When initializing a wallet from a seed, remember that only seeds generated by `ed25519` and `secp256k1` algorithms are supported. Learn more about XRPL cryptographic keys in the [official documentation](https://xrpl.org/docs/concepts/accounts/cryptographic-keys).

:::

#### Signing and multisigning transactions

A wallet lets the developer sign and multisign transactions easily. The `Wallet` type exposes the following signing methods:

```go
// Signing methods
func (w *Wallet) Sign(tx map[string]interface{}) (string, string, error)
func (w *Wallet) Multisign(tx map[string]interface{}) (string, string, error)
```

The `Sign` method signs a [flat transaction](Afegir link) and returns the signed transaction blob and the signature.

On the other hand, the `Multisign` method multisigns a flat transaction by adding the wallet's signature to the transaction and returning the resulting transaction blob and the blob hash. Learn more about how multisigns work in the [official documentation](https://xrpl.org/docs/concepts/accounts/multi-signing).

#### Usage

In this section, we will see how to generate a `Wallet`, call the faucet to get XRP, and send the XRP to another account.
First step is to generate a `Wallet` using the `New` constructor (in this case, we will use the `ed25519` algorithm):

```go
wallet, err := wallet.New(crypto.ED25519())
if err != nil {
    // ...
}
```
Once we have the `Wallet`, we can call the faucet to get XRP. For this example, we will use the `DevnetFaucetProvider` to get XRP on the `devnet` ledger:

```go
devnetFaucet := faucet.NewDevnetFaucetProvider()

err := devnetFaucet.FundWallet(wallet.ClassicAddress)
if err != nil {
    // ...
}
```

Once we have the XRP, we can create a `Payment` transaction. For this example, we will send the XRP to the `rJ96831v5JXxna35JYvsW9VRmENwq23ib9` account.

```go
payment := transaction.Payment{
    BaseTx: transaction.BaseTx{
        Account: wallet.ClassicAddress,
    },
    Destination: "rJ96831v5JXxna35JYvsW9VRmENwq23ib9",
    Amount:      types.XRPCurrencyAmount(10000000),
    DeliverMax:  types.XRPCurrencyAmount(10000000),
}
```

Finally, we can sign the flat payment transaction:

```go
blob, hash, err := wallet.Sign(payment.Flatten())
if err != nil {
    // ...
}
```

Summarizing, the complete code to generate a wallet, call the faucet to get XRP, create a payment transaction and sign it is the following:

```go
package main

import (
	"fmt"
	"log"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

func main() {
	wallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		log.Fatal(err)
	}

	payment := transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account: wallet.ClassicAddress,
		},
		Destination: "rJ96831v5JXxna35JYvsW9VRmENwq23ib9",
		Amount:      types.XRPCurrencyAmount(10000000),
		DeliverMax:  types.XRPCurrencyAmount(10000000),
	}

	blob, hash, err := wallet.Sign(payment.Flatten())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tx blob: ", blob)
	fmt.Println("Tx hash: ", hash)
}
```

### websocket

The `websocket` package provides a WebSocket client for interacting with the XRPL network via its WebSocket API. This client handles the communication with XRPL nodes, allowing you to:

- Send requests to query the ledger state
- Submit transactions to the network 
- Receive responses and handle errors
- Manage the connections configuration

#### Config

The `websocket` package provides a `Config` struct that allows you to configure the WebSocket client. Every time you create a new `Client`, you need to pass a `Config` struct as argument. You can initialize a `Config` struct using the `NewClientConfig` function.

`Config` struct follows the options pattern, so you can pass different options to the `NewClientConfig` function:

##### Host

The `WithHost` option allows you to set the host of the WebSocket client.

```go
func (wc ClientConfig) WithHost(host string) ClientConfig
```

##### FaucetProvider

The `FaucetProvider` option allows you to set the faucet provider of the WebSocket client. There're two predefined faucet providers: `TestnetFaucetProvider` and `DevnetFaucetProvider`. You can also implement your own faucet provider by implementing the `FaucetProvider` interface.

```go
func (wc ClientConfig) WithFaucetProvider(fp common.FaucetProvider) ClientConfig
```

##### MaxRetries

The `WithMaxRetries` option allows you to set the maximum number of retries for a transaction.

```go
func (wc ClientConfig) WithMaxRetries(maxRetries int) ClientConfig
```

##### RetryDelay

The `WithRetryDelay` option allows you to set the delay between retries for a transaction.

```go
func (wc ClientConfig) WithRetryDelay(retryDelay time.Duration) ClientConfig
```

##### FeeCushion

The `WithFeeCushion` option allows you to set the fee cushion for a transaction.

```go
func (wc ClientConfig) WithFeeCushion(feeCushion float32) ClientConfig
```

##### MaxFeeXRP

The `WithMaxFeeXRP` option allows you to set the maximum fee in XRP that the WebSocket client will use.

```go
func (wc ClientConfig) WithMaxFeeXRP(maxFeeXrp float32) ClientConfig
```

#### Connection

As the `websocket` package is a WebSocket client, it needs to be connected to a WebSocket server. The `Client` type exposes the following methods to connect to a WebSocket server:

```go
// Connection methods
func (c *Client) Connect() error
func (c *Client) Disconnect() error

// Connection status
func (c *Client) IsConnected() bool

// Connection
func (c *Client) Conn() *websocket.Conn
```


So, for example, if you want to connect to the `devnet` ledger, you can do it this way:

```go
client := websocket.NewClient(websocket.NewClientConfig().WithHost("wss://s.altnet.rippletest.net:51233"))
defer client.Disconnect()

err := client.Connect()
if err != nil {
    // ...
}

if !client.IsConnected() {
    // ...
}
```

#### Methods

The `Client` type exposes the following methods to interact with the XRPL network:

##### Request

The `Request` method is used to send a request to the server and returns the response. This method is mostly used to send client [queries](/docs/docs/packages/xrpl#queries) to the server.

```go
func (c *Client) Request(reqParams XRPLRequest) (*ClientResponse, error)
```



##### Autofill/AutofillMultisigned

The `Autofill` method is used to autofill some fields in a flat transaction. This method is useful for adding dynamic fields like `LastLedgerSequence` or `Fee`. It returns an error if the transaction is not valid or some internall call fails. There's also a `AutofillMultisigned` method that works the same way but for multisigned transactions.

```go
func (c *Client) Autofill(tx *transaction.FlatTransaction) error
func (c *Client) AutofillMultisigned(tx *transaction.FlatTransaction, nSigners uint64) error
```

##### Submit/SubmitMultisigned

The `Submit` method is used to submit a transaction to the XRPL network. It returns a `TxResponse` struct containing the transaction result for the blob submitted. `txBlob` must be signed. There's also a `SubmitMultisigned` method that works the same way but for multisigned transactions.

```go
func (c *Client) Submit(txBlob string, failHard bool) (*requests.SubmitResponse, error)
func (c *Client) SubmitMultisigned(txBlob string, failHard bool) (*requests.SubmitMultisignedResponse, error)
```

##### SubmitAndWait

The `SubmitAndWait` method is used to submit a transaction to the XRPL network and wait for it to be included in a ledger. It returns a `TxResponse` struct containing the transaction result for the blob submitted. 

```go
func (c *Client) SubmitAndWait(txBlob string, failHard bool) (*requests.TxResponse, error)
```

#### Queries

The `websocket` package provides a `Query` struct that allows you to send client [queries](/docs/docs/packages/xrpl#queries) to the server.

#### Examples

##### How to send a payment transaction

This example shows how to send a payment transaction to the XRPL testnet with the `websocket` package.

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/Peersyst/xrpl-go/xrpl/currency"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
)

func main() {

	// Create a new websocket client with a testnet faucet provider
	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.altnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	defer client.Disconnect()

	// Connect to the testnet
	if err := client.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	// Check if the client is connected
	if !client.IsConnected() {
		return
	}

	// Create a new wallet with the ed25519 algorithm
	w, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fund the wallet with the testnet faucet
	if err := client.FundWallet(&w); err != nil {
		fmt.Println(err)
		return
	}

	// Convert the amount to drops
	xrpAmount, err := currency.XrpToDrops("1")
	if err != nil {
		fmt.Println(err)
		return
	}

	xrpAmountInt, err := strconv.ParseInt(xrpAmount, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account: types.Address(w.GetAddress()),
		},
		Destination: "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
		Amount:      types.XRPCurrencyAmount(xrpAmountInt),
		DeliverMax:  types.XRPCurrencyAmount(xrpAmountInt),
	}

	flattenedTx := p.Flatten()

	// Autofill the transaction with the client's config
	if err := client.Autofill(&flattenedTx); err != nil {
		fmt.Println(err)
		return
	}

	// Sign the transaction with the wallet
	txBlob, _, err := w.Sign(flattenedTx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Submit the transaction to the network and wait for it to be included in a ledge
	res, err := client.SubmitAndWait(txBlob, false)
	if err != nil {
		fmt.Println(err)
		return
	}
}

```



## Guides
