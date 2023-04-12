package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

type AttestationKeeper interface {
	// Methods imported from attestation should be defined here
	CheckOrchestratorValidatorInSet(ctx sdk.Context, orchestrator string) error
	ClaimHandlerCommon(ctx sdk.Context, msgAny *codectypes.Any, msg attestationTypes.Claim) error
	ConfirmHandlerCommon(ctx sdk.Context, ethAddress string, orchestrator sdk.AccAddress, signature string, checkpoint []byte) error
	GetRouterID(ctx sdk.Context) string
}

type MultichainKeeper interface {
	// Methods imported from multichain should be defined here
	GetLastObservedEventNonce(ctx sdk.Context, chainId string) uint64
	SetLastObservedEventNonce(ctx sdk.Context, chainId string, nonce uint64)
	GetLastObservedEventBlockHeight(ctx sdk.Context, chainId string) uint64
	SetLastObservedEventBlockHeight(ctx sdk.Context, chainId string, blockHeight uint64)
	SetLastObservedValsetNonce(ctx sdk.Context, chainId string, nonce uint64)
	GetLastObservedValsetNonce(ctx sdk.Context, chainId string) uint64
	GetAllChainConfig(ctx sdk.Context) (list []multichainTypes.ChainConfig)
	GetChainConfig(ctx sdk.Context, chainId string) (chainConfig multichainTypes.ChainConfig, found bool)
}

type OracleKeeper interface {
	// Methods imported from oracle should be defined here
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
