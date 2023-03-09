package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewIncomingTx(chainType multichainTypes.ChainType, chainId string, eventNonce uint64, blockHeight uint64, sourceSender string, sourceTxHash string, sourceTimeStamp uint64, routerBridgeContract string, gasLimit uint64, routeAmount sdk.Int, routeRecipient []byte, feePayer []byte, payload []byte) *IncomingTx {
	return &IncomingTx{
		ChainType:            chainType,
		ChainId:              chainId,
		EventNonce:           eventNonce,
		BlockHeight:          blockHeight,
		SourceSender:         sourceSender,
		SourceTimestamp:      sourceTimeStamp,
		SourceTxHash:         sourceTxHash,
		RouterBridgeContract: routerBridgeContract,
		Payload:              payload,
		GasLimit:             gasLimit,
		FeePayer:             feePayer,
		RouteRecipient:       routeRecipient,
		RouteAmount:          routeAmount,
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
	sourceTimeStamp uint64,
	routerBridgeContract string,
	gasLimit uint64,
	routeAmount sdk.Int,
	routeRecipient []byte,
	payload []byte) *InboundRequestClaimHash {
	return &InboundRequestClaimHash{
		ChainType:            chainType,
		ChainId:              chainId,
		EventNonce:           eventNonce,
		BlockHeight:          blockHeight,
		SourceSender:         sourceSender,
		SourceTxHash:         sourceTxHash,
		SourceTimestamp:      sourceTimeStamp,
		RouteRecipient:       routeRecipient,
		RouteAmount:          routeAmount,
		RouterBridgeContract: routerBridgeContract,
		Payload:              payload,
		GasLimit:             gasLimit,
	}
}
