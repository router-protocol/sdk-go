package types

import (
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error

	// Methods imported from bank should be defined here
}

type MultichainKeeper interface {
	GetChainConfig(ctx sdk.Context, chainId string) (val multichainTypes.ChainConfig, found bool)
	IsNonceObserved(ctx sdk.Context, chainId string, contractAddress string, nonce uint64) (isObserved bool)
	// Methods imported from multichain should be defined here
}

type AttestationKeeper interface {
	// Methods imported from attestation should be defined here
	CheckOrchestratorValidatorInSet(ctx sdk.Context, orchestrator string) error
	ClaimHandlerCommon(ctx sdk.Context, msgAny *codectypes.Any, msg attestationTypes.Claim, isAdhoc bool) error
	TallyAdhocAttestation(ctx sdk.Context, chainId string, txHash string, claimHash []byte)
	GetObservedAttestation(ctx sdk.Context, chainId string, contract string, eventNonce uint64, claimHash []byte) (val attestationTypes.Attestation, found bool)
}

type WasmKeeper interface {
	// Methods imported from wasmd should be defined here
	Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
}

type StakingKeeper interface {
	GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) int64
}
