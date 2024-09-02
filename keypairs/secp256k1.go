package keypairs

import (
	"crypto/sha512"
	"math/big"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/btcsuite/btcd/btcec/v2"
)

var _ CryptoImplementation = (*secp256k1Alg)(nil)

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

func (c *secp256k1Alg) derivateKeypair(seed []byte, validator bool, accountIndex uint32) (string, string, error) {
	curve := btcec.S256()
	order := curve.N

	privateGen := deriveScalar(seed, nil)

	if validator {
		return "", "", &secp256k1ValidatorError{}
	}

	rootPrivateKey, _ := btcec.PrivKeyFromBytes(privateGen.Bytes())

	derivatedScalar := deriveScalar(rootPrivateKey.PubKey().SerializeCompressed(), big.NewInt(int64(accountIndex)))
	scalarWithPrivateGen := derivatedScalar.Add(derivatedScalar, privateGen)
	privateKey := scalarWithPrivateGen.Mod(scalarWithPrivateGen, order)

	privKeyBytes := privateKey.Bytes()
	privKeyBytes = append([]byte{addresscodec.SECP256K1}, privKeyBytes...)
	private := formatKey(privKeyBytes)

	_, pubKey := btcec.PrivKeyFromBytes(privKeyBytes)
	

	


	return private, public, nil
}

func (c *secp256k1Alg) sign(msg, privKey string) (string, error) {
	return "", nil
}

func (c *secp256k1Alg) validate(msg, pubkey, sig string) bool {
	return false
}