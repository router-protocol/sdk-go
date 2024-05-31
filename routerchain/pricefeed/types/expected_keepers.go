package types

import (
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

type MultichainKeeper interface {
	GetChainConfig(ctx sdk.Context, chainId string) (val multichainTypes.ChainConfig, found bool)
	// Methods imported from multichain should be defined here
}
