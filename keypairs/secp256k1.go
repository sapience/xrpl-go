package keypairs

import (
	"crypto/sha512"
	"errors"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
)

// var _ CryptoImplementation = (*secp256k1Alg)(nil)

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
	return "", nil
}

func (c *secp256k1Alg) validate(msg, pubkey, sig string) bool {
	return false
}