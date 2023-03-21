package types

import (
	context "context"
	"math/big"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/x/evm/statedb"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	oracleTypes "github.com/router-protocol/sdk-go/routerchain/oracle/types"
)

type MultichainKeeper interface {
	// Methods imported from multichain should be defined here
	GetLastObservedEventNonce(ctx sdk.Context, chainType multichainTypes.ChainType, chainId string) uint64
	GetChainConfig(ctx sdk.Context, chainType multichainTypes.ChainType, chainId string) (val multichainTypes.ChainConfig, found bool)
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
	GetTokenPrice(ctx sdk.Context, oracletype oracleTypes.OracleType, symbol string) *sdk.Dec
	GetGasPrice(ctx sdk.Context, chainType multichainTypes.ChainType, chainID string) oracleTypes.GasPriceState
	ConvertNativeTokenFeeToRouter(ctx sdk.Context, chainType multichainTypes.ChainType, chainId string, feeConsumed *big.Int) sdk.Coin
}

// EVMKeeper defines the expected EVM keeper interface used on erc20
type EVMKeeper interface {
	GetParams(ctx sdk.Context) evmtypes.Params
	GetAccountWithoutBalance(ctx sdk.Context, addr common.Address) *statedb.Account
	EstimateGas(c context.Context, req *evmtypes.EthCallRequest) (*evmtypes.EstimateGasResponse, error)
	ApplyMessage(ctx sdk.Context, msg core.Message, tracer vm.EVMLogger, commit bool) (*evmtypes.MsgEthereumTxResponse, error)
	ApplyTransaction(ctx sdk.Context, tx *ethtypes.Transaction) (*evmtypes.MsgEthereumTxResponse, error)
	ApplyInternalMessage(ctx sdk.Context, msg core.Message, tracer vm.EVMLogger) (*evmtypes.MsgEthereumTxResponse, error)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authTypes.AccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)
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
