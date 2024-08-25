package xrpl

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/keypairs"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// A utility for deriving a wallet composed of a keypair (publicKey/privateKey).
// A wallet can be derived from either a seed, mnemonic, or entropy (array of random numbers).
// It provides functionality to sign/verify transactions offline.
type Wallet struct {
	PublicKey      string
	PrivateKey     string
	ClassicAddress string
	Seed           string
}

// creates a new random Wallet. In order to make this a valid account on ledger, you must
// Send XRP to it.
func NewWallet(alg addresscodec.CryptoAlgorithm) (Wallet, error) {
	seed, err := keypairs.GenerateSeed("", alg)
	if err != nil {
		return Wallet{}, err
	}
	return NewWalletFromSeed(seed, "")
}

// Derives a wallet from a seed.
// Returns a Wallet object. If an error occurs, it will be returned.
func NewWalletFromSeed(seed string, masterAddress string) (Wallet, error) {
	privKey, pubKey, err := keypairs.DeriveKeypair(seed, false)
	if err != nil {
		return Wallet{}, err
	}

	var classicAddr string
	if ok := addresscodec.IsValidClassicAddress(masterAddress); ok {
		classicAddr = masterAddress
	} else {
		addr, err := keypairs.DeriveClassicAddress(pubKey)
		if err != nil {
			return Wallet{}, err
		}
		classicAddr = addr
	}

	return Wallet{
		PublicKey:      pubKey,
		PrivateKey:     privKey,
		Seed:           seed,
		ClassicAddress: classicAddr,
	}, nil

}

// Derives a wallet from a secret (AKA a seed).
// Returns a Wallet object. If an error occurs, it will be returned.
func NewWalletFromSecret(seed string) (Wallet, error) {
	return NewWalletFromSeed(seed, "")
}

// Derives a wallet from a bip39 or RFC1751 mnemonic (Defaults to bip39).
// Returns a Wallet object. If an error occurs, it will be returned.
func NewWalletFromMnemonic(mnemonic string) (*Wallet, error) {
	// Validate the mnemonic
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("invalid mnemonic")
	}

	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "")

	// Derive the master key
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	// Derive the key using the path m/44'/144'/0'/0/0
	path := []uint32{
		44 + bip32.FirstHardenedChild,
		144 + bip32.FirstHardenedChild,
		bip32.FirstHardenedChild,
		0,
		0,
	}

	key := masterKey
	for _, childNum := range path {
		key, err = key.NewChildKey(childNum)
		if err != nil {
			return nil, err
		}
	}

	// Convert the private key to the format expected by the XRPL library
	privKey := strings.ToUpper(hex.EncodeToString(key.Key))
	pubKey := strings.ToUpper(hex.EncodeToString(key.PublicKey().Key))

	// Derive classic address
	classicAddr, err := keypairs.DeriveClassicAddress(pubKey)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		PublicKey:      pubKey,
		PrivateKey:     fmt.Sprintf("00%s", privKey),
		ClassicAddress: classicAddr,
		Seed:           "", // We don't have the seed in this case
	}, nil
}

// Signs a transaction offline.
// In order for a transaction to be validated, it must be signed by the account sending the transaction to prove
// that the owner is actually the one deciding to take that action.
func (w *Wallet) Sign(tx transactions.BaseTx) (txBlob, hash string, err error) {
	return "", "", errors.New("not implemented")

	// encodedTx, _ := binarycodec.EncodeForSigning(tx)
	// hexTx, err := hex.DecodeString(encodedTx)
	// if err != nil {
	// 	return "", "", err
	// }

	// hash, err = keypairs.Sign(string(hexTx), w.PrivateKey)
	// if err != nil {
	// 	return "", "", err
	// }

	// tx.TxnSignature = hash

	// txBlob, err = binarycodec.Encode(tx)
	// if err != nil {
	// 	return "", "", err
	// }
	// return txBlob, hash, nil
}
