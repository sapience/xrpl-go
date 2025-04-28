package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	txnTypes "github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/websocket/types"
)

func main() {
	fmt.Println("⏳ Connecting to devnet...")
	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.devnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewDevnetFaucetProvider()),
	)

	defer client.Disconnect()

	if err := client.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	if !client.IsConnected() {
		fmt.Println("❌ Failed to connect to devnet")
		return
	}

	fmt.Println("✅ Connected to devnet")
	fmt.Println()

	fmt.Println("⏳ Funding wallets...")
	// Create and fund the nft wallets
	nftMinter, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println("❌ Error creating nft minter wallet:", err)
		return
	}
	if err := client.FundWallet(&nftMinter); err != nil {
		fmt.Println("❌ Error funding nft minter wallet:", err)
		return
	}
	fmt.Println("💸 NFT minter wallet funded!")

	nftBuyer, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println("❌ Error creating nft buyer wallet:", err)
		return
	}
	if err := client.FundWallet(&nftBuyer); err != nil {
		fmt.Println("❌ Error funding nft buyer wallet:", err)
		return
	}
	fmt.Println("💸 NFT buyer wallet funded!")
	fmt.Println()

	// Mint NFT
	nftMint := transaction.NFTokenMint{
		BaseTx: transaction.BaseTx{
			Account:         nftMinter.ClassicAddress,
			TransactionType: transaction.NFTokenMintTx,
		},
		Destination:  nftBuyer.ClassicAddress,
		Amount:       txnTypes.XRPCurrencyAmount(1000000), // 1 XRP
		NFTokenTaxon: 0,
		URI:          txnTypes.NFTokenURI("68747470733A2F2F676F6F676C652E636F6D"), // https://google.com
	}
	nftMint.SetTransferableFlag()

	responseMint, err := client.SubmitTxAndWait(nftMint.Flatten(), &types.SubmitOptions{
		Autofill: true,
		Wallet:   &nftMinter,
	})
	if err != nil {
		fmt.Println("❌ Error minting NFT:", err)
		return
	}
	if !responseMint.Validated {
		fmt.Println("❌ NFTokenMint txn is not in a validated ledger", responseMint)
		return
	}
	fmt.Println("✅ NFT minted successfully! - 🌎 Hash: ", responseMint.Hash)
	fmt.Println()

	// Get the NFT token offer ID
	metaMap, ok := responseMint.Meta.(map[string]any)
	if !ok {
		fmt.Println("❌ Meta is not a map[string]any")
		return
	}

	offerID, ok := metaMap["offer_id"].(string)
	if !ok {
		fmt.Println("❌ offer_id not found or not a string")
		return
	}

	fmt.Println("🌎 offer_id:", offerID)
	fmt.Println()

	// Accept NFT
	nftAccept := transaction.NFTokenAcceptOffer{
		BaseTx: transaction.BaseTx{
			Account:         nftBuyer.ClassicAddress,
			TransactionType: transaction.NFTokenAcceptOfferTx,
		},
		NFTokenSellOffer: txnTypes.Hash256(offerID),
	}

	responseModify, err := client.SubmitTxAndWait(nftAccept.Flatten(), &types.SubmitOptions{
		Autofill: true,
		Wallet:   &nftBuyer,
	})
	if err != nil {
		fmt.Println("❌ Error accepting NFT offer:", err)
		return
	}
	if !responseModify.Validated {
		fmt.Println("❌ NFTokenAcceptOffer txn is not in a validated ledger", responseModify)
		return
	}
	fmt.Println("✅ NFT offer accepted successfully! - 🌎 Hash: ", responseModify.Hash)
}
