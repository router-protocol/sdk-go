package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	metastoreTypes "github.com/router-protocol/sdk-go/routerchain/metastore/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	oracleTypes "github.com/router-protocol/sdk-go/routerchain/oracle/types"
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
	GetLastObservedEventNonce(ctx sdk.Context, chainType multichainTypes.ChainType, chainId string) uint64
	GetChainConfig(ctx sdk.Context, chainType multichainTypes.ChainType, chainId string) (chainConfig multichainTypes.ChainConfig, found bool)
}

type MetastoreKeeper interface {
	GetMetaInfo(ctx sdk.Context, chainType multichainTypes.ChainType, chainId string, dappAddress []byte) (val metastoreTypes.MetaInfo, found bool)
}

type OracleKeeper interface {
	// Methods imported from oracle should be defined here
	GetGasPrice(ctx sdk.Context, chainType multichainTypes.ChainType, chainID string) oracleTypes.GasPriceState
	GetTokenPrice(ctx sdk.Context, oracletype oracleTypes.OracleType, symbol string) *sdk.Dec
	ConvertNativeTokenFeeToRouter(ctx sdk.Context, chainType multichainTypes.ChainType, chainId string, feeConsumed *big.Int) sdk.Coin
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	// Methods imported from bank should be defined here
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SetDenomMetaData(ctx sdk.Context, denomMeta banktypes.Metadata)
}
