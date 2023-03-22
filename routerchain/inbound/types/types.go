package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewIncomingTx(chainType multichainTypes.ChainType, chainId string, eventNonce uint64, blockHeight uint64, sourceSender []byte, sourceTxHash string, sourceTimeStamp uint64, routerBridgeContract string, gasLimit uint64, routeAmount sdk.Int, routeRecipient string, payload []byte, asmAddress []byte) *IncomingTx {
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
		RouteRecipient:       routeRecipient,
		RouteAmount:          routeAmount,
		AsmAddress:           asmAddress,
		Status:               INCOMING_TX_CREATED,
	}
}

func NewInboundRequestClaimHash(
	chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	blockHeight uint64,
	sourceSender []byte,
	sourceTxHash string,
	sourceTimeStamp uint64,
	routerBridgeContract string,
	gasLimit uint64,
	routeAmount sdk.Int,
	routeRecipient string,
	asmAddress []byte,
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
		AsmAddress:           asmAddress,
	}
}
