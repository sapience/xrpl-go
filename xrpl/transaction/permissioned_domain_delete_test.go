package transaction

import (
	"errors"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/stretchr/testify/assert"
)

func TestPermissionedDomainDelete_TxType(t *testing.T) {
	tx := &PermissionedDomainDelete{}
	assert.Equal(t, PermissionedDomainDeleteTx, tx.TxType())
}

func TestPermissionedDomainDelete_Flatten(t *testing.T) {
	tx := &PermissionedDomainDelete{
		BaseTx: BaseTx{
			Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			TransactionType: PermissionedDomainDeleteTx,
		},
		DomainID: "domain123",
	}
	expected := `{
		"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		"TransactionType": "PermissionedDomainDelete",
		"DomainID": "domain123"
	}`
	err := testutil.CompareFlattenAndExpected(tx.Flatten(), []byte(expected))
	if err != nil {
		t.Error(err)
	}
}

func TestPermissionedDomainDelete_Validate(t *testing.T) {
	tests := []struct {
		name        string
		tx          *PermissionedDomainDelete
		wantValid   bool
		wantErr     bool
		expectedErr error
	}{
		{
			name: "pass - valid transaction",
			tx: &PermissionedDomainDelete{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PermissionedDomainDeleteTx,
				},
				DomainID: "domain123",
			},
			wantValid:   true,
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name: "fail - missing DomainID",
			tx: &PermissionedDomainDelete{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PermissionedDomainDeleteTx,
				},
				DomainID: "",
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: errors.New("missing required field: DomainID"),
		},
		{
			name: "fail - invalid base transaction (missing Account)",
			tx: &PermissionedDomainDelete{
				BaseTx: BaseTx{
					Account:         "",
					TransactionType: PermissionedDomainDeleteTx,
				},
				DomainID: "domain123",
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidAccount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.tx.Validate()
			if tt.expectedErr != nil && err != nil {
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			}
			assert.Equal(t, tt.wantValid, valid)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
