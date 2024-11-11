package crypto

type Algorithm struct {
	prefix           uint8
	familySeedPrefix uint8
}

func (c Algorithm) Prefix() uint8 {
	return c.prefix
}

func (c Algorithm) FamilySeedPrefix() uint8 {
	return c.familySeedPrefix
}

func (c Algorithm) DeriveKeypair(_ []byte, _ bool) (string, string, error) {
	return "", "", nil
}

func (c Algorithm) Sign(_ string, _ string) (string, error) {
	return "", nil
}

func (c Algorithm) Validate(_ string, _ string, _ string) bool {
	return false
}
