package integration

import (
	"encoding/hex"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil/integration"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/results"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/stretchr/testify/require"
)

type CredentialCreateTest struct {
	Name             string
	CredentialCreate *transaction.CredentialCreate
	ExpectedError    string
}

func TestIntegrationCredentialCreateWebsocket(t *testing.T) {
	env := integration.GetWebsocketEnv(t)
	client := websocket.NewClient(websocket.NewClientConfig().WithHost(env.Host).WithFaucetProvider(env.FaucetProvider))

	runner := integration.NewRunner(t, client, &integration.RunnerConfig{
		WalletCount: 1,
	})

	err := runner.Setup()
	require.NoError(t, err)

	sender := runner.GetWallet(0)

	tt := []CredentialCreateTest{
		{
			Name: "pass - valid CredentialCreate",
			CredentialCreate: &transaction.CredentialCreate{
				BaseTx: transaction.BaseTx{
					Account: sender.GetAddress(),
				},
				CredentialType: "6D795F63726564656E7469616C",
				Subject:        sender.GetAddress(),
				Expiration:     10000,
				URI:            hex.EncodeToString([]byte("https://example.com")),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			// check validity
			ok, err := tc.CredentialCreate.Validate()
			require.NoError(t, err)
			require.True(t, ok)

			// flatten and test
			flatTx := tc.CredentialCreate.Flatten()

			_, err = runner.TestTransaction(&flatTx, sender, results.TesSUCCESS.String())

			if tc.ExpectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ExpectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}

	err = runner.Teardown()
	require.NoError(t, err)
}
