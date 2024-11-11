package keypairs

import (
	"crypto/rand"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/pkg/random"
)

var r random.Randomizer

const (
	VERIFICATIONMESSAGE = "This test message should verify."
)

func init() {
	r.Reader = rand.Reader
}

func GenerateSeed(entropy string, alg CryptoImplementation) (string, error) {
	var pe []byte
	if entropy == "" {
		b, err := r.GenerateBytes(addresscodec.FamilySeedLength)
		pe = b
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
		return
	}
	private, public, err = alg.DeriveKeypair(ds, validator)
	if err != nil {
		return
	}
	signature, err := alg.Sign(VERIFICATIONMESSAGE, private)

	if !alg.Validate(VERIFICATIONMESSAGE, public, signature) {
		return "", "", &InvalidSignatureError{}
	}
	return
}

func DeriveClassicAddress(pubkey string) (string, error) {
	return addresscodec.EncodeClassicAddressFromPublicKeyHex(pubkey)
}

func Sign(msg, privKey string) (string, error) {
	alg := getCryptoImplementationFromKey(privKey)
	if alg == nil {
		return "", &CryptoImplementationError{}
	}
	return alg.Sign(msg, privKey)
}

func Validate(msg, pubKey, sig string) (bool, error) {
	alg := getCryptoImplementationFromKey(pubKey)
	if alg == nil {
		return false, &CryptoImplementationError{}
	}
	return alg.Validate(msg, pubKey, sig), nil
}

type InvalidSignatureError struct{}

func (e *InvalidSignatureError) Error() string {
	return "derived keypair did not generate verifiable signature"
}
