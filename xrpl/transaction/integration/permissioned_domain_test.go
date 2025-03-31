package integration

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/rpc"
	"github.com/Peersyst/xrpl-go/xrpl/testutil/integration"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/stretchr/testify/require"
)

func TestIntegrationPermissionedDomainSetAndDelete_Websocket(t *testing.T) {
	// Setup integration environment using websocket.
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

	// --- Create a new permissioned domain using PermissionedDomainSet ---
	setTx := &transaction.PermissionedDomainSet{
		BaseTx: transaction.BaseTx{
			Account:         wallet.GetAddress(),
			TransactionType: transaction.PermissionedDomainSetTx,
		},
		// Omit DomainID to create a new domain.
		AcceptedCredentials: []types.AuthorizeCredential{
			{
				// Using the wallet address as the issuer.
				Issuer:         wallet.GetAddress(),
				CredentialType: types.CredentialType("6D795F63726564656E7469616C"),
			},
		},
	}
	flatSetTx := setTx.Flatten() // returns FlatTransaction (a map[string]interface{})
	txResp, _, err := runner.ProcessTransactionAndWait(&flatSetTx, wallet)
	require.NoError(t, err)
	require.True(t, txResp.Validated, "permissioned domain set transaction not validated")

	// Inline extraction of DomainID from the meta field.
	var domainID string
	meta, ok := txResp.Meta.(map[string]interface{})
	require.True(t, ok, "expected meta field in tx response")
	affectedNodes, ok := meta["AffectedNodes"].([]interface{})
	require.True(t, ok, "expected AffectedNodes in meta")
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
					domainID = id
					break
				}
			}
		}
	}
	require.NotEmpty(t, domainID, "expected DomainID from permissioned domain creation")
	t.Logf("Created DomainID: %s", domainID)

	// --- Delete the created domain using PermissionedDomainDelete ---
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
	// Setup integration environment using RPC.
	env := integration.GetRPCEnv(t)
	clientCfg, err := rpc.NewClientConfig(env.Host, rpc.WithFaucetProvider(env.FaucetProvider))
	require.NoError(t, err)
	client := rpc.NewClient(clientCfg)

	runner := integration.NewRunner(t, client, integration.NewRunnerConfig(
		integration.WithWallets(1),
	))
	err = runner.Setup()
	require.NoError(t, err)
	defer runner.Teardown()

	wallet := runner.GetWallet(0)

	// --- Create domain ---
	setTx := &transaction.PermissionedDomainSet{
		BaseTx: transaction.BaseTx{
			Account:         wallet.GetAddress(),
			TransactionType: transaction.PermissionedDomainSetTx,
		},
		AcceptedCredentials: []types.AuthorizeCredential{
			{
				Issuer:         wallet.GetAddress(),
				CredentialType: types.CredentialType("6D795F63726564656E7469616C"),
			},
		},
	}
	flatSetTx := setTx.Flatten()
	txResp, _, err := runner.ProcessTransactionAndWait(&flatSetTx, wallet)
	require.NoError(t, err)
	require.True(t, txResp.Validated, "permissioned domain set transaction not validated")
	jsonData, err := json.MarshalIndent(txResp, "", "  ")
	if err != nil {
		log.Printf("failed to marshal tx response: %v", err)
	} else {
		log.Printf("Transaction Response:\n%s", string(jsonData))
	}
	// Inline extraction of DomainID from the meta field.
	var domainID string
	meta, ok := txResp.Meta.(map[string]interface{})
	require.True(t, ok, "expected meta field in tx response")
	affectedNodes, ok := meta["AffectedNodes"].([]interface{})
	require.True(t, ok, "expected AffectedNodes in meta")
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
					domainID = id
					break
				}
			}
		}
	}
	require.NotEmpty(t, domainID, "expected DomainID from permissioned domain creation")
	t.Logf("Created DomainID: %s", domainID)

	// --- Delete domain ---
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
