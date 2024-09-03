package keypairs

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/Peersyst/xrpl-go/pkg/secp256k1"
	"github.com/btcsuite/btcd/btcec/v2"

	"golang.org/x/crypto/sha3"
)

var _ CryptoImplementation = (*secp256k1Alg)(nil) 

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	b := make([]byte, 32)
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	d.Sum(b)
	return b
}

func Keccak256Hash(data ...[]byte) (h [32]byte) {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	d.Sum(h[:])
	return h
}


type secp256k1Alg struct{}

func deriveScalar(bytes []byte, discrim *big.Int) *big.Int {

	order := btcec.S256().N
	for i := 0; i <= 0xffffffff; i++ {
		hash := sha512.New()
		
		hash.Write(bytes)

		if discrim != nil {
			discrimBytes := make([]byte, 4)
			bytes[0] = byte(uint32(discrim.Uint64()))
			bytes[1] = byte(uint32(discrim.Uint64()) >> 8)
			bytes[2] = byte(uint32(discrim.Uint64()) >> 16)
			bytes[3] = byte(uint32(discrim.Uint64()) >> 24)

			hash.Write(discrimBytes)
		}

		shiftBytes := make([]byte, 4)
		bytes[0] = byte(uint32(i))
		bytes[1] = byte(uint32(i) >> 8)
		bytes[2] = byte(uint32(i) >> 16)
		bytes[3] = byte(uint32(i) >> 24)

		hash.Write(shiftBytes)

		key := new(big.Int).SetBytes(hash.Sum(nil)[:32])

		if key.Cmp(big.NewInt(0)) > 0 && key.Cmp(order) < 0 {
			return key
		}
	}
	// This error is practically impossible to reach.
	// The order of the curve describes the (finite) amount of points on the curve.
	panic("impossible unicorn ;)")
}

func (c *secp256k1Alg) deriveKeypair(seed []byte, validator bool) (string, string, error) {
	curve := btcec.S256()
	order := curve.N

	privateGen := deriveScalar(seed, nil)

	if validator {
		return "", "", errors.New("validator keypair derivation not supported")
	}

	rootPrivateKey, _ := btcec.PrivKeyFromBytes(privateGen.Bytes())

	derivatedScalar := deriveScalar(rootPrivateKey.PubKey().SerializeCompressed(), big.NewInt(0))
	scalarWithPrivateGen := derivatedScalar.Add(derivatedScalar, privateGen)
	privateKey := scalarWithPrivateGen.Mod(scalarWithPrivateGen, order)

	privKeyBytes := privateKey.Bytes()
	private := formatKey(privKeyBytes)

	_, pubKey := btcec.PrivKeyFromBytes(privKeyBytes)
	
	pubKeyBytes := pubKey.SerializeCompressed()
	public := formatKey(pubKeyBytes)

	private = "00" + private
	return private, public, nil
}

func (c *secp256k1Alg) sign(msg, privKey string) (string, error) {
	if len(privKey) == 66 {
		privKey = privKey[2:]
	}
	fmt.Println("privKey", privKey)
	privKeyBytes, err := hex.DecodeString(privKey)
	if err != nil {
		return "", err
	}
	// Convert msg to hex
	msgHex := hex.EncodeToString([]byte(msg))
	hash := Keccak256Hash([]byte(msgHex))

	sig, err := secp256k1.Sign(hash[:], privKeyBytes)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(sig)), nil
}

func (c *secp256k1Alg) validate(msg, pubkey, sig string) bool {
	return false
}