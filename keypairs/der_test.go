package keypairs

import (
	"encoding/hex"
	"strings"
	"testing"
)

func BenchmarkDERHexToSig(b *testing.B) {
	hexSignature := "3045022100E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C829761802206FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64"
	for i := 0; i < b.N; i++ {
		DERHexToSig(hexSignature)
	}
}

func TestDERHexToSig(t *testing.T) {
	testCases := []struct {
		name         string
		hexSignature string
		expectedR    string
		expectedS    string
		expectError  bool
	}{
		{
			name:         "Valid DER signature",
			hexSignature: "3045022100E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C829761802206FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64",
			expectedR:    "E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C8297618",
			expectedS:    "6FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64",
			expectError:  false,
		},
		{
			name:         "Invalid hex string",
			hexSignature: "invalid",
			expectedR:    "",
			expectedS:    "",
			expectError:  true,
		},
		{
			name:         "Invalid signature tag",
			hexSignature: "3145022100E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C829761802206FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64",
			expectedR:    "",
			expectedS:    "",
			expectError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, s, err := DERHexToSig(tc.hexSignature)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if hex.EncodeToString(r) != strings.ToLower(tc.expectedR) {
					t.Errorf("Expected R %s, but got %s", tc.expectedR, hex.EncodeToString(r))
				}
				if hex.EncodeToString(s) != strings.ToLower(tc.expectedS) {
					t.Errorf("Expected S %s, but got %s", tc.expectedS, hex.EncodeToString(s))
				}
			}
		})
	}
}

func BenchmarkDERHexFromSig(b *testing.B) {
	rHex := "E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C8297618"
	sHex := "6FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64"
	for i := 0; i < b.N; i++ {
		DERHexFromSig(rHex, sHex)
	}
}

func TestDERHexFromSig(t *testing.T) {
	testCases := []struct {
		name        string
		rHex        string
		sHex        string
		expectedDER string
		expectError bool
	}{
		{
			name:        "Valid r and s values",
			rHex:        "E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C8297618",
			sHex:        "6FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64",
			expectedDER: strings.ToLower("3045022100E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C829761802206FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64"),
			expectError: false,
		},
		{
			name:        "Invalid r hex string",
			rHex:        "invalid",
			sHex:        "6FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64",
			expectedDER: "",
			expectError: true,
		},
		{
			name:        "Invalid s hex string",
			rHex:        "E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C8297618",
			sHex:        "invalid",
			expectedDER: "",
			expectError: true,
		},
		{
			name:        "r value with leading zero",
			rHex:        "00E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C8297618",
			sHex:        "6FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64",
			expectedDER: strings.ToLower("3045022100E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C829761802206FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64"),
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := DERHexFromSig(tc.rHex, tc.sHex)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tc.expectedDER {
					t.Errorf("Expected %s, but got %s", tc.expectedDER, result)
				}
			}
		})
	}
}
