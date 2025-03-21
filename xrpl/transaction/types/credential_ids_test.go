package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCredentialIDs_IsValid(t *testing.T) {
	tests := []struct {
		name          string
		credentialIDs CredentialIDs
		expected      bool
	}{
		{
			name:          "empty",
			credentialIDs: CredentialIDs{},
			expected:      false,
		},
		{
			name:          "valid",
			credentialIDs: CredentialIDs{"0000000000000000000000000000000000000000000000000000000000000000"},
			expected:      true,
		},
		{
			name: "invalid",
			credentialIDs: CredentialIDs{
				"0000000000000000000000000000000000000000000000000000000000000000",
				"6D795F63726564656E7469616C",
			},
			expected: true,
		},
		{
			name: "invalid",
			credentialIDs: CredentialIDs{
				"0000000000000000000000000000000000000000000000000000000000000000",
				"invalid",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.credentialIDs.IsValid()
			require.Equal(t, test.expected, result)
		})
	}
}
