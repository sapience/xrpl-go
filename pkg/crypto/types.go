package crypto

type CryptoAlgorithm struct {
	prefix           uint8
	familySeedPrefix uint8
}

func (c CryptoAlgorithm) Prefix() uint8 {
	return c.prefix
}

func (c CryptoAlgorithm) FamilySeedPrefix() uint8 {
	return c.familySeedPrefix
}

func (c CryptoAlgorithm) DeriveKeypair(decodedSeed []byte, validator bool) (string, string, error) {
	return "", "", nil
}

func (c CryptoAlgorithm) Sign(msg, privKey string) (string, error) {
	return "", nil
}

func (c CryptoAlgorithm) Validate(msg, pubkey, sig string) bool {
	return false
}
