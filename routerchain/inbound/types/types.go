package types

import (
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewIncomingTx(chainType multichainTypes.ChainType, chainId string, eventNonce uint64, blockHeight uint64, sourceSender string, sourceTxHash string, routerBridgeContract string, gasLimit uint64, feePayer []byte, payload []byte) *IncomingTx {
	return &IncomingTx{
		ChainType:            chainType,
		ChainId:              chainId,
		EventNonce:           eventNonce,
		BlockHeight:          blockHeight,
		SourceSender:         sourceSender,
		SourceTxHash:         sourceTxHash,
		RouterBridgeContract: routerBridgeContract,
		Payload:              payload,
		GasLimit:             gasLimit,
		FeePayer:             feePayer,
		Status:               INCOMING_TX_CREATED,
	}
}

func NewInboundRequestClaimHash(
	chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	blockHeight uint64,
	sourceSender string,
	sourceTxHash string,
	routerBridgeContract string,
	gasLimit uint64,
	payload []byte) *InboundRequestClaimHash {
	return &InboundRequestClaimHash{
		ChainType:            chainType,
		ChainId:              chainId,
		EventNonce:           eventNonce,
		BlockHeight:          blockHeight,
		SourceSender:         sourceSender,
		SourceTxHash:         sourceTxHash,
		RouterBridgeContract: routerBridgeContract,
		Payload:              payload,
		GasLimit:             gasLimit,
	}
}
