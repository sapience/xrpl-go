package integration

import (
	"os"
	"testing"
)

// GetEnv returns the integration environment.
// If the environment is not set, it skips the tests.
// This function is intended to be used in tests that need to run against a specific environment.
// Run it before creating the runner to retrieve the environment host and faucet provider.
func GetEnv(t *testing.T) Env {
	if _, ok := IntegrationEnvs[EnvKey(os.Getenv("INTEGRATION"))]; !ok {
		t.Skip("skipping integration tests")
	}

	return IntegrationEnvs[EnvKey(os.Getenv("INTEGRATION"))]
}
