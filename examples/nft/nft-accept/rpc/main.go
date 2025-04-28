package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/xrpl/rpc/types"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	txnTypes "github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

func main() {
	cfg, err := rpc.NewClientConfig(
		"https://s.devnet.rippletest.net:51234/",
		rpc.WithFaucetProvider(faucet.NewDevnetFaucetProvider()),
	)
	if err != nil {
		panic(err)
	}

	client := rpc.NewClient(cfg)

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
