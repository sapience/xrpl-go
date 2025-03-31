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

// extractDomainID extracts the DomainID (LedgerIndex) from the Meta field
// of a transaction response. It returns an empty string if no matching node is found.
func extractDomainID(meta any) string {
	m, ok := meta.(map[string]interface{})
	if !ok {
		return ""
	}
	affectedNodes, ok := m["AffectedNodes"].([]interface{})
	if !ok {
		return ""
	}
	for _, node := range affectedNodes {
		nodeMap, ok := node.(map[string]interface{})
		if !ok {
			continue
		}
		if created, exists := nodeMap["CreatedNode"]; exists {
			createdMap, ok := created.(map[string]interface{})
			if !ok {
				continue
			}
			if entryType, exists := createdMap["LedgerEntryType"]; exists && entryType == "PermissionedDomain" {
				if id, ok := createdMap["LedgerIndex"].(string); ok {
					return id
				}
			}
		}
	}
	return ""
}

func TestIntegrationPermissionedDomainSetAndDelete_Websocket(t *testing.T) {
	env := integration.GetWebsocketEnv(t)
	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost(env.Host).
			WithFaucetProvider(env.FaucetProvider),
	)
	runner := integration.NewRunner(t, client, &integration.RunnerConfig{
		WalletCount: 1,
	})
	err := runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	wallet := runner.GetWallet(0)

	setTx := &transaction.PermissionedDomainSet{
		BaseTx: transaction.BaseTx{
			Account:         wallet.GetAddress(),
			TransactionType: transaction.PermissionedDomainSetTx,
		},
		AcceptedCredentials: []types.AuthorizeCredential{
			{
				Credential: struct {
					Issuer         types.Address
					CredentialType types.CredentialType
				}{
					Issuer:         wallet.GetAddress(),
					CredentialType: types.CredentialType("6D795F63726564656E7469616C"),
				},
			},
		},
	}
	flatSetTx := setTx.Flatten()
	txResp, _, err := runner.ProcessTransactionAndWait(&flatSetTx, wallet)
	require.NoError(t, err)
	require.True(t, txResp.Validated, "permissioned domain set transaction not validated")

	domainID := extractDomainID(txResp.Meta)
	require.NotEmpty(t, domainID, "expected DomainID from permissioned domain creation")

	delTx := &transaction.PermissionedDomainDelete{
		BaseTx: transaction.BaseTx{
			Account:         wallet.GetAddress(),
			TransactionType: transaction.PermissionedDomainDeleteTx,
		},
		DomainID: domainID,
	}
	flatDelTx := delTx.Flatten()
	delResp, _, err := runner.ProcessTransactionAndWait(&flatDelTx, wallet)
	require.NoError(t, err)
	require.True(t, delResp.Validated, "permissioned domain delete transaction not validated")
}

func TestIntegrationPermissionedDomainSetAndDelete_RPCClient(t *testing.T) {
	env := integration.GetRPCEnv(t)
	clientCfg, err := rpc.NewClientConfig(env.Host, rpc.WithFaucetProvider(env.FaucetProvider))
	require.NoError(t, err)
	client := rpc.NewClient(clientCfg)
	runner := integration.NewRunner(t, client, integration.NewRunnerConfig(integration.WithWallets(1)))
	err = runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	wallet := runner.GetWallet(0)

	setTx := &transaction.PermissionedDomainSet{
		BaseTx: transaction.BaseTx{
			Account:         wallet.GetAddress(),
			TransactionType: transaction.PermissionedDomainSetTx,
		},
		AcceptedCredentials: []types.AuthorizeCredential{
			{
				Credential: struct {
					Issuer         types.Address
					CredentialType types.CredentialType
				}{
					Issuer:         wallet.GetAddress(),
					CredentialType: types.CredentialType("6D795F63726564656E7469616C"),
				},
			},
		},
	}
	flatSetTx := setTx.Flatten()
	txResp, _, err := runner.ProcessTransactionAndWait(&flatSetTx, wallet)
	require.NoError(t, err)
	require.True(t, txResp.Validated, "permissioned domain set transaction not validated")

	domainID := extractDomainID(txResp.Meta)
	require.NotEmpty(t, domainID, "expected DomainID from permissioned domain creation")

	delTx := &transaction.PermissionedDomainDelete{
		BaseTx: transaction.BaseTx{
			Account:         wallet.GetAddress(),
			TransactionType: transaction.PermissionedDomainDeleteTx,
		},
		DomainID: domainID,
	}
	flatDelTx := delTx.Flatten()
	delResp, _, err := runner.ProcessTransactionAndWait(&flatDelTx, wallet)
	require.NoError(t, err)
	require.True(t, delResp.Validated, "permissioned domain delete transaction not validated")
}
