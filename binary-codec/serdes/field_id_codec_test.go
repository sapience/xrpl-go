//go:build unit
// +build unit

package serdes

import (
	"encoding/hex"
	"testing"

	"github.com/Peersyst/xrpl-go/binary-codec/definitions"
	"github.com/stretchr/testify/require"
)

// Returns the field name represented by the given field ID in hex string form.
func decodeFieldID(h string) (string, error) {
	b, err := hex.DecodeString(h)
	if err != nil {
		return "", err
	}
	if len(b) == 1 {
		return definitions.Get().GetFieldNameByFieldHeader(definitions.CreateFieldHeader(int32(b[0]>>4), int32(b[0]&byte(15))))
	}
	if len(b) == 2 {
		firstByteHighBits := b[0] >> 4
		firstByteLowBits := b[0] & byte(15)
		if firstByteHighBits == 0 {
			return definitions.Get().GetFieldNameByFieldHeader(definitions.CreateFieldHeader(int32(b[1]), int32(firstByteLowBits)))
		}
		return definitions.Get().GetFieldNameByFieldHeader(definitions.CreateFieldHeader(int32(firstByteHighBits), int32(b[1])))
	}
	if len(b) == 3 {
		return definitions.Get().GetFieldNameByFieldHeader(definitions.CreateFieldHeader(int32(b[1]), int32(b[2])))
	}
	return "", nil
}

func TestEncodeFieldID(t *testing.T) {
	tt := []struct {
		description string
		input       string
		expected    []byte
		expectedErr error
	}{
		{
			description: "Type Code and Field Code < 16",
			input:       "Sequence",
			expected:    []byte{36},
			expectedErr: nil,
		},
		{
			description: "Additional Type Code and Field Code < 16",
			input:       "Flags",
			expected:    []byte{34},
			expectedErr: nil,
		},
		{
			description: "Additional Type Code and Field Code < 16",
			input:       "DestinationTag",
			expected:    []byte{46},
			expectedErr: nil,
		},
		{
			description: "Type Code >= 16 and Field Code < 16",
			input:       "Paths",
			expected:    []byte{1, 18},
			expectedErr: nil,
		},
		{
			description: "Additional Type Code >= 16 and Field Code < 16",
			input:       "CloseResolution",
			expected:    []byte{1, 16},
			expectedErr: nil,
		},
		{
			description: "Type Code < 16 and Field Code >= 16",
			input:       "SetFlag",
			expected:    []byte{32, 33},
			expectedErr: nil,
		},
		{
			description: "Additional Type Code < 16 and Field Code >= 16",
			input:       "Nickname",
			expected:    []byte{80, 18},
			expectedErr: nil,
		},
		{
			description: "Type Code and Field Code >= 16",
			input:       "TickSize",
			expected:    []byte{0, 16, 16},
			expectedErr: nil,
		},
		{
			description: "Additional Type Code and Field Code >= 16",
			input:       "UNLModifyDisabling",
			expected:    []byte{0, 16, 17},
			expectedErr: nil,
		},
		{
			description: "Non existent field name",
			input:       "yurt",
			expected:    nil,
			expectedErr: &definitions.NotFoundError{Instance: "FieldName", Input: "yurt"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got, err := encodeFieldID(tc.input)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				require.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, got)
			}
		})
	}
}

func TestDecodeFieldID(t *testing.T) {
	tt := []struct {
		description string
		input       []byte
		expected    string
		expectedErr error
	}{
		{
			description: "Decode Sequence fieldId (Type Code and Field Code < 16)",
			input:       []byte{36},
			expected:    "Sequence",
			expectedErr: nil,
		},
		{
			description: "Decode DestinationTag fieldId (Type Code and Field Code < 16)",
			input:       []byte{46},
			expected:    "DestinationTag",
			expectedErr: nil,
		},
		{
			description: "Decode Paths fieldId (Type Code >= 16 and Field Code < 16)",
			input:       []byte{1, 18},
			expected:    "Paths",
			expectedErr: nil,
		},
		{
			description: "Decode CloseResolution fieldId (Type Code >= 16 and Field Code < 16)",
			input:       []byte{1, 16},
			expected:    "CloseResolution",
			expectedErr: nil,
		},
		{
			description: "Decode SetFlag fieldId (Type Code < 16 and Field Code >= 16)",
			input:       []byte{32, 33},
			expected:    "SetFlag",
			expectedErr: nil,
		},
		{
			description: "Decode Nickname fieldId (Type Code < 16 and Field Code >= 16)",
			input:       []byte{80, 18},
			expected:    "Nickname",
			expectedErr: nil,
		},
		{
			description: "Decode TickSize fieldId (Type Code and Field Code >= 16)",
			input:       []byte{0, 16, 16},
			expected:    "TickSize",
			expectedErr: nil,
		},
		{
			description: "Decode UNLModifyDisabling fieldId (Type Code and Field Code >= 16)",
			input:       []byte{0, 16, 17},
			expected:    "UNLModifyDisabling",
			expectedErr: nil,
		},
		{
			description: "Non existent field name",
			input:       []byte{255},
			expected:    "",
			expectedErr: &definitions.NotFoundErrorFieldHeader{Instance: "FieldHeader"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			hex := hex.EncodeToString(tc.input)
			// fmt.Println("hex string:", hex)
			actual, err := decodeFieldID(hex)
			// fmt.Println(actual)

			if tc.expectedErr != nil {
				require.Error(t, err, tc.expectedErr.Error())
				require.Zero(t, actual)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, actual)
			}
		})
	}
}
