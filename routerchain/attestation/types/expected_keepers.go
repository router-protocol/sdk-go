package types

import (
	"time"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"

	"cosmossdk.io/math"
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type MultichainKeeper interface {
	// Methods imported from multichain should be defined here
	GetLastObservedEventNonce(ctx sdk.Context, chainId string, contract string) uint64
	SetLastObservedEventNonce(ctx sdk.Context, chainId string, contract string, nonce uint64)
	GetLastObservedEventBlockHeight(ctx sdk.Context, contract string, chainId string) uint64
	SetLastObservedEventBlockHeight(ctx sdk.Context, chainId string, contract string, blockHeight uint64)
	SetNonceObservedStatus(ctx sdk.Context, nonceObservedStatus multichainTypes.NonceObservedStatus)
	SetLastObservedValsetNonce(ctx sdk.Context, chainId string, nonce uint64)
	GetLastObservedValsetNonce(ctx sdk.Context, chainId string) uint64
	GetAllChainConfig(ctx sdk.Context) (list []multichainTypes.ChainConfig)
	GetAllContractConfig(ctx sdk.Context) (list []multichainTypes.ContractConfig)
	GetChainConfig(ctx sdk.Context, chainId string) (chainConfig multichainTypes.ChainConfig, found bool)
	GetContractConfig(ctx sdk.Context, chainId string, contract string) (val multichainTypes.ContractConfig, found bool)
	IsNonceObserved(ctx sdk.Context, chainId string, contractAddress string, nonce uint64) (isObserved bool)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) auth.AccountI
	// Methods imported from account should be defined here
}

// StakingKeeper defines the expected staking keeper methods
type StakingKeeper interface {
	GetBondedValidatorsByPower(ctx sdk.Context) []stakingtypes.Validator
	GetAllValidators(ctx sdk.Context) (validators []stakingtypes.Validator)
	GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) int64
	GetLastTotalPower(ctx sdk.Context) (power sdkmath.Int)
	ValidatorQueueIterator(ctx sdk.Context, endTime time.Time, endHeight int64) sdk.Iterator
	GetParams(ctx sdk.Context) stakingtypes.Params
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	IterateBondedValidatorsByPower(sdk.Context, func(index int64, validator stakingtypes.ValidatorI) (stop bool))
	IterateLastValidators(sdk.Context, func(index int64, validator stakingtypes.ValidatorI) (stop bool))

	// iterate through validators by operator address, execute func for each validator
	IterateValidators(sdk.Context,
		func(index int64, validator stakingtypes.ValidatorI) (stop bool))

	Validator(sdk.Context, sdk.ValAddress) stakingtypes.ValidatorI            // get a particular validator by operator address
	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingtypes.ValidatorI // get a particular validator by consensus address

	// slash the validator and delegators of the validator, specifying offence height, offence power, and slash fraction
	Slash(sdk.Context, sdk.ConsAddress, int64, int64, sdk.Dec) math.Int
	Jail(sdk.Context, sdk.ConsAddress)   // jail a validator
	Unjail(sdk.Context, sdk.ConsAddress) // unjail a validator

	// Delegation allows for getting a particular delegation for a given validator
	// and delegator outside the scope of the staking module.
	Delegation(sdk.Context, sdk.AccAddress, sdk.ValAddress) stakingtypes.DelegationI

	// MaxValidators returns the maximum amount of bonded validators
	MaxValidators(sdk.Context) uint32

	GetHistoricalInfo(ctx sdk.Context, height int64) (stakingtypes.HistoricalInfo, bool)
	GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator stakingtypes.Validator, found bool)
}

// BankKeeper defines the expected bank keeper methods
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type SlashingKeeper interface {
	GetValidatorSigningInfo(ctx sdk.Context, address sdk.ConsAddress) (info slashingtypes.ValidatorSigningInfo, found bool)
}

type DistributionKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
	GetFeePool(ctx sdk.Context) (feePool types.FeePool)
	SetFeePool(ctx sdk.Context, feePool types.FeePool)
}
