package types

import (
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewValsetConfirmation(valsetNonce uint64, ethSigner string, signature string, orchestrator string) *ValsetConfirmation {
	return &ValsetConfirmation{
		ValsetNonce:  valsetNonce,
		EthAddress:   ethSigner,
		Signature:    signature,
		Orchestrator: orchestrator,
	}
}

func NewValsetUpdatedClaim(chainType multichainTypes.ChainType, chainId string, eventNonce uint64, blockHeight uint64,
	valsetNonce uint64,
) *ValsetUpdatedClaim {
	return &ValsetUpdatedClaim{
		ChainType:   chainType,
		ChainId:     chainId,
		EventNonce:  eventNonce,
		BlockHeight: blockHeight,
		ValsetNonce: valsetNonce,
	}
}
