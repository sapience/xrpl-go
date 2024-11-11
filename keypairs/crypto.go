package keypairs

import (
	"encoding/hex"
	"errors"

	"github.com/Peersyst/xrpl-go/keypairs/interfaces"
	"github.com/Peersyst/xrpl-go/pkg/crypto"
)

// Errors
var (
	ErrInvalidCryptoImplementation = errors.New("not a valid crypto implementation")
)

// GetCryptoImplementationFromKey returns the CryptoImplementation based on the key
func getCryptoImplementationFromKey(k string) interfaces.CryptoImplementation {
	prefix, err := hex.DecodeString(k[:2])
	if err != nil {
		return nil
	}

	if ed25519 := crypto.ED25519(); prefix[0] == ed25519.Prefix() {
		return ed25519
	}
	if secp256k1 := crypto.SECP256K1(); prefix[0] == secp256k1.Prefix() {
		return secp256k1
	}
	return nil
}
