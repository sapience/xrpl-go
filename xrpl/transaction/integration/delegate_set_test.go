package integration

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/xrpl/testutil/integration"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/stretchr/testify/require"
)

type DelegateSetTest struct {
	Name          string
	DelegateSet   *transaction.DelegateSet
	ExpectedError string
}

func TestIntegrationDelegateSet_Websocket(t *testing.T) {
	env := integration.GetWebsocketEnv(t)
	client := websocket.NewClient(websocket.NewClientConfig().WithHost(env.Host).WithFaucetProvider(env.FaucetProvider))

	runner := integration.NewRunner(t, client, &integration.RunnerConfig{
		WalletCount: 2,
	})

	err := runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	delegator := runner.GetWallet(0)
	delegatee := runner.GetWallet(1)

	tt := []DelegateSetTest{
		{
			Name: "pass - single transaction type permission",
			DelegateSet: &transaction.DelegateSet{
				BaseTx: transaction.BaseTx{
					Account: delegator.GetAddress(),
				},
				Authorize: delegatee.GetAddress(),
				Permissions: []types.Permission{
					{
						Permission: types.PermissionValue{
							PermissionValue: "Payment",
						},
					},
				},
			},
		},
		{
			Name: "pass - multiple granular permissions",
			DelegateSet: &transaction.DelegateSet{
				BaseTx: transaction.BaseTx{
					Account: delegator.GetAddress(),
				},
				Authorize: delegatee.GetAddress(),
				Permissions: []types.Permission{
					{
						Permission: types.PermissionValue{
							PermissionValue: "TrustlineAuthorize",
						},
					},
					{
						Permission: types.PermissionValue{
							PermissionValue: "AccountDomainSet",
						},
					},
					{
						Permission: types.PermissionValue{
							PermissionValue: "PaymentMint",
						},
					},
				},
			},
		},
		{
			Name: "pass - mixed transaction type and granular permissions",
			DelegateSet: &transaction.DelegateSet{
				BaseTx: transaction.BaseTx{
					Account: delegator.GetAddress(),
				},
				Authorize: delegatee.GetAddress(),
				Permissions: []types.Permission{
					{
						Permission: types.PermissionValue{
							PermissionValue: "Payment",
						},
					},
					{
						Permission: types.PermissionValue{
							PermissionValue: "TrustSet",
						},
					},
					{
						Permission: types.PermissionValue{
							PermissionValue: "TrustlineAuthorize",
						},
					},
				},
			},
		},
		{
			Name: "fail - missing Authorize field",
			DelegateSet: &transaction.DelegateSet{
				BaseTx: transaction.BaseTx{
					Account: delegator.GetAddress(),
				},
				Permissions: []types.Permission{
					{
						Permission: types.PermissionValue{
							PermissionValue: "Payment",
						},
					},
				},
			},
			ExpectedError: ErrInvalidTransaction,
		},
		{
			Name: "fail - empty Permissions array",
			DelegateSet: &transaction.DelegateSet{
				BaseTx: transaction.BaseTx{
					Account: delegator.GetAddress(),
				},
				Authorize:   delegatee.GetAddress(),
				Permissions: []types.Permission{},
			},
			ExpectedError: ErrInvalidTransaction,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			flatTx := tc.DelegateSet.Flatten()
			_, err := runner.TestTransaction(&flatTx, delegator, "tesSUCCESS")
			if tc.ExpectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ExpectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIntegrationDelegateSet_RPCClient(t *testing.T) {
	env := integration.GetRPCEnv(t)
	clientCfg, err := rpc.NewClientConfig(env.Host, rpc.WithFaucetProvider(env.FaucetProvider))
	require.NoError(t, err)
	client := rpc.NewClient(clientCfg)

	runner := integration.NewRunner(t, client, integration.NewRunnerConfig(
		integration.WithWallets(2),
	))

	err = runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	delegator := runner.GetWallet(0)
	delegatee := runner.GetWallet(1)

	tt := []DelegateSetTest{
		{
			Name: "pass - single transaction type permission",
			DelegateSet: &transaction.DelegateSet{
				BaseTx: transaction.BaseTx{
					Account: delegator.GetAddress(),
				},
				Authorize: delegatee.GetAddress(),
				Permissions: []types.Permission{
					{
						Permission: types.PermissionValue{
							PermissionValue: "Payment",
						},
					},
				},
			},
		},
		{
			Name: "pass - mixed transaction type and granular permissions",
			DelegateSet: &transaction.DelegateSet{
				BaseTx: transaction.BaseTx{
					Account: delegator.GetAddress(),
				},
				Authorize: delegatee.GetAddress(),
				Permissions: []types.Permission{
					{
						Permission: types.PermissionValue{
							PermissionValue: "Payment",
						},
					},
					{
						Permission: types.PermissionValue{
							PermissionValue: "TrustlineAuthorize",
						},
					},
				},
			},
		},
		{
			Name: "fail - missing Authorize field",
			DelegateSet: &transaction.DelegateSet{
				BaseTx: transaction.BaseTx{
					Account: delegator.GetAddress(),
				},
				Permissions: []types.Permission{
					{
						Permission: types.PermissionValue{
							PermissionValue: "Payment",
						},
					},
				},
			},
			ExpectedError: ErrInvalidTransaction,
		},
		{
			Name: "fail - empty Permissions array",
			DelegateSet: &transaction.DelegateSet{
				BaseTx: transaction.BaseTx{
					Account: delegator.GetAddress(),
				},
				Authorize:   delegatee.GetAddress(),
				Permissions: []types.Permission{},
			},
			ExpectedError: ErrInvalidTransaction,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			flatTx := tc.DelegateSet.Flatten()
			_, err := runner.TestTransaction(&flatTx, delegator, "tesSUCCESS")
			if tc.ExpectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ExpectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TestDelegateSetUsage_Websocket tests the complete delegation flow with granular permissions
func TestDelegateSetUsage_Websocket(t *testing.T) {
	env := integration.GetWebsocketEnv(t)
	client := websocket.NewClient(websocket.NewClientConfig().WithHost(env.Host).WithFaucetProvider(env.FaucetProvider))

	runner := integration.NewRunner(t, client, &integration.RunnerConfig{
		WalletCount: 2,
	})

	err := runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	delegator := runner.GetWallet(0)
	delegatee := runner.GetWallet(1)

	// Step 1: Set up delegation with both transaction type and granular permissions
	delegateSetTx := &transaction.DelegateSet{
		BaseTx: transaction.BaseTx{
			Account: delegator.GetAddress(),
		},
		Authorize: delegatee.GetAddress(),
		Permissions: []types.Permission{
			{
				Permission: types.PermissionValue{
					PermissionValue: "Payment",
				},
			},
			{
				Permission: types.PermissionValue{
					PermissionValue: "TrustlineAuthorize",
				},
			},
		},
	}

	flatDelegateSet := delegateSetTx.Flatten()
	_, err = runner.TestTransaction(&flatDelegateSet, delegator, "tesSUCCESS")
	require.NoError(t, err)

	// Step 2: Use the delegation - delegated payment
	delegatedPaymentTx := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account:  delegator.GetAddress(),
			Delegate: delegatee.GetAddress(),
		},
		Destination: delegatee.GetAddress(),
		Amount:      types.XRPCurrencyAmount(1000000), // 1 XRP
	}

	flatPayment := delegatedPaymentTx.Flatten()
	_, err = runner.TestTransaction(&flatPayment, delegatee, "tesSUCCESS")
	require.NoError(t, err)
}

// TestDelegateSetUsage_RPCClient tests the complete delegation flow using RPC client
func TestDelegateSetUsage_RPCClient(t *testing.T) {
	env := integration.GetRPCEnv(t)
	clientCfg, err := rpc.NewClientConfig(env.Host, rpc.WithFaucetProvider(env.FaucetProvider))
	require.NoError(t, err)
	client := rpc.NewClient(clientCfg)

	runner := integration.NewRunner(t, client, integration.NewRunnerConfig(
		integration.WithWallets(2),
	))

	err = runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	delegator := runner.GetWallet(0)
	delegatee := runner.GetWallet(1)

	// Step 1: Set up delegation with both transaction type and granular permissions
	delegateSetTx := &transaction.DelegateSet{
		BaseTx: transaction.BaseTx{
			Account: delegator.GetAddress(),
		},
		Authorize: delegatee.GetAddress(),
		Permissions: []types.Permission{
			{
				Permission: types.PermissionValue{
					PermissionValue: "Payment",
				},
			},
			{
				Permission: types.PermissionValue{
					PermissionValue: "TrustlineAuthorize",
				},
			},
		},
	}

	flatDelegateSet := delegateSetTx.Flatten()
	_, err = runner.TestTransaction(&flatDelegateSet, delegator, "tesSUCCESS")
	require.NoError(t, err)

	// Step 2: Use the delegation - delegated payment
	delegatedPaymentTx := &transaction.Payment{
		BaseTx: transaction.BaseTx{
			Account:  delegator.GetAddress(),
			Delegate: delegatee.GetAddress(),
		},
		Destination: delegatee.GetAddress(),
		Amount:      types.XRPCurrencyAmount(1000000), // 1 XRP
	}

	flatPayment := delegatedPaymentTx.Flatten()
	_, err = runner.TestTransaction(&flatPayment, delegatee, "tesSUCCESS")
	require.NoError(t, err)
}
