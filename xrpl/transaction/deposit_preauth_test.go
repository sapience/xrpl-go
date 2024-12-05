package transaction

import (
	"testing"

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
			name: "fail - must set either authorize or unauthorize",
			tx: &DepositPreauth{
				BaseTx: BaseTx{
					Account:         "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
					TransactionType: DepositPreauthTx,
				},
			},
			expected:    false,
			expectedErr: ErrDepositPreauthMustSetEitherAuthorizeOrUnauthorize,
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
			ok, err := testcase.tx.Validate()
			if testcase.expectedErr != nil {
				require.Error(t, err)
				require.Equal(t, testcase.expectedErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, testcase.expected, ok)
			}
		})
	}
}
