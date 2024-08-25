package xrpl

import "testing"

func TestNewWalletFromSeed(t *testing.T) {
	testCases := []struct {
		Seed           string
		PublicKey      string
		PrivateKey     string
		ClassicAddress string
	}{
		{
			Seed:           "sEd7io6yt5dFJrcePgRiFVHvmkJhJD1",
			PublicKey:      "EDC9DA1AA7513D891B58B3C9BEBAE3EB12620AFF4ABBA806B23BB3FA62109CE87F",
			PrivateKey:     "EDE01A1644C9FDE0367A7A285CA69798066C131C1133E1128B170CA65AEA5C6D19",
			ClassicAddress: "rn5M6BQCmQAzBxms9A84qEpx1Fdn9y7jdD",
		},
		{
			Seed:           "sEdTLE1G6QVc8znymeRZD3s5oajQcY5",
			PublicKey:      "ED676AD70576E126B46F6AF52D908FAB8F352F3A0BA05F48613DF017F6B83205B6",
			PrivateKey:     "ED31053AC2D97A74EA3A401CF23EEDA20400500CCFB82F442A8E0E6096A4150A9F",
			ClassicAddress: "r9D79PwpT5Z5xztgiQQgmxcYbF249PefnW",
		},
		{
			Seed:           "sEd71D6u2LkA36TMfJ5rApVsgZXQE9F",
			PublicKey:      "ED0A8A14F3226B2109047662605898F96F61764A9269B7823453A04A7B4F524C0E",
			PrivateKey:     "ED2C0EAB27E1411DBB8FACC88D531A69967DA0E45AC7821A4041A5AEE24BB8FF29",
			ClassicAddress: "rs7cvHcsEF54DEs2y24Tpph3Xf71xUUrFu",
		},
	}

	for _, tc := range testCases {
		wallet, err := NewWalletFromSeed(tc.Seed, "")
		if err != nil {
			t.Errorf("Error generating wallet from seed: %s", err)
		}

		if wallet.PublicKey != tc.PublicKey {
			t.Errorf("Public key does not match expected value. Expected: %s, got: %s", tc.PublicKey, wallet.PublicKey)
		}

		if wallet.PrivateKey != tc.PrivateKey {
			t.Errorf("Private key does not match expected value. Expected: %s, got: %s", tc.PrivateKey, wallet.PrivateKey)
		}

		if wallet.ClassicAddress != tc.ClassicAddress {
			t.Errorf("Classic address does not match expected value. Expected: %s, got: %s", tc.ClassicAddress, wallet.ClassicAddress)
		}
	}
}

func TestNewWalletFromSecret(t *testing.T) {
	testCases := []struct {
		Seed           string
		PublicKey      string
		PrivateKey     string
		ClassicAddress string
	}{
		{
			Seed:           "sEd7io6yt5dFJrcePgRiFVHvmkJhJD1",
			PublicKey:      "EDC9DA1AA7513D891B58B3C9BEBAE3EB12620AFF4ABBA806B23BB3FA62109CE87F",
			PrivateKey:     "EDE01A1644C9FDE0367A7A285CA69798066C131C1133E1128B170CA65AEA5C6D19",
			ClassicAddress: "rn5M6BQCmQAzBxms9A84qEpx1Fdn9y7jdD",
		},
		{
			Seed:           "sEdTLE1G6QVc8znymeRZD3s5oajQcY5",
			PublicKey:      "ED676AD70576E126B46F6AF52D908FAB8F352F3A0BA05F48613DF017F6B83205B6",
			PrivateKey:     "ED31053AC2D97A74EA3A401CF23EEDA20400500CCFB82F442A8E0E6096A4150A9F",
			ClassicAddress: "r9D79PwpT5Z5xztgiQQgmxcYbF249PefnW",
		},
		{
			Seed:           "sEd71D6u2LkA36TMfJ5rApVsgZXQE9F",
			PublicKey:      "ED0A8A14F3226B2109047662605898F96F61764A9269B7823453A04A7B4F524C0E",
			PrivateKey:     "ED2C0EAB27E1411DBB8FACC88D531A69967DA0E45AC7821A4041A5AEE24BB8FF29",
			ClassicAddress: "rs7cvHcsEF54DEs2y24Tpph3Xf71xUUrFu",
		},
	}

	for _, tc := range testCases {
		wallet, err := NewWalletFromSecret(tc.Seed)
		if err != nil {
			t.Errorf("Error generating wallet from seed: %s", err)
		}

		if wallet.PublicKey != tc.PublicKey {
			t.Errorf("Public key does not match expected value. Expected: %s, got: %s", tc.PublicKey, wallet.PublicKey)
		}

		if wallet.PrivateKey != tc.PrivateKey {
			t.Errorf("Private key does not match expected value. Expected: %s, got: %s", tc.PrivateKey, wallet.PrivateKey)
		}

		if wallet.ClassicAddress != tc.ClassicAddress {
			t.Errorf("Classic address does not match expected value. Expected: %s, got: %s", tc.ClassicAddress, wallet.ClassicAddress)
		}
	}
}

func TestNewWalletFromMnemonic(t *testing.T) {
	testCases := []struct {
		Mnemonic       string
		PublicKey      string
		PrivateKey     string
		ClassicAddress string
	}{
		{
			Mnemonic: "midnight help already frost arena force omit physical please dwarf envelope royal dice surge eight often muscle tired blast begin waste fat rescue debate",
			PublicKey: "028E831F16FD85ABEDA7577B6F4F26500FAB80AEA54B8A89EEC6FA44BCC7AF5678",
			PrivateKey: "00C503FC86436D384F37F946E8DE3B8D9B4D09961424B7ABEF47DEE229A499D557",
			ClassicAddress: "rpa9S5fRbS2ZAf2cdFGtezhGYgom1iD4yh",
		},
		{
			Mnemonic: "honey tip lunch empower omit invite nuclear tent brother sadness still exercise odor harbor alcohol huge wait swamp vessel tired swallow supreme silk spawn",
			PublicKey: "038D4EA460687B0FF43E95D9CB56E439AC2C65D890C9FD8B553FF24997C5F91D9A",
			PrivateKey: "00C2C6F996F5FAD5168CAAAD6E58E518641A41E8E164A6FADC540F52D728E2A4AB",
			ClassicAddress: "rPdkbYi6ok7HFpbRo6CSxeDUg951tVdA1m",
		},
		{
			Mnemonic: "hen toe quarter robust elevator badge coconut all place desk pen school topic life seminar run salute paddle hurdle impact push amount oblige citizen",
			PublicKey: "0388AE366DF0D8819760B82319C7A04CA06CC1D49EB5D41ED8DA0C8903DE8FF812",
			PrivateKey: "001CBDDCA87FB57FB5A9D96ECE29D5710DA8421891F40CFA2262120DBB61E5050D",
			ClassicAddress: "rsKbuMTkzR5HU96j8pdGsSuzmZEiZ6mKh5",
		},
	}

	for _, tc := range testCases {
		wallet, err := NewWalletFromMnemonic(tc.Mnemonic)
		if err != nil {
			t.Errorf("Error generating wallet from mnemonic: %s", err)
		}

		if wallet.PublicKey != tc.PublicKey {
			t.Errorf("Public key does not match expected value. Expected: %s, got: %s", tc.PublicKey, wallet.PublicKey)
		}

		if wallet.PrivateKey != tc.PrivateKey {
			t.Errorf("Private key does not match expected value. Expected: %s, got: %s", tc.PrivateKey, wallet.PrivateKey)
		}

		if wallet.ClassicAddress != tc.ClassicAddress {
			t.Errorf("Classic address does not match expected value. Expected: %s, got: %s", tc.ClassicAddress, wallet.ClassicAddress)
		}
	}
}



