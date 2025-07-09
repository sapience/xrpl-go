package wallet

import (
	"testing"

	binarycodec "github.com/Peersyst/xrpl-go/binary-codec"
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
			RawTransactions: []types.RawTransaction{
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
			RawTransactions: []types.RawTransaction{
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

func TestCombineBatchSigners(t *testing.T) {
	// Create test wallets using the same seeds as in TypeScript tests
	// rPZsMhM7jNaixFiiipWUuDPifUXCVNYfb6
	edWallet, err := FromSeed("sEdStM1pngFcLQqVfH3RQcg2Qr6ov9e", "")
	require.NoError(t, err)

	// rPMh7Pi9ct699iZUTWaytJUoHcJ7cgyziK
	secpWallet, err := FromSeed("spkcsko6Ag3RbCSVXV2FJ8Pd4Zac1", "")
	require.NoError(t, err)

	// rJCxK2hX9tDMzbnn3cg1GU2g19Kfmhzxkp
	submitWallet, err := FromSeed("sEd7HmQFsoyj5TAm6d98gytM9LJA1MF", "")
	require.NoError(t, err)

	// Helper function to create original batch transaction
	createOriginalBatchTx := func() *transaction.Batch {
		paymentTx1 := &transaction.Payment{
			BaseTx: transaction.BaseTx{
				Account:         types.Address(edWallet.ClassicAddress.String()),
				TransactionType: transaction.PaymentTx,
				Flags:           0x40000000,
				Fee:             types.XRPCurrencyAmount(0),
				Sequence:        215,
				SigningPubKey:   "",
			},
			Destination: types.Address(secpWallet.ClassicAddress.String()),
			Amount:      types.XRPCurrencyAmount(5000000),
		}

		paymentTx2 := &transaction.Payment{
			BaseTx: transaction.BaseTx{
				Account:         types.Address(secpWallet.ClassicAddress.String()),
				TransactionType: transaction.PaymentTx,
				Flags:           0x40000000,
				Fee:             types.XRPCurrencyAmount(0),
				Sequence:        470,
				SigningPubKey:   "",
			},
			Destination: types.Address(submitWallet.ClassicAddress.String()),
			Amount:      types.XRPCurrencyAmount(1000000),
		}

		return &transaction.Batch{
			BaseTx: transaction.BaseTx{
				Account:            types.Address(submitWallet.ClassicAddress.String()),
				TransactionType:    transaction.BatchTx,
				Flags:              1, // tfAllOrNothing
				LastLedgerSequence: 14973,
				NetworkID:          21336,
				Sequence:           215,
			},
			RawTransactions: []types.RawTransaction{
				{
					RawTransaction: paymentTx1.Flatten(),
				},
				{
					RawTransaction: paymentTx2.Flatten(),
				},
			},
		}
	}

	// Helper function to create batch transaction with submitter transaction
	createBatchTxWithSubmitter := func() *transaction.Batch {
		originalTx := createOriginalBatchTx()

		paymentTx3 := &transaction.Payment{
			BaseTx: transaction.BaseTx{
				Account:         types.Address(submitWallet.ClassicAddress.String()), // submitter account
				TransactionType: transaction.PaymentTx,
				Flags:           0x40000000,
				Fee:             types.XRPCurrencyAmount(0),
				Sequence:        470,
				SigningPubKey:   "",
			},
			Destination: types.Address(secpWallet.ClassicAddress.String()),
			Amount:      types.XRPCurrencyAmount(1000000),
		}

		originalTx.RawTransactions = append(originalTx.RawTransactions, types.RawTransaction{
			RawTransaction: paymentTx3.Flatten(),
		})

		return originalTx
	}

	testCases := []struct {
		name          string
		setupTxs      func() []transaction.Batch
		expectedError error
		postCheck     func(t *testing.T, result string, err error)
	}{
		{
			name: "pass - combines valid transactions",
			setupTxs: func() []transaction.Batch {
				tx1 := createOriginalBatchTx()
				tx2 := createOriginalBatchTx()

				err := SignMultiBatch(edWallet, tx1, &SignMultiBatchOptions{})
				require.NoError(t, err)

				err = SignMultiBatch(secpWallet, tx2, &SignMultiBatchOptions{})
				require.NoError(t, err)

				return []transaction.Batch{*tx1, *tx2}
			},
			expectedError: nil,
			postCheck: func(t *testing.T, result string, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, result)

				// Decode the result to verify structure
				decoded, err := binarycodec.Decode(result)
				require.NoError(t, err)
				require.Contains(t, decoded, "BatchSigners")
			},
		},
		{
			name: "pass - sorts the signers",
			setupTxs: func() []transaction.Batch {
				tx1 := createOriginalBatchTx()
				tx2 := createOriginalBatchTx()

				err := SignMultiBatch(edWallet, tx1, &SignMultiBatchOptions{})
				require.NoError(t, err)

				err = SignMultiBatch(secpWallet, tx2, &SignMultiBatchOptions{})
				require.NoError(t, err)

				return []transaction.Batch{*tx2, *tx1} // Note: reversed order to test sorting
			},
			expectedError: nil,
			postCheck: func(t *testing.T, result string, err error) {
				require.NoError(t, err)

				// Decode and verify that signers are sorted by account address
				decoded, err := binarycodec.Decode(result)
				require.NoError(t, err)

				batchSigners, ok := decoded["BatchSigners"].([]interface{})
				require.True(t, ok)
				require.Len(t, batchSigners, 2)

				// Extract the account addresses from the signers
				accounts := make([]string, len(batchSigners))
				for i, signerInterface := range batchSigners {
					signer, ok := signerInterface.(map[string]interface{})
					require.True(t, ok)
					batchSigner, ok := signer["BatchSigner"].(map[string]interface{})
					require.True(t, ok)
					account, ok := batchSigner["Account"].(string)
					require.True(t, ok)
					accounts[i] = account
				}

				// Verify that the accounts are sorted
				require.True(t, accounts[0] < accounts[1], "Accounts should be sorted: %v", accounts)
			},
		},
		{
			name: "pass - removes signer for Batch submitter",
			setupTxs: func() []transaction.Batch {
				originalTx := createBatchTxWithSubmitter()

				tx1 := &transaction.Batch{}
				*tx1 = *originalTx
				tx2 := &transaction.Batch{}
				*tx2 = *originalTx
				tx3 := &transaction.Batch{}
				*tx3 = *originalTx

				err := SignMultiBatch(edWallet, tx1, &SignMultiBatchOptions{})
				require.NoError(t, err)

				err = SignMultiBatch(secpWallet, tx2, &SignMultiBatchOptions{})
				require.NoError(t, err)

				err = SignMultiBatch(submitWallet, tx3, &SignMultiBatchOptions{})
				require.NoError(t, err)

				return []transaction.Batch{*tx1, *tx2, *tx3}
			},
			expectedError: nil,
			postCheck: func(t *testing.T, result string, err error) {
				require.NoError(t, err)

				// Decode and verify that only 2 signers remain (not 3)
				decoded, err := binarycodec.Decode(result)
				require.NoError(t, err)

				batchSigners, ok := decoded["BatchSigners"].([]interface{})
				require.True(t, ok)
				require.Len(t, batchSigners, 2) // Should exclude the submitter's signer
			},
		},
		{
			name: "fail - fails with no transactions provided",
			setupTxs: func() []transaction.Batch {
				return []transaction.Batch{}
			},
			expectedError: ErrNoTransactionsProvided,
			postCheck: func(t *testing.T, result string, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), ErrNoTransactionsProvided.Error())
			},
		},
		{
			name: "fail - fails with no BatchSigners provided in a transaction",
			setupTxs: func() []transaction.Batch {
				tx1 := createOriginalBatchTx()
				tx2 := createOriginalBatchTx()

				// Sign only one transaction
				err := SignMultiBatch(edWallet, tx1, &SignMultiBatchOptions{})
				require.NoError(t, err)

				// tx2 has no BatchSigners
				return []transaction.Batch{*tx1, *tx2}
			},
			expectedError: ErrTxMustIncludeBatchSigner,
			postCheck: func(t *testing.T, result string, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), ErrTxMustIncludeBatchSigner.Error())
			},
		},
		{
			name: "fail - fails with signed inner transaction",
			setupTxs: func() []transaction.Batch {
				tx1 := createOriginalBatchTx()
				tx2 := createOriginalBatchTx()

				err := SignMultiBatch(edWallet, tx1, &SignMultiBatchOptions{})
				require.NoError(t, err)

				err = SignMultiBatch(secpWallet, tx2, &SignMultiBatchOptions{})
				require.NoError(t, err)

				// Sign the transaction completely (add TxnSignature to make it signed)
				tx1.TxnSignature = "some_signature"

				return []transaction.Batch{*tx1, *tx2}
			},
			expectedError: ErrTransactionAlreadySigned,
			postCheck: func(t *testing.T, result string, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), ErrTransactionAlreadySigned.Error())
			},
		},
		{
			name: "fail - fails with different flags signed",
			setupTxs: func() []transaction.Batch {
				tx1 := createOriginalBatchTx()
				tx2 := createOriginalBatchTx()

				// Change flags on tx2
				tx2.Flags = 4 // tfIndependent

				err := SignMultiBatch(edWallet, tx1, &SignMultiBatchOptions{})
				require.NoError(t, err)

				err = SignMultiBatch(secpWallet, tx2, &SignMultiBatchOptions{})
				require.NoError(t, err)

				return []transaction.Batch{*tx1, *tx2}
			},
			expectedError: ErrBatchSignableNotEqual,
			postCheck: func(t *testing.T, result string, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), ErrBatchSignableNotEqual.Error())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			txs := tc.setupTxs()
			result, err := CombineBatchSigners(txs)
			tc.postCheck(t, result, err)
		})
	}
}
