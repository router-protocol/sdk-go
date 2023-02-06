package types

import (
	"math/big"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	oracleTypes "github.com/router-protocol/sdk-go/routerchain/oracle/types"
)

type MultichainKeeper interface {
	// Methods imported from multichain should be defined here
	GetChainConfig(ctx sdk.Context, chainType multichainTypes.ChainType, chainId string) (val multichainTypes.ChainConfig, found bool)
	ConvertNativeTokenFeeToRouter(ctx sdk.Context, chainType multichainTypes.ChainType, chainId string, feeConsumed *big.Int) sdk.Coin
}

type AttestationKeeper interface {
	// Methods imported from attestation should be defined here
	ConfirmHandlerCommon(ctx sdk.Context, ethAddress string, orchestrator sdk.AccAddress, signature string, checkpoint []byte) error
	GetRouterID(ctx sdk.Context) string
	CheckOrchestratorValidatorInSet(ctx sdk.Context, orchestrator string) error
	ClaimHandlerCommon(ctx sdk.Context, msgAny *codectypes.Any, msg attestationTypes.Claim) error
}

type OracleKeeper interface {
	// Methods imported from oracle should be defined here
	GetPrice(ctx sdk.Context, oracletype oracleTypes.OracleType, symbol string) *sdk.Dec
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

type WasmKeeper interface {
	// Methods imported from wasmd should be defined here
	Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
}
