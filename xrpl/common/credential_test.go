package common

import (
	"strings"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestIsValidCredentialType(t *testing.T) {
	tests := []struct {
		name     string
		credType types.CredentialType
		expected bool
	}{
		{
			name:     "empty string",
			credType: "",
			expected: false,
		},
		{
			name:     "valid hex string",
			credType: "6D795F63726564656E7469616C",
			expected: true,
		},
		{
			name:     "invalid hex string",
			credType: "invalid",
			expected: false,
		},
		{
			name:     "short hex string",
			credType: types.CredentialType(strings.Repeat("0", MinCredentialTypeLength-1)),
			expected: false,
		},
		{
			name:     "long hex string",
			credType: types.CredentialType(strings.Repeat("0", MaxCredentialTypeLength+1)),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsValidCredentialType(test.credType)
			assert.Equal(t, test.expected, result)
		})
	}
}
