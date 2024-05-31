package types

import (
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AttestationKeeper interface {
	CheckOrchestratorValidatorInSet(ctx sdk.Context, orchestrator string) error
	ClaimHandlerCommon(ctx sdk.Context, msgAny *codectypes.Any, msg attestationTypes.Claim) error
	TallyAttestation(ctx sdk.Context, chainId string, contract string, nonce uint64)
}

type MultichainKeeper interface {
	GetLastObservedEventNonce(ctx sdk.Context, chainId string, contract string) uint64
	GetAllIbcRelayerConfig(ctx sdk.Context) (list []multichainTypes.IbcRelayerConfig)
	GetIbcRelayerConfig(ctx sdk.Context, chainId string, connectionType multichainTypes.IbcRelayerConnectionType) (val multichainTypes.IbcRelayerConfig, found bool)
	GetChainConfig(ctx sdk.Context, chainId string) (chainConfig multichainTypes.ChainConfig, found bool)
	GetContractConfig(ctx sdk.Context, chainId string, contract string) (val multichainTypes.ContractConfig, found bool)
	GetAllContractConfigByChain(ctx sdk.Context, chainId string) (list []multichainTypes.ContractConfig)
	IsNonceObserved(ctx sdk.Context, chainId string, contractAddress string, nonce uint64) (isObserved bool)
}

type WasmKeeper interface {
	// Methods imported from wasmd should be defined here
	Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
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
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}
