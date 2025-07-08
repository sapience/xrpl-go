package wallet

import (
	"testing"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	wallettypes "github.com/Peersyst/xrpl-go/xrpl/wallet/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignMultiBatch_ED25519(t *testing.T) {
	// Create test wallets using the same seeds as in TypeScript tests
	// rPZsMhM7jNaixFiiipWUuDPifUXCVNYfb6
	edWallet, err := FromSeed("sEdTCFHBquP36KursdZ17ZiuZenJZHg", "")
	require.NoError(t, err)

	// rJCxK2hX9tDMzbnn3cg1GU2g19Kfmhzxkp
	submitWallet, err := FromSeed("sEd7HmQFsoyj5TAm6d98gytM9LJA1MF", "")
	require.NoError(t, err)

	// rwRNeznwHzdfYeKWpevYmax2NSDioyeEtT
	regkeyWallet, err := FromSeed("sEdStM1pngFcLQqVfH3RQcg2Qr6ov9e", "")
	require.NoError(t, err)

	// Create a wallet not included in the batch
	otherWallet, err := New(crypto.ED25519())
	require.NoError(t, err)

	paymentTx1 := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account:         types.Address("rPZsMhM7jNaixFiiipWUuDPifUXCVNYfb6"),
			TransactionType: transaction.PaymentTx,
			Flags:           0x40000000,
			Fee:             types.XRPCurrencyAmount(0),
			Sequence:        215,
			SigningPubKey:   "",
		},
		Destination: types.Address("rPMh7Pi9ct699iZUTWaytJUoHcJ7cgyziK"),
		Amount:      types.XRPCurrencyAmount(5000000),
	}

	paymentTx2 := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account:         types.Address("rPMh7Pi9ct699iZUTWaytJUoHcJ7cgyziK"),
			TransactionType: transaction.PaymentTx,
			Flags:           0x40000000,
			Fee:             types.XRPCurrencyAmount(0),
			Sequence:        470,
			SigningPubKey:   "",
		},
		Destination: types.Address("rJCxK2hX9tDMzbnn3cg1GU2g19Kfmhzxkp"),
		Amount:      types.XRPCurrencyAmount(1000000),
	}

	flatPaymentTx1 := paymentTx1.Flatten()
	flatPaymentTx2 := paymentTx2.Flatten()

	// Create test batch transaction
	createBatchTx := func() *transaction.Batch {
		return &transaction.Batch{
			BaseTx: transaction.BaseTx{
				Account:         types.Address("rJCxK2hX9tDMzbnn3cg1GU2g19Kfmhzxkp"),
				TransactionType: transaction.BatchTx,
			},
			RawTransactions: []transaction.RawTransactionWrapper{
				{
					RawTransaction: flatPaymentTx1,
				},
				{
					RawTransaction: flatPaymentTx2,
				},
			},
		}
	}

	tc := []struct {
		name          string
		wallet        Wallet
		tx            *transaction.Batch
		opts          SignMultiBatchOptions
		postCheck     func(t *testing.T, tx *transaction.Batch)
		expectedError error
	}{
		{
			name:   "pass - succeeds with ed25519 seed",
			wallet: edWallet,
			tx:     createBatchTx(),
			opts:   SignMultiBatchOptions{},
			postCheck: func(t *testing.T, tx *transaction.Batch) {
				require.NotNil(t, tx.BatchSigners)
				require.Len(t, tx.BatchSigners, 1)

				batchSigner := tx.BatchSigners[0]
				require.Equal(t, types.Address("rPZsMhM7jNaixFiiipWUuDPifUXCVNYfb6"), batchSigner.BatchSigner.Account)
				require.NotEmpty(t, batchSigner.BatchSigner.SigningPubKey)
				require.NotEmpty(t, batchSigner.BatchSigner.TxnSignature)
			},
			expectedError: nil,
		},
		{
			name:   "pass - succeeds with a different account",
			wallet: regkeyWallet,
			tx:     createBatchTx(),
			opts: SignMultiBatchOptions{
				BatchAccount: wallettypes.NewBatchAccount("rPZsMhM7jNaixFiiipWUuDPifUXCVNYfb6"),
			},
			postCheck: func(t *testing.T, tx *transaction.Batch) {
				require.NotNil(t, tx.BatchSigners)
				require.Len(t, tx.BatchSigners, 1)

				batchSigner := tx.BatchSigners[0]
				require.Equal(t, types.Address("rPZsMhM7jNaixFiiipWUuDPifUXCVNYfb6"), batchSigner.BatchSigner.Account)
				require.NotEmpty(t, batchSigner.BatchSigner.SigningPubKey)
				require.NotEmpty(t, batchSigner.BatchSigner.TxnSignature)
			},
			expectedError: nil,
		},
		{
			name:   "pass - succeeds with multisign",
			wallet: regkeyWallet,
			tx:     createBatchTx(),
			opts: SignMultiBatchOptions{
				BatchAccount: wallettypes.NewBatchAccount(edWallet.ClassicAddress.String()),
				Multisign:    true,
			},
			postCheck: func(t *testing.T, tx *transaction.Batch) {
				require.NotNil(t, tx.BatchSigners)
				require.Len(t, tx.BatchSigners, 1)

				batchSigner := tx.BatchSigners[0]
				assert.Equal(t, types.Address("rPZsMhM7jNaixFiiipWUuDPifUXCVNYfb6"), batchSigner.BatchSigner.Account)
				require.NotNil(t, batchSigner.BatchSigner.Signers)
				require.Len(t, batchSigner.BatchSigner.Signers, 1)

				signer := batchSigner.BatchSigner.Signers[0]
				require.Equal(t, types.Address("rwRNeznwHzdfYeKWpevYmax2NSDioyeEtT"), signer.SignerData.Account)
				require.NotEmpty(t, signer.SignerData.SigningPubKey)
				require.NotEmpty(t, signer.SignerData.TxnSignature)
			},
			expectedError: nil,
		},
		{
			name:   "pass - succeeds with multisign + regular key",
			wallet: regkeyWallet,
			tx:     createBatchTx(),
			opts: SignMultiBatchOptions{
				BatchAccount:     wallettypes.NewBatchAccount(edWallet.ClassicAddress.String()),
				MultisignAccount: submitWallet.ClassicAddress.String(),
			},
			postCheck: func(t *testing.T, tx *transaction.Batch) {
				require.NotNil(t, tx.BatchSigners)
				require.Len(t, tx.BatchSigners, 1)

				batchSigner := tx.BatchSigners[0]
				require.Equal(t, types.Address("rPZsMhM7jNaixFiiipWUuDPifUXCVNYfb6"), batchSigner.BatchSigner.Account)
				require.NotNil(t, batchSigner.BatchSigner.Signers)
				require.Len(t, batchSigner.BatchSigner.Signers, 1)

				signer := batchSigner.BatchSigner.Signers[0]
				require.Equal(t, types.Address("rJCxK2hX9tDMzbnn3cg1GU2g19Kfmhzxkp"), signer.SignerData.Account)
				require.NotEmpty(t, signer.SignerData.SigningPubKey)
				require.NotEmpty(t, signer.SignerData.TxnSignature)
			},
			expectedError: nil,
		},
		{
			name:          "fail - fails with not-included account",
			wallet:        otherWallet,
			tx:            createBatchTx(),
			opts:          SignMultiBatchOptions{},
			postCheck:     func(t *testing.T, tx *transaction.Batch) {},
			expectedError: ErrBatchAccountNotFound,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			err := SignMultiBatch(tt.wallet, tt.tx, &tt.opts)
			if tt.expectedError == nil {
				require.NoError(t, err)
				tt.postCheck(t, tt.tx)
			} else {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
			}
		})
	}
}

func TestSignMultiBatch_SECP256K1(t *testing.T) {
	// Create test wallets using the same seeds as in TypeScript tests
	// rPMh7Pi9ct699iZUTWaytJUoHcJ7cgyziK
	secpWallet, err := FromSeed("spkcsko6Ag3RbCSVXV2FJ8Pd4Zac1", "")
	require.NoError(t, err)

	// Create a wallet not included in the batch
	otherWallet, err := New(crypto.SECP256K1())
	require.NoError(t, err)

	paymentTx1 := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account:         types.Address("rJy554HmWFFJQGnRfZuoo8nV97XSMq77h7"),
			TransactionType: transaction.PaymentTx,
			Flags:           0x40000000,
			Fee:             types.XRPCurrencyAmount(0),
			Sequence:        215,
			SigningPubKey:   "",
		},
		Destination: types.Address("rPMh7Pi9ct699iZUTWaytJUoHcJ7cgyziK"),
		Amount:      types.XRPCurrencyAmount(5000000),
	}

	paymentTx2 := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account:         types.Address("rPMh7Pi9ct699iZUTWaytJUoHcJ7cgyziK"),
			TransactionType: transaction.PaymentTx,
			Flags:           0x40000000,
			Fee:             types.XRPCurrencyAmount(0),
			Sequence:        470,
			SigningPubKey:   "",
		},
		Destination: types.Address("rJCxK2hX9tDMzbnn3cg1GU2g19Kfmhzxkp"),
		Amount:      types.XRPCurrencyAmount(1000000),
	}

	flatPaymentTx1 := paymentTx1.Flatten()
	flatPaymentTx2 := paymentTx2.Flatten()
	// Create test batch transaction
	createBatchTx := func() *transaction.Batch {
		return &transaction.Batch{
			BaseTx: transaction.BaseTx{
				Account:         types.Address("rJCxK2hX9tDMzbnn3cg1GU2g19Kfmhzxkp"),
				TransactionType: transaction.BatchTx,
			},
			RawTransactions: []transaction.RawTransactionWrapper{
				{
					RawTransaction: flatPaymentTx1,
				},
				{
					RawTransaction: flatPaymentTx2,
				},
			},
		}
	}

	tc := []struct {
		name          string
		wallet        Wallet
		tx            *transaction.Batch
		opts          SignMultiBatchOptions
		postCheck     func(t *testing.T, tx *transaction.Batch)
		expectedError error
	}{
		{
			name:   "pass - succeeds with secp256k1 seed",
			wallet: secpWallet,
			tx:     createBatchTx(),
			opts:   SignMultiBatchOptions{},
			postCheck: func(t *testing.T, tx *transaction.Batch) {
				require.NotNil(t, tx.BatchSigners)
				require.Len(t, tx.BatchSigners, 1)

				batchSigner := tx.BatchSigners[0]
				require.Equal(t, types.Address("rPMh7Pi9ct699iZUTWaytJUoHcJ7cgyziK"), batchSigner.BatchSigner.Account)
				require.NotEmpty(t, batchSigner.BatchSigner.SigningPubKey)
				require.NotEmpty(t, batchSigner.BatchSigner.TxnSignature)
			},
			expectedError: nil,
		},
		{
			name:          "fail - fails with not-included account",
			wallet:        otherWallet,
			tx:            createBatchTx(),
			opts:          SignMultiBatchOptions{},
			postCheck:     func(t *testing.T, tx *transaction.Batch) {},
			expectedError: ErrBatchAccountNotFound,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			err := SignMultiBatch(tt.wallet, tt.tx, &tt.opts)
			if tt.expectedError == nil {
				require.NoError(t, err)
				tt.postCheck(t, tt.tx)
			} else {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
			}
		})
	}
}
