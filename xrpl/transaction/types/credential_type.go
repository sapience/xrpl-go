package types

type CredentialType string

// String returns the string representation of a CredentialType.
func (c *CredentialType) String() string {
	return string(*c)
}
