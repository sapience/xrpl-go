package types

import (
	"errors"
	"testing"

	"github.com/Peersyst/xrpl-go/binary-codec/definitions"
	"github.com/Peersyst/xrpl-go/binary-codec/serdes"
	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
	"github.com/Peersyst/xrpl-go/binary-codec/types/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestPermissionValue_FromJSON(t *testing.T) {

	tt := []struct {
		name        string
		input       any
		expected    []byte
		expectedErr error
	}{
		{
			name:        "pass - string permission name",
			input:       "TrustlineAuthorize",
			expected:    []byte{0, 1, 0, 1},
			expectedErr: nil,
		},
		{
			name:        "pass - transaction type name",
			input:       "Payment",
			expected:    []byte{0, 0, 0, 1},
			expectedErr: nil,
		},
		{
			name:        "pass - integer value",
			input:       65537,
			expected:    []byte{0, 1, 0, 1},
			expectedErr: nil,
		},
		{
			name:     "fail - invalid permission name",
			input:    "InvalidPermission",
			expected: nil,
			expectedErr: &definitions.NotFoundError{
				Instance: "DelegatablePermissionName",
				Input:    "InvalidPermission",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			pv := &PermissionValue{}
			actual, err := pv.FromJSON(tc.input)
			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, actual)
			}
		})
	}
}

func TestPermissionValue_ToJSON(t *testing.T) {
	defs := definitions.Get()

	tt := []struct {
		name        string
		input       []byte
		malleate    func(t *testing.T) interfaces.BinaryParser
		expected    any
		expectedErr error
	}{
		{
			name:  "fail - invalid permission value",
			input: []byte{0, 1, 0, 1},
			malleate: func(t *testing.T) interfaces.BinaryParser {
				parserMock := testutil.NewMockBinaryParser(gomock.NewController(t))
				parserMock.EXPECT().ReadBytes(gomock.Any()).Return([]byte{}, errors.New("binary parser has no data"))
				return parserMock
			},
			expected:    nil,
			expectedErr: errors.New("binary parser has no data"),
		},
		{
			name:  "pass - known permission value returns name",
			input: []byte{0, 1, 0, 1},
			malleate: func(t *testing.T) interfaces.BinaryParser {
				return serdes.NewBinaryParser([]byte{0, 1, 0, 1}, defs)
			},
			expected:    "TrustlineAuthorize",
			expectedErr: nil,
		},
		{
			name:  "pass - unknown permission value returns number",
			input: []byte{0, 0, 0, 100},
			malleate: func(t *testing.T) interfaces.BinaryParser {
				return serdes.NewBinaryParser([]byte{0, 0, 0, 100}, defs)
			},
			expected:    uint32(100),
			expectedErr: nil,
		},
		{
			name:  "pass - Payment value returns Payment name",
			input: []byte{0, 0, 0, 1},
			malleate: func(t *testing.T) interfaces.BinaryParser {
				return serdes.NewBinaryParser([]byte{0, 0, 0, 1}, defs)
			},
			expected:    "Payment",
			expectedErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			pv := &PermissionValue{}
			parser := tc.malleate(t)
			actual, err := pv.ToJSON(parser)
			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.Equal(t, tc.expected, actual)
			}
		})
	}
}
