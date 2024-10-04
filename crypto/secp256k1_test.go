package crypto

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
			seedBytes, _, err := addresscodec.DecodeSeed(tc.seed)
			if err != nil {
				t.Fatal(err)
			}
			privKey, pubKey, err := SECP256K1().DeriveKeypair(seedBytes, tc.validator)
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
	testCases := []struct {
		name              string
		message           string
		privKey           string
		expectedSignature string
		wantErr           bool
	}{
		{
			name:              "Valid signature",
			message:           "Hello World",
			privKey:           "00B167A9F3B9E60A4F93695713682C102438620AA1785C3AE635F53E5B6261071A",
			expectedSignature: "3045022100E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C829761802206FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64",
			wantErr:           false,
		},
		{
			name:              "Valid signature 2",
			message:           "test",
			privKey:           "00B167A9F3B9E60A4F93695713682C102438620AA1785C3AE635F53E5B6261071A",
			expectedSignature: "30450221008F0B50BEEA0C9787E85EEF9624E9385CCBE565B221BAEC2F2DA5F1D9D6D976F7022022C1B1829AE0E758FB690110F245F15433A0579C44910785FE75F93B9D0FB41F",
			wantErr:           false,
		},
		{
			name:              "Valid signature 3",
			message:           "message",
			privKey:           "00B167A9F3B9E60A4F93695713682C102438620AA1785C3AE635F53E5B6261071A",
			expectedSignature: "3045022100F07CD8D749AAD8F972475A34591336162A959FCC7F8E692D56410CB70B9634F702201B96AF63E166371D8A2C4C3D4CDA69F6064212D1C28D01F598653BE05C323E8C",
			wantErr:           false,
		},
		{
			name:              "Valid signature 4",
			message:           "message2",
			privKey:           "00B167A9F3B9E60A4F93695713682C102438620AA1785C3AE635F53E5B6261071A",
			expectedSignature: "3045022100A2847849BC186B227DB941B1D0A4C39FABBE04A10BF364FC4E394E8B53FD308D02202D47CA9DC35B7FE3E04B578A935CCBE1827B610911709AC13343344F311BD799",
			wantErr:           false,
		},
		{
			name:              "Valid signature 5",
			message:           "message3",
			privKey:           "00B167A9F3B9E60A4F93695713682C102438620AA1785C3AE635F53E5B6261071A",
			expectedSignature: "304402202D5CDBCF251868046CB07FC2CB49200FED9FF216D4B38455A1D222ED29E6123B022057E9962B336D180F0B8DCD99B72C30BB09A5451D2059556E3C1E45C1F5D018B6",
			wantErr:           false,
		},
		{
			name:              "Valid signature 6",
			message:           "message4",
			privKey:           "00B167A9F3B9E60A4F93695713682C102438620AA1785C3AE635F53E5B6261071A",
			expectedSignature: "3045022100A07B597B3D61C3E97B3CCC2DB65F40B26BAEEF7A3EAF8969C0F4E879DDAD1314022058296AC8B4A6E2D5F33891B5BB2211D2AEF1853DF42452649865AB2FE2C83922",
			wantErr:           false,
		},
		{
			name:              "Valid signature 7",
			message:           "message5",
			privKey:           "00B167A9F3B9E60A4F93695713682C102438620AA1785C3AE635F53E5B6261071A",
			expectedSignature: "3044022033950382A62160DBD731D3108C34B07AFD5CD816943931B64E3A25440E8C911902200ABEF5FB3E8B0C4CBD304421B8D3BD6F135D54831FE5426BE74D340ECDFE1F8F",
			wantErr:           false,
		},
		{
			name:              "Empty private key",
			message:           "Hello World",
			privKey:           "",
			expectedSignature: "",
			wantErr:           true,
		},
		{
			name:              "Invalid private key",
			message:           "Hello World",
			privKey:           "invalid_key",
			expectedSignature: "",
			wantErr:           true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			signature, err := SECP256K1().Sign(tc.message, tc.privKey)

			if signature != tc.expectedSignature {
				t.Errorf("sign() returned %v, want %v", signature, tc.expectedSignature)
			}

			if (err != nil) != tc.wantErr {
				t.Errorf("sign() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && len(signature) == 0 {
				t.Errorf("sign() returned empty signature for valid input")
			}
		})
	}
}

func TestSecp256k1_validate(t *testing.T) {
	testCases := []struct {
		name      string
		message   string
		signature string
		pubKey    string
		wantValid bool
	}{
		{
			name:      "Valid signature",
			message:   "Hello World",
			signature: "3045022100E1617F1A3C85B5BC8FA6224F893FE9068BEA8F8D075EE144F6F9D255C829761802206FD9B361CDE83A0C3D5654232F1D7CFB1A614E9A8F9B1A861564029065516E64",
			pubKey:    "02950F4710101A25073BF37086D73FBBD00C7A6B0F91097D8F0BC6D268C400D56E",
			wantValid: true,
		},
		{
			name:      "Valid signature",
			message:   "test",
			signature: "30450221008F0B50BEEA0C9787E85EEF9624E9385CCBE565B221BAEC2F2DA5F1D9D6D976F7022022C1B1829AE0E758FB690110F245F15433A0579C44910785FE75F93B9D0FB41F",
			pubKey:    "02950F4710101A25073BF37086D73FBBD00C7A6B0F91097D8F0BC6D268C400D56E",
			wantValid: true,
		},
		{
			name:      "Invalid signature",
			message:   "Hello, World!",
			signature: "3045022100B1629F44BB12A86AE5A3D79A4E2BE6A473DBBD3F4FB4E3898A2E9A9BE1A54EF502204C3B0C33C46F5ABDE7C2C1A3F2B79B8A9F3A69D8C7C248B2B5C16A39A9C3B5F6",
			pubKey:    "02950F4710101A25073BF37086D73FBBD00C7A6B0F91097D8F0BC6D268C400D56E",
			wantValid: false,
		},
		{
			name:      "Invalid public key",
			message:   "Hello, World!",
			signature: "3045022100B1629F44BB12A86AE5A3D79A4E2BE6A473DBBD3F4FB4E3898A2E9A9BE1A54EF502204C3B0C33C46F5ABDE7C2C1A3F2B79B8A9F3A69D8C7C248B2B5C16A39A9C3B5F5",
			pubKey:    "invalid_key",
			wantValid: false,
		},
		{
			name:      "Empty signature",
			message:   "Hello, World!",
			signature: "",
			pubKey:    "02950F4710101A25073BF37086D73FBBD00C7A6B0F91097D8F0BC6D268C400D56E",
			wantValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid := SECP256K1().Validate(tc.message, tc.pubKey, tc.signature)

			if isValid != tc.wantValid {
				t.Errorf("validate() = %v, want %v", isValid, tc.wantValid)
			}
		})
	}
}
