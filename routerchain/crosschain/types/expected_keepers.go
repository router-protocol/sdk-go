package types

import (
	context "context"

	"cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v6/modules/apps/29-fee/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/x/evm/statedb"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	metastoreTypes "github.com/router-protocol/sdk-go/routerchain/metastore/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	pricefeedTypes "github.com/router-protocol/sdk-go/routerchain/pricefeed/types"
)

type AttestationKeeper interface {
	// Methods imported from attestation should be defined here
	CheckOrchestratorValidatorInSet(ctx sdk.Context, orchestrator string) error
	ClaimHandlerCommon(ctx sdk.Context, msgAny *codectypes.Any, msg attestationTypes.Claim) error
	ConfirmHandlerCommon(ctx sdk.Context, ethAddress string, orchestrator sdk.AccAddress, signature string, checkpoint []byte) error
	GetRouterID(ctx sdk.Context) string
}

type MetastoreKeeper interface {
	GetMetaInfo(ctx sdk.Context, chainId string, dappAddress string) (val metastoreTypes.MetaInfo, found bool)
	GetMetadataRequest(ctx sdk.Context, chainId string, eventNonce uint64, claimHash []byte) (val metastoreTypes.MetadataRequest, found bool)
	SetMetadataRequest(ctx sdk.Context, claimHash []byte, metadataRequest metastoreTypes.MetadataRequest)
}

type MultichainKeeper interface {
	// Methods imported from multichain should be defined here
	GetLastObservedEventNonce(ctx sdk.Context, chainId string, contract string) uint64
	SetLastObservedEventNonce(ctx sdk.Context, chainId string, contract string, nonce uint64)
	GetLastObservedEventBlockHeight(ctx sdk.Context, contract string, chainId string) uint64
	SetLastObservedEventBlockHeight(ctx sdk.Context, chainId string, contract string, blockHeight uint64)
	GetAllChainConfig(ctx sdk.Context) (list []multichainTypes.ChainConfig)
	GetAllContractConfig(ctx sdk.Context) (list []multichainTypes.ContractConfig)
	GetChainConfig(ctx sdk.Context, chainId string) (chainConfig multichainTypes.ChainConfig, found bool)
	GetContractConfig(ctx sdk.Context, chainId string, contract string) (val multichainTypes.ContractConfig, found bool)
}

type PriceFeedKeeper interface {
	// Methods imported from oracle should be defined here
	GetTokenPriceState(ctx sdk.Context, symbol string) (*pricefeedTypes.Price, error)
	GetGasPriceState(ctx sdk.Context, chainID string) (pricefeedTypes.GasPrice, error)
	ConvertNativeTokenFeeToRouter(ctx sdk.Context, chainId string, feeConsumedInDecimals math.Int) (sdk.Coin, error)
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
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}

type WasmKeeper interface {
	// Methods imported from wasmd should be defined here
	Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
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

// EVMKeeper defines the expected EVM keeper interface used on erc20
type IBCFeeKeeper interface {
	TotalRecvFees(goCtx context.Context, req *ibcfeetypes.QueryTotalRecvFeesRequest) (*ibcfeetypes.QueryTotalRecvFeesResponse, error)
	TotalAckFees(goCtx context.Context, req *ibcfeetypes.QueryTotalAckFeesRequest) (*ibcfeetypes.QueryTotalAckFeesResponse, error)
	TotalTimeoutFees(goCtx context.Context, req *ibcfeetypes.QueryTotalTimeoutFeesRequest) (*ibcfeetypes.QueryTotalTimeoutFeesResponse, error)
}
