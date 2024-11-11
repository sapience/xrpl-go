package transaction

// Common flags for AMM transactions (Deposit and Withdraw).
const (
	// Perform a double-asset withdrawal/deposit and receive the specified amount of LP Tokens.
	tfLPToken uint32 = 65536
	// Perform a single-asset withdrawal/deposit with a specified amount of the asset to deposit.
	tfSingleAsset uint32 = 524288
	// Perform a double-asset withdrawal/deposit with specified amounts of both assets.
	tfTwoAsset uint32 = 1048576
	// Perform a single-asset withdrawal/deposit and receive the specified amount of LP Tokens.
	tfOneAssetLPToken uint32 = 2097152
	// Perform a single-asset withdrawal/deposit with a specified effective price.
	tfLimitLPToken uint32 = 4194304
)
