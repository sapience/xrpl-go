package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOracleDelete_TxType(t *testing.T) {
	tx := &OracleDelete{}
	assert.Equal(t, OracleDeleteTx, tx.TxType())
}

func TestOracleDelete_Flatten(t *testing.T) {
	testcases := []struct {
		name     string
		tx       *OracleDelete
		expected FlatTransaction
	}{
		{
			name: "pass - empty",
			tx:   &OracleDelete{},
			expected: FlatTransaction{
				"TransactionType":  "OracleDelete",
				"OracleDocumentID": uint32(0),
			},
		},
		{
			name: "pass - complete",
			tx: &OracleDelete{
				BaseTx: BaseTx{
					Account:         "r9cZA1mTh4KVPD5PXPBGVdqw9XRybCz6z",
					TransactionType: OracleDeleteTx,
				},
				OracleDocumentID: 34,
			},
			expected: FlatTransaction{
				"Account":          "r9cZA1mTh4KVPD5PXPBGVdqw9XRybCz6z",
				"TransactionType":  "OracleDelete",
				"OracleDocumentID": uint32(34),
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			actual := testcase.tx.Flatten()
			assert.Equal(t, testcase.expected, actual)
		})
	}
}

func TestOracleDelete_Validate(t *testing.T) {
	testcases := []struct {
		name string
		tx   *OracleDelete
		err  error
	}{
		{
			name: "fail - missing account",
			tx: &OracleDelete{
				BaseTx: BaseTx{
					TransactionType: OracleDeleteTx,
				},
			},
			err: ErrInvalidAccount,
		},
		{
			name: "pass - complete",
			tx: &OracleDelete{
				BaseTx: BaseTx{
					Account:         "r9cZA1mTh4KVPD5PXPBGVdqw9XRybCz6z",
					TransactionType: OracleDeleteTx,
				},
			},
			err: nil,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			ok, err := testcase.tx.Validate()
			assert.Equal(t, testcase.err, err)
			assert.Equal(t, ok, testcase.err == nil)
		})
	}
}
