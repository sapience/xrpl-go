package crypto

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"strings"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
)

const (
	ed25519Prefix = 0xED
)

type ed25519CryptoAlgorithm CryptoAlgorithm

func ED25519() ed25519CryptoAlgorithm {
	return ed25519CryptoAlgorithm{
		prefix: ed25519Prefix,
	}
}

func (c ed25519CryptoAlgorithm) Prefix() byte {
	return c.prefix
}

func (c ed25519CryptoAlgorithm) FamilySeedPrefix() byte {
	return c.familySeedPrefix
}

func (c ed25519CryptoAlgorithm) DeriveKeypair(decodedSeed []byte, validator bool) (string, string, error) {
	if validator {
		return "", "", &ed25519ValidatorError{}
	}
	rawPriv := Sha512Half(decodedSeed)
	pubKey, privKey, err := ed25519.GenerateKey(bytes.NewBuffer(rawPriv))
	if err != nil {
		return "", "", err
	}
	pubKey = append([]byte{addresscodec.ED25519}, pubKey...)
	public := strings.ToUpper(hex.EncodeToString(pubKey))
	privKey = append([]byte{addresscodec.ED25519}, privKey...)
	private := strings.ToUpper(hex.EncodeToString(privKey[:32+len([]byte{addresscodec.ED25519})]))
	return private, public, nil
}

func (c ed25519CryptoAlgorithm) Sign(msg, privKey string) (string, error) {
	b, err := hex.DecodeString(privKey)
	if err != nil {
		return "", err
	}
	rawPriv := ed25519.NewKeyFromSeed(b[1:])
	signedMsg := ed25519.Sign(rawPriv, []byte(msg))
	return strings.ToUpper(hex.EncodeToString(signedMsg)), nil
}

func (c ed25519CryptoAlgorithm) Validate(msg, pubkey, sig string) bool {
	bp, err := hex.DecodeString(pubkey)
	if err != nil {
		return false
	}

	bs, err := hex.DecodeString(sig)
	if err != nil {
		return false
	}

	return ed25519.Verify(ed25519.PublicKey(bp[1:]), []byte(msg), bs)
}

type ed25519ValidatorError struct{}

func (e *ed25519ValidatorError) Error() string {
	return "validator keypairs can not use Ed25519"
}
