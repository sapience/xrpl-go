package interfaces

type KeypairCryptoAlg interface {
	DeriveKeypair(decodedSeed []byte, validator bool) (string, string, error)
	Sign(msg, privKey string) (string, error)
	Validate(msg, pubkey, sig string) bool
}

type NodeDerivationCryptoAlg interface {
	DerivePublicKeyFromPublicGenerator(pubKey []byte) ([]byte, error)
}
