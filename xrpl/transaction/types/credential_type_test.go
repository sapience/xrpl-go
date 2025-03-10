package types

import (
	"testing"
)

func TestCredentialType_String(t *testing.T) {
	tests := []struct {
		name     string
		input    CredentialType
		expected string
	}{
		{
			name:     "pass - valid CredentialType",
			input:    CredentialType("my_credential"),
			expected: "my_credential",
		},
		{
			name:     "pass - empty URI",
			input:    CredentialType(""),
			expected: "",
		},
		{
			name:     "pass - CredentialType with special characters",
			input:    CredentialType("https://example.com/nft/1?query=param&another=param"),
			expected: "https://example.com/nft/1?query=param&another=param",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.String(); got != tt.expected {
				t.Errorf("CredentialType.String(), got: %v but we want: %v", got, tt.expected)
			}
		})
	}
}
