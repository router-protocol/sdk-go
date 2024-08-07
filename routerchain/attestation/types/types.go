package types

import (
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewValsetConfirmation(valsetNonce uint64, destChainType multichainTypes.ChainType, ethSigner string, signature string, orchestrator string) *ValsetConfirmation {
	return &ValsetConfirmation{
		ValsetNonce:   valsetNonce,
		DestChainType: destChainType,
		EthAddress:    ethSigner,
		Signature:     signature,
		Orchestrator:  orchestrator,
	}
}

func NewValsetUpdatedClaim(chainId string, contract string, eventNonce uint64, blockHeight uint64,
	valsetNonce uint64, sourceTxHash string,
) *ValsetUpdatedClaim {
	return &ValsetUpdatedClaim{
		ChainId:      chainId,
		Contract:     contract,
		EventNonce:   eventNonce,
		BlockHeight:  blockHeight,
		ValsetNonce:  valsetNonce,
		SourceTxHash: sourceTxHash,
	}
}

func NewValsetUpdatedClaimHash(
	chainId string,
	contract string,
	eventNonce uint64,
	blockHeight uint64,
	valsetNonce uint64,
	sourceTxHash string,
	members []BridgeValidator,
) *ValsetUpdatedClaimHash {
	return &ValsetUpdatedClaimHash{
		ChainId:      chainId,
		Contract:     contract,
		EventNonce:   eventNonce,
		BlockHeight:  blockHeight,
		ValsetNonce:  valsetNonce,
		SourceTxHash: sourceTxHash,
		Members:      members,
	}
}
