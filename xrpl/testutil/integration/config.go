package integration

import "github.com/Peersyst/xrpl-go/xrpl/websocket"

const (
	DefaultMaxRetries = 3
)

var (
	DefaultClient Client = websocket.NewClient(websocket.NewClientConfig().WithHost(IntegrationEnvs[LocalnetEnv].Host))
)

// RunnerConfig is the configuration for the integration test runner.
// It contains the configuration for the websocket client and the number of wallets to create.
type RunnerConfig struct {
	WalletCount int
	Client      Client
	MaxRetries  int
}

// Option is a function that modifies the RunnerConfig.
type Option func(*RunnerConfig)

// WithClient sets the client for the runner.
func WithClient(client Client) Option {
	return func(c *RunnerConfig) {
		c.Client = client
	}
}

// WithWallets sets the number of wallets to create.
func WithWallets(count int) Option {
	return func(c *RunnerConfig) {
		c.WalletCount = count
	}
}

// WithMaxRetries sets the maximum number of retries for a transaction.
func WithMaxRetries(maxRetries int) Option {
	return func(c *RunnerConfig) {
		c.MaxRetries = maxRetries
	}
}

// NewRunnerConfig creates a new RunnerConfig with the given websocket configuration and options.
func NewRunnerConfig(opts ...Option) *RunnerConfig {
	config := &RunnerConfig{
		Client:      DefaultClient,
		WalletCount: 0,
		MaxRetries:  DefaultMaxRetries,
	}

	for _, opt := range opts {
		opt(config)
	}

	return config
}
