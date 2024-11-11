package keypairs

import (
	"errors"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/keypairs/interfaces"
)

var (
	// Errors

	// Derived keypair did not generate verifiable signature
	ErrInvalidSignature = errors.New("derived keypair did not generate verifiable signature")
)

const (
	verificationMessage = "This test message should verify."
)

func GenerateSeed(entropy string, alg interfaces.CryptoImplementation, r interfaces.Randomizer) (string, error) {
	var pe []byte
	var err error
	if entropy == "" {
		pe, err = r.GenerateBytes(addresscodec.FamilySeedLength)
		if err != nil {
			return "", err
		}
	} else {
		pe = []byte(entropy)[:addresscodec.FamilySeedLength]
	}
	return addresscodec.EncodeSeed(pe, alg)
}

// Derives a keypair from a given seed. Returns a tuple of private key and public key
func DeriveKeypair(seed string, validator bool) (private, public string, err error) {
	ds, alg, err := addresscodec.DecodeSeed(seed)
	if err != nil {
		return "", "", err
	}
	private, public, err = alg.DeriveKeypair(ds, validator)
	if err != nil {
		return "", "", err
	}
	signature, err := alg.Sign(verificationMessage, private)
	if err != nil {
		return "", "", err
	}
	if !alg.Validate(verificationMessage, public, signature) {
		return "", "", ErrInvalidSignature
	}
	return private, public, nil
}

func DeriveClassicAddress(pubkey string) (string, error) {
	return addresscodec.EncodeClassicAddressFromPublicKeyHex(pubkey)
}

func Sign(msg, privKey string) (string, error) {
	alg := getCryptoImplementationFromKey(privKey)
	if alg == nil {
		return "", ErrInvalidCryptoImplementation
	}
	return alg.Sign(msg, privKey)
}

func Validate(msg, pubKey, sig string) (bool, error) {
	alg := getCryptoImplementationFromKey(pubKey)
	if alg == nil {
		return false, ErrInvalidCryptoImplementation
	}
	return alg.Validate(msg, pubKey, sig), nil
}
