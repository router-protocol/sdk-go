package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ROUTER defines the default coin denomination used in Ethermint in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Ethermint.
	RouterCoin string = "route"

	// BaseDenomUnit defines the base denomination unit for Photons.
	// 1 photon = 1x10^{BaseDenomUnit} router
	BaseDenomUnit = 18
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewRouterCoin is a utility function that returns an "router" coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewRouterCoin(amount sdk.Int) sdk.Coin {
	return sdk.NewCoin(RouterCoin, amount)
}

// NewRouterDecCoin is a utility function that returns an "router" decimal coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewRouterDecCoin(amount sdk.Int) sdk.DecCoin {
	return sdk.NewDecCoin(RouterCoin, amount)
}

// NewRouterCoinInt64 is a utility function that returns an "router" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewRouterCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(RouterCoin, amount)
}
