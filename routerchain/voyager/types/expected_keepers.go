package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
)

type AttestationKeeper interface {
	CheckOrchestratorValidatorInSet(ctx sdk.Context, orchestrator string) error
	ClaimHandlerCommon(ctx sdk.Context, msgAny *codectypes.Any, msg attestationTypes.Claim) error
}

type MultichainKeeper interface {
	GetLastObservedEventNonce(ctx sdk.Context, chainId string, contract string) uint64
	// Methods imported from multichain should be defined here
}

type PricefeedKeeper interface {
	// Methods imported from pricefeed should be defined here
}

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
