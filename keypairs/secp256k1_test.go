package keypairs

import (
	"errors"
	"testing"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
)

func TestSecp256k1_deriveKeypair(t *testing.T) {
	testCases := []struct {
		name            string
		seed            string
		validator       bool
		expectedPrivKey string
		expectedPubKey  string
		expectedErr     error
	}{
		{
			name:            "valid seed (1)",
			seed:            "sntbkd2DsouBx8BAdJdi35p1HRw6h",
			validator:       false,
			expectedPubKey:  "02950F4710101A25073BF37086D73FBBD00C7A6B0F91097D8F0BC6D268C400D56E",
			expectedPrivKey: "00B167A9F3B9E60A4F93695713682C102438620AA1785C3AE635F53E5B6261071A",
			expectedErr:     nil,
		},
		{
			name:            "valid seed (2)",
			seed:            "shSDdnXqsS7zAjbdWX86fT6H5oCxK",
			validator:       false,
			expectedPubKey:  "031FBCFDD2EC6C2EDFBBA3866BDBAC28E5253C6A01FE9EFF8CAAE01871F009E837",
			expectedPrivKey: "00A3D1513DBE784107428B363A1F8EAF1377AB63D4D137AB9E28E0BC614C71D8C0",
			expectedErr:     nil,
		},
		{
			name:            "validator set to true",
			seed:            "shSDdnXqsS7zAjbdWX86fT6H5oCxK",
			validator:       true,
			expectedPubKey:  "",
			expectedPrivKey: "",
			expectedErr:     errors.New("validator keypair derivation not supported"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			secp256k1 := secp256k1Alg{}
			seedBytes, _, err := addresscodec.DecodeSeed(tc.seed)
			if err != nil {
				t.Fatal(err)
			}
			privKey, pubKey, err := secp256k1.deriveKeypair(seedBytes, tc.validator)
			if tc.expectedErr != nil {
				if err == nil || err.Error() != tc.expectedErr.Error() {
					t.Fatalf("expected error %v, got %v", tc.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("expected valid keypair, got error: %v", err)
				}
				if privKey != tc.expectedPrivKey {
					t.Errorf("expected private key %s, got %s", tc.expectedPrivKey, privKey)
				}
				if pubKey != tc.expectedPubKey {
					t.Errorf("expected public key %s, got %s", tc.expectedPubKey, pubKey)
				}
			}
		})
	}
}


func TestSecp256k1_sign(t *testing.T) {

	secp256k1 := secp256k1Alg{}
	seed, _, err := addresscodec.DecodeSeed("sntbkd2DsouBx8BAdJdi35p1HRw6h")
	if err != nil {
		t.Fatal(err)
	}
	privKey, _, _ := secp256k1.deriveKeypair(seed, false)
	signature, err := secp256k1.sign("Hello World", privKey)
	
	if err != nil {
		t.Fatal(err)
	}
	
	if signature != "3045022100E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C829761802206FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64" {
		t.Errorf("invalid signature %s", signature)
	}
}