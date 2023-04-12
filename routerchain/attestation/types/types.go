package types

func NewValsetConfirmation(valsetNonce uint64, ethSigner string, signature string, orchestrator string) *ValsetConfirmation {
	return &ValsetConfirmation{
		ValsetNonce:  valsetNonce,
		EthAddress:   ethSigner,
		Signature:    signature,
		Orchestrator: orchestrator,
	}
}

func NewValsetUpdatedClaim(chainId string, eventNonce uint64, blockHeight uint64,
	valsetNonce uint64, sourceTxHash string,
) *ValsetUpdatedClaim {
	return &ValsetUpdatedClaim{
		ChainId:      chainId,
		EventNonce:   eventNonce,
		BlockHeight:  blockHeight,
		ValsetNonce:  valsetNonce,
		SourceTxHash: sourceTxHash,
	}
}

func NewValsetUpdatedClaimHash(
	chainId string,
	eventNonce uint64,
	blockHeight uint64,
	valsetNonce uint64,
	sourceTxHash string,
	members []BridgeValidator,
) *ValsetUpdatedClaimHash {
	return &ValsetUpdatedClaimHash{
		ChainId:      chainId,
		EventNonce:   eventNonce,
		BlockHeight:  blockHeight,
		ValsetNonce:  valsetNonce,
		SourceTxHash: sourceTxHash,
		Members:      members,
	}
}
