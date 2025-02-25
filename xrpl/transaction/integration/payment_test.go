package integration

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil/integration"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/stretchr/testify/require"
)

func TestIntegrationPayment(t *testing.T) {
	env := integration.GetEnv(t)
	runner := integration.NewRunner(t, &integration.RunnerConfig{
		WebsocketConfig: websocket.NewClientConfig().WithHost(env.Host).WithFaucetProvider(env.FaucetProvider),
		WalletCount:     2,
	})

	err := runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	sender := runner.GetWallet(0)
	receiver := runner.GetWallet(1)

	tt := []struct {
		name          string
		payment       *transaction.Payment
		expectedError string
	}{
		{
			name: "pass - XRP to XRP",
			payment: &transaction.Payment{
				BaseTx: transaction.BaseTx{
					Account: sender.GetAddress(),
				},
				Amount:      types.XRPCurrencyAmount(1),
				Destination: receiver.GetAddress(),
			},
		},
		{
			name: "fail - missing Destination",
			payment: &transaction.Payment{
				BaseTx: transaction.BaseTx{
					Account: sender.GetAddress(),
				},
				Amount: types.XRPCurrencyAmount(1),
			},
			expectedError: "invalidTransaction",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			flatTx := tc.payment.Flatten()
			_, err := runner.TestTransaction(&flatTx, sender)
			if tc.expectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
