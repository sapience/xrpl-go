package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/require"
)

func TestDepositPreauth_TxType(t *testing.T) {
	tx := &DepositPreauth{}
	require.Equal(t, DepositPreauthTx, tx.TxType())
}

func TestDepositPreauth_Flatten(t *testing.T) {
	testcases := []struct {
		name     string
		tx       *DepositPreauth
		expected FlatTransaction
	}{
		{
			name: "pass - base transaction",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				},
			},
			expected: FlatTransaction{
				"Account":         "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"TransactionType": "DepositPreauth",
			},
		},
		{
			name: "pass - authorize",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				},
				Authorize: "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
			},
			expected: FlatTransaction{
				"Account":         "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"TransactionType": "DepositPreauth",
				"Authorize":       "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
			},
		},
		{
			name: "pass - unauthorize",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				},
				Unauthorize: "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
			},
			expected: FlatTransaction{
				"Account":         "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"TransactionType": "DepositPreauth",
				"Unauthorize":     "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
			},
		},
		{
			name: "pass - authorize credentials",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				},
				AuthorizeCredentials: []types.AuthorizeCredentials{
					{
						Issuer:         "rIssuer",
						CredentialType: "6D795F63726564656E7469616C",
					},
				},
			},
			expected: FlatTransaction{
				"Account":         "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"TransactionType": "DepositPreauth",
				"AuthorizeCredentials": []interface{}{
					map[string]interface{}{"Issuer": "rIssuer", "CredentialType": "6D795F63726564656E7469616C"},
				},
			},
		},
		{
			name: "pass - unauthorize credentials",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				},
				UnauthorizeCredentials: []types.AuthorizeCredentials{
					{
						Issuer:         "rIssuer",
						CredentialType: "6D795F63726564656E7469616C",
					},
				},
			},
			expected: FlatTransaction{
				"Account":         "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"TransactionType": "DepositPreauth",
				"UnauthorizeCredentials": []interface{}{
					map[string]interface{}{"Issuer": "rIssuer", "CredentialType": "6D795F63726564656E7469616C"},
				},
			},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			require.Equal(t, testcase.expected, testcase.tx.Flatten())
		})
	}
}

func TestDepositPreauth_Validate(t *testing.T) {
	testcases := []struct {
		name        string
		tx          *DepositPreauth
		expected    bool
		expectedErr error
	}{
		{
			name: "fail - base transaction missing account",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					TransactionType: DepositPreauthTx,
					Account:         "",
				},
			},
			expected:    false,
			expectedErr: ErrInvalidAccount,
		},
		{
			name: "fail - invalid authorize",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					TransactionType: DepositPreauthTx,
					Account:         "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				},
				Authorize: "invalid",
			},
			expected:    false,
			expectedErr: ErrDepositPreauthInvalidAuthorize,
		},
		{
			name: "fail - invalid unauthorize",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					TransactionType: DepositPreauthTx,
					Account:         "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				},
				Unauthorize: "invalid",
			},
			expected:    false,
			expectedErr: ErrDepositPreauthInvalidUnauthorize,
		},
		{
			name: "fail - must set only one field",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					Account:         "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
					TransactionType: DepositPreauthTx,
				},
			},
			expected:    false,
			expectedErr: ErrDepositPreauthMustSetOnlyOneField,
		},
		{
			name: "pass - authorize",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					Account:         "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
					TransactionType: DepositPreauthTx,
				},
				Authorize: "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
			},
			expected: true,
		},
		{
			name: "pass - unauthorize",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					Account:         "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
					TransactionType: DepositPreauthTx,
				},
				Unauthorize: "rEhxGqkqPPSxQ3P25J66ft5TwpzV14k2de",
			},
			expected: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			valid, err := testcase.tx.Validate()
			require.Equal(t, testcase.expected, valid)
			if testcase.expectedErr != nil {
				require.Error(t, err)
				require.Equal(t, testcase.expectedErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDepositPreauth_IsOnlyOneFieldSet(t *testing.T) {
	testcases := []struct {
		name     string
		dp       *DepositPreauth
		expected bool
	}{
		{
			name: "pass - only Authorize set",
			dp: &DepositPreauth{
				Authorize: "rAuthorize",
			},
			expected: true,
		},
		{
			name: "pass - only AuthorizeCredentials set",
			dp: &DepositPreauth{
				AuthorizeCredentials: []types.AuthorizeCredentials{
					{
						Issuer:         "rIssuer",
						CredentialType: "6D795F63726564656E7469616C",
					},
				},
			},
			expected: true,
		},
		{
			name: "pass - only Unauthorize set",
			dp: &DepositPreauth{
				Unauthorize: "rUnauthorize",
			},
			expected: true,
		},
		{
			name: "pass - only UnauthorizeCredentials set",
			dp: &DepositPreauth{
				UnauthorizeCredentials: []types.AuthorizeCredentials{
					{
						Issuer:         "rIssuer",
						CredentialType: "6D795F63726564656E7469616C",
					},
				},
			},
			expected: true,
		},
		{
			name:     "fail - no fields set",
			dp:       &DepositPreauth{},
			expected: false,
		},
		{
			name: "fail - Authorize and AuthorizeCredentials set",
			dp: &DepositPreauth{
				Authorize: "rAuthorize",
				AuthorizeCredentials: []types.AuthorizeCredentials{
					{
						Issuer:         "rIssuer",
						CredentialType: "6D795F63726564656E7469616C",
					},
				},
			},
			expected: false,
		},
		{
			name: "fail - Authorize and Unauthorize set",
			dp: &DepositPreauth{
				Authorize:   "rAuthorize",
				Unauthorize: "rUnauthorize",
			},
			expected: false,
		},
		{
			name: "fail - AuthorizeCredentials and UnauthorizeCredentials set",
			dp: &DepositPreauth{
				AuthorizeCredentials: []types.AuthorizeCredentials{
					{
						Issuer:         "rIssuer",
						CredentialType: "6D795F63726564656E7469616C",
					},
				},
				UnauthorizeCredentials: []types.AuthorizeCredentials{
					{
						Issuer:         "rIssuer",
						CredentialType: "6D795F63726564656E7469616C",
					},
				},
			},
			expected: false,
		},
		{
			name: "fail - all fields set",
			dp: &DepositPreauth{
				Authorize: "rAuthorize",
				AuthorizeCredentials: []types.AuthorizeCredentials{
					{
						Issuer:         "rIssuer",
						CredentialType: "6D795F63726564656E7469616C",
					},
				},
				Unauthorize: "rUnauthorize",
				UnauthorizeCredentials: []types.AuthorizeCredentials{
					{
						Issuer:         "rIssuer",
						CredentialType: "6D795F63726564656E7469616C",
					},
				},
			},
			expected: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			result := testcase.dp.IsOnlyOneFieldSet()
			require.Equal(t, testcase.expected, result)
		})
	}
}
