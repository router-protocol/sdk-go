package types

import (
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

type AttestationKeeper interface {
	// Methods imported from attestation should be defined here
	CheckOrchestratorValidatorInSet(ctx sdk.Context, orchestrator string) error
	ClaimHandlerCommon(ctx sdk.Context, msgAny *codectypes.Any, msg attestationTypes.Claim, isAdhoc bool) error
	TallyAttestation(ctx sdk.Context, chainId string, contract string, nonce uint64)
}

type MultichainKeeper interface {
	GetLastObservedEventNonce(ctx sdk.Context, chainId string, contract string) uint64
	GetChainConfig(ctx sdk.Context, chainId string) (chainConfig multichainTypes.ChainConfig, found bool)
	GetContractConfig(ctx sdk.Context, chainId string, contract string) (val multichainTypes.ContractConfig, found bool)
	IsNonceObserved(ctx sdk.Context, chainId string, contractAddress string, nonce uint64) (isObserved bool)
	// Methods imported from multichain should be defined here
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	// Methods imported from bank should be defined here
}
