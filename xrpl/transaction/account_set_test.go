package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountSetTfFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*AccountSet)
		expected uint32
	}{
		{
			name: "pass - SetRequireDestTag",
			setter: func(s *AccountSet) {
				s.SetRequireDestTag()
			},
			expected: tfRequireDestTag,
		},
		{
			name: "pass - SetRequireAuth",
			setter: func(s *AccountSet) {
				s.SetRequireAuth()
			},
			expected: tfRequireAuth,
		},
		{
			name: "pass - SetDisallowXRP",
			setter: func(s *AccountSet) {
				s.SetDisallowXRP()
			},
			expected: tfDisallowXRP,
		},
		{
			name: "pass - SetOptionalDestTag",
			setter: func(s *AccountSet) {
				s.SetOptionalDestTag()
			},
			expected: tfOptionalDestTag,
		},
		{
			name: "pass - SetRequireDestTag and SetRequireAuth",
			setter: func(s *AccountSet) {
				s.SetRequireDestTag()
				s.SetRequireAuth()
			},
			expected: tfRequireDestTag | tfRequireAuth,
		},
		{
			name: "pass - SetDisallowXRP and SetOptionalDestTag",
			setter: func(s *AccountSet) {
				s.SetDisallowXRP()
				s.SetOptionalDestTag()
			},
			expected: tfDisallowXRP | tfOptionalDestTag,
		},
		{
			name: "pass - SetRequireDestTag, SetRequireAuth, and SetDisallowXRP",
			setter: func(s *AccountSet) {
				s.SetRequireDestTag()
				s.SetRequireAuth()
				s.SetDisallowXRP()
			},
			expected: tfRequireDestTag | tfRequireAuth | tfDisallowXRP,
		},
		{
			name: "pass - All flags",
			setter: func(s *AccountSet) {
				s.SetRequireDestTag()
				s.SetRequireAuth()
				s.SetDisallowXRP()
				s.SetOptionalDestTag()
				s.SetOptionalAuth()
				s.SetAllowXRP()
			},
			expected: tfRequireDestTag | tfRequireAuth | tfDisallowXRP | tfOptionalDestTag | tfOptionalAuth | tfAllowXRP,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AccountSet{}
			tt.setter(s)
			if s.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, s.Flags)
			}
		})
	}
}

func TestAccountSetAsfFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*AccountSet)
		expected uint32
	}{
		{
			name: "pass - SetAsfRequireDest",
			setter: func(s *AccountSet) {
				s.SetAsfRequireDest()
			},
			expected: asfRequireDest,
		},
		{
			name: "pass - SetAsfRequireAuth",
			setter: func(s *AccountSet) {
				s.SetAsfRequireAuth()
			},
			expected: asfRequireAuth,
		},
		{
			name: "pass - SetAsfDisallowXRP",
			setter: func(s *AccountSet) {
				s.SetAsfDisallowXRP()
			},
			expected: asfDisallowXRP,
		},
		{
			name: "pass - SetAsfDisableMaster",
			setter: func(s *AccountSet) {
				s.SetAsfDisableMaster()
			},
			expected: asfDisableMaster,
		},
		{
			name: "pass - SetAsfAccountTxnID",
			setter: func(s *AccountSet) {
				s.SetAsfAccountTxnID()
			},
			expected: asfAccountTxnID,
		},
		{
			name: "pass - SetAsfNoFreeze",
			setter: func(s *AccountSet) {
				s.SetAsfNoFreeze()
			},
			expected: asfNoFreeze,
		},
		{
			name: "pass - SetAsfGlobalFreeze",
			setter: func(s *AccountSet) {
				s.SetAsfGlobalFreeze()
			},
			expected: asfGlobalFreeze,
		},
		{
			name: "pass - SetAsfDefaultRipple",
			setter: func(s *AccountSet) {
				s.SetAsfDefaultRipple()
			},
			expected: asfDefaultRipple,
		},
		{
			name: "pass - SetAsfDepositAuth",
			setter: func(s *AccountSet) {
				s.SetAsfDepositAuth()
			},
			expected: asfDepositAuth,
		},
		{
			name: "pass - SetAsfAuthorizedNFTokenMinter",
			setter: func(s *AccountSet) {
				s.SetAsfAuthorizedNFTokenMinter()
			},
			expected: asfAuthorizedNFTokenMinter,
		},
		{
			name: "pass - SetAsfDisallowIncomingNFTokenOffer",
			setter: func(s *AccountSet) {
				s.SetAsfDisallowIncomingNFTokenOffer()
			},
			expected: asfDisallowIncomingNFTokenOffer,
		},
		{
			name: "pass - SetAsfDisallowIncomingCheck",
			setter: func(s *AccountSet) {
				s.SetAsfDisallowIncomingCheck()
			},
			expected: asfDisallowIncomingCheck,
		},
		{
			name: "pass - SetAsfDisallowIncomingPayChan",
			setter: func(s *AccountSet) {
				s.SetAsfDisallowIncomingPayChan()
			},
			expected: asfDisallowIncomingPayChan,
		},
		{
			name: "pass - SetAsfDisallowIncomingTrustLine",
			setter: func(s *AccountSet) {
				s.SetAsfDisallowIncomingTrustLine()
			},
			expected: asfDisallowIncomingTrustLine,
		},
		{
			name: "pass - SetAsfAllowTrustLineClawback",
			setter: func(s *AccountSet) {
				s.SetAsfAllowTrustLineClawback()
			},
			expected: asfAllowTrustLineClawback,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AccountSet{}
			tt.setter(s)
			if s.SetFlag != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, s.Flags)
			}
		})
	}
}

func TestAccountClearAsfFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*AccountSet)
		expected uint32
	}{
		{
			name: "pass - ClearAsfRequireDest",
			setter: func(s *AccountSet) {
				s.ClearAsfRequireDest()
			},
			expected: asfRequireDest,
		},
		{
			name: "pass - ClearAsfRequireAuth",
			setter: func(s *AccountSet) {
				s.ClearAsfRequireAuth()
			},
			expected: asfRequireAuth,
		},
		{
			name: "pass - ClearAsfDisallowXRP",
			setter: func(s *AccountSet) {
				s.ClearAsfDisallowXRP()
			},
			expected: asfDisallowXRP,
		},
		{
			name: "pass - ClearAsfDisableMaster",
			setter: func(s *AccountSet) {
				s.ClearAsfDisableMaster()
			},
			expected: asfDisableMaster,
		},
		{
			name: "pass - ClearAsfAccountTxnID",
			setter: func(s *AccountSet) {
				s.ClearAsfAccountTxnID()
			},
			expected: asfAccountTxnID,
		},
		{
			name: "pass - asfNoFreeze",
			setter: func(s *AccountSet) {
				s.ClearAsfNoFreeze()
			},
			expected: asfNoFreeze,
		},
		{
			name: "pass - asfGlobalFreeze",
			setter: func(s *AccountSet) {
				s.ClearAsfGlobalFreeze()
			},
			expected: asfGlobalFreeze,
		},
		{
			name: "pass - ClearAsfDefaultRipple",
			setter: func(s *AccountSet) {
				s.ClearAsfDefaultRipple()
			},
			expected: asfDefaultRipple,
		},
		{
			name: "pass - ClearAsfDepositAuth",
			setter: func(s *AccountSet) {
				s.ClearAsfDepositAuth()
			},
			expected: asfDepositAuth,
		},
		{
			name: "pass - ClearAsfAuthorizedNFTokenMinter",
			setter: func(s *AccountSet) {
				s.ClearAsfAuthorizedNFTokenMinter()
			},
			expected: asfAuthorizedNFTokenMinter,
		},
		{
			name: "pass - ClearAsfDisallowIncomingNFTokenOffer",
			setter: func(s *AccountSet) {
				s.ClearAsfDisallowIncomingNFTokenOffer()
			},
			expected: asfDisallowIncomingNFTokenOffer,
		},
		{
			name: "pass - ClearAsfDisallowIncomingCheck",
			setter: func(s *AccountSet) {
				s.ClearAsfDisallowIncomingCheck()
			},
			expected: asfDisallowIncomingCheck,
		},
		{
			name: "pass - ClearAsfDisallowIncomingPayChan",
			setter: func(s *AccountSet) {
				s.ClearAsfDisallowIncomingPayChan()
			},
			expected: asfDisallowIncomingPayChan,
		},
		{
			name: "pass - ClearAsfDisallowIncomingTrustLine",
			setter: func(s *AccountSet) {
				s.ClearAsfDisallowIncomingTrustLine()
			},
			expected: asfDisallowIncomingTrustLine,
		},
		{
			name: "pass - ClearAsfAllowTrustLineClawback",
			setter: func(s *AccountSet) {
				s.ClearAsfAllowTrustLineClawback()
			},
			expected: asfAllowTrustLineClawback,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AccountSet{}
			tt.setter(s)
			if s.ClearFlag != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, s.Flags)
			}
		})
	}
}

func TestAccountSet_Validate(t *testing.T) {
	testCases := []struct {
		name       string
		accountSet *AccountSet
		valid      bool
	}{
		{
			name: "pass - Valid AccountSet",
			accountSet: &AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				ClearFlag:    1,
				SetFlag:      2,
				Domain:       types.Domain("A5B21758D2318FA2C"),
				EmailHash:    types.EmailHash("1234567890abcdef"),
				MessageKey:   types.MessageKey("messageKey"),
				TransferRate: TransferRate(1000000001),
				TickSize:     TickSize(5),
			},
			valid: true,
		},
		{
			name: "pass - Valid AccountSet without options, just the commons fields",
			accountSet: &AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
			},
			valid: true,
		},
		{
			name: "fail - Invalid AccountSet with high SetFlag",
			accountSet: &AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				SetFlag: 18, // too high
			},
			valid: false,
		},
		{
			name: "fail - Invalid AccountSet with low TickSize",
			accountSet: &AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				TickSize: TickSize(2),
			},
			valid: false,
		},
		{
			name: "fail - Invalid AccountSet with high TickSize",
			accountSet: &AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				TickSize: TickSize(16),
			},
			valid: false,
		},
		{
			name: "pass - Valid AccountSet TickSize set to 0 to disable it",
			accountSet: &AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				TickSize: TickSize(0),
			},
			valid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			valid, err := tc.accountSet.Validate()
			if valid != tc.valid {
				t.Errorf("Validation result for %s is incorrect. Expected: %v, Got: %v", tc.name, tc.valid, valid)
			}
			if err != nil && tc.valid {
				t.Errorf("Validation failed for %s: %s", tc.name, err)
			}
			if err == nil && !tc.valid {
				t.Errorf("Validation should have failed for %s", tc.name)
			}
		})
	}
}

func TestAccountSet_Flatten(t *testing.T) {
	tests := []struct {
		name       string
		accountSet *AccountSet
		expected   FlatTransaction
	}{
		{
			name: "pass - Flatten with all fields",
			accountSet: &AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				ClearFlag:     asfRequireDest,
				Domain:        types.Domain("A5B21758D2318FA2C"),
				EmailHash:     types.EmailHash("1234567890abcdef"),
				MessageKey:    types.MessageKey("messagekey"),
				NFTokenMinter: types.NFTokenMinter("nftokenminter"),
				SetFlag:       asfRequireAuth,
				TransferRate:  TransferRate(1000000001),
				TickSize:      TickSize(5),
				WalletLocator: WalletLocator("walletLocator"),
				WalletSize:    WalletSize(10),
			},
			expected: FlatTransaction{
				"Account":         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
				"TransactionType": "AccountSet",
				"Fee":             "1",
				"Sequence":        uint32(1234),
				"SigningPubKey":   "ghijk",
				"TxnSignature":    "A1B2C3D4E5F6",
				"ClearFlag":       asfRequireDest,
				"Domain":          "A5B21758D2318FA2C",
				"EmailHash":       "1234567890abcdef",
				"MessageKey":      "messagekey",
				"NFTokenMinter":   "nftokenminter",
				"SetFlag":         asfRequireAuth,
				"TransferRate":    uint32(1000000001),
				"TickSize":        uint8(5),
				"WalletLocator":   "walletLocator",
				"WalletSize":      uint32(10),
			},
		},
		{
			name: "pass - Flatten with empty string or value set to 0 to remove/disable the fields",
			accountSet: &AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
				Domain:        types.Domain(""),
				EmailHash:     types.EmailHash(""),
				TickSize:      TickSize(0),
				TransferRate:  TransferRate(0),
				NFTokenMinter: types.NFTokenMinter(""),
				WalletLocator: WalletLocator(""),
				WalletSize:    WalletSize(0),
			},
			expected: FlatTransaction{
				"Account":         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
				"TransactionType": "AccountSet",
				"Fee":             "1",
				"Sequence":        uint32(1234),
				"SigningPubKey":   "ghijk",
				"TxnSignature":    "A1B2C3D4E5F6",
				"Domain":          "",
				"EmailHash":       "",
				"TickSize":        uint8(0),
				"TransferRate":    uint32(0),
				"NFTokenMinter":   "",
				"WalletLocator":   "",
				"WalletSize":      uint32(0),
			},
		},
		{
			name: "pass - Flatten with required strings only",
			accountSet: &AccountSet{
				BaseTx: BaseTx{
					Account:         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
					TransactionType: AccountSetTx,
					Fee:             types.XRPCurrencyAmount(1),
					Sequence:        1234,
					SigningPubKey:   "ghijk",
					TxnSignature:    "A1B2C3D4E5F6",
				},
			},
			expected: FlatTransaction{
				"Account":         "r7dawf5hSG71faLnCrPiAQ5DkXfVxULPs",
				"TransactionType": "AccountSet",
				"Fee":             "1",
				"Sequence":        uint32(1234),
				"SigningPubKey":   "ghijk",
				"TxnSignature":    "A1B2C3D4E5F6",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			flattened := tc.accountSet.Flatten()
			require.Equal(t, tc.expected, flattened)
		})
	}
}

func TestAccountSet_TxType(t *testing.T) {
	entry := &AccountSet{}
	assert.Equal(t, AccountSetTx, entry.TxType())
}
