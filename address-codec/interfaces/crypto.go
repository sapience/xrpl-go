package interfaces

type CryptoImplementation interface {
	DeriveKeypair(decodedSeed []byte, validator bool) (string, string, error)
	Sign(msg, privKey string) (string, error)
	Validate(msg, pubkey, sig string) bool
}
