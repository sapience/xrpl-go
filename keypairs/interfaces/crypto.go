package interfaces

type KeypairCryptoImplementation interface {
	DeriveKeypair(decodedSeed []byte, validator bool) (string, string, error)
	Sign(msg, privKey string) (string, error)
	Validate(msg, pubkey, sig string) bool
}

type NodeDerivationCryptoImplementation interface {
	DerivePublicKeyFromPublicGenerator(pubKey []byte) ([]byte, error)
}
