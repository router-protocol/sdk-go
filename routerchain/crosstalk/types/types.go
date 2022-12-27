package types

import (
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewCrossTalkRequest(
	eventNonce uint64,
	blockHeight uint64,
	sourceChainType multichainTypes.ChainType,
	sourceChainId string,
	sourceTxHash string,
	destinationChainType multichainTypes.ChainType,
	destinationChainId string,
	destinationGasLimit uint64,
	destinationGasPrice uint64,
	requestSender string,
	requestNonce uint64,
	isAtomic bool,
	expiryTimestamp uint64,
	destContractAddresses [][]byte,
	destContractPayloads [][]byte,
	ackType uint64,
	ackGasLimit uint64,
	ackGasPrice uint64) *CrossTalkRequest {
	return &CrossTalkRequest{
		EventNonce:            eventNonce,
		BlockHeight:           blockHeight,
		SourceChainType:       sourceChainType,
		SourceChainId:         sourceChainId,
		SourceTxHash:          sourceTxHash,
		DestinationChainType:  destinationChainType,
		DestinationChainId:    destinationChainId,
		DestinationGasLimit:   destinationGasLimit,
		DestinationGasPrice:   destinationGasPrice,
		RequestSender:         requestSender,
		RequestNonce:          requestNonce,
		IsAtomic:              isAtomic,
		ExpiryTimestamp:       expiryTimestamp,
		DestContractAddresses: destContractAddresses,
		DestContractPayloads:  destContractPayloads,
		AckType:               ackType,
		AckGasLimit:           ackGasLimit,
		AckGasPrice:           ackGasPrice,
		Status:                CROSSTALK_REQUEST_CREATED,
	}
}

func NewCrosstalkRequestConfirm(sourceChainType multichainTypes.ChainType, sourceChainId string, eventNonce uint64, claimHash []byte, orchestrator string, ethSigner string, signature string) *CrosstalkRequestConfirm {
	return &CrosstalkRequestConfirm{
		SourceChainType: sourceChainType,
		SourceChainId:   sourceChainId,
		EventNonce:      eventNonce,
		ClaimHash:       claimHash,
		Orchestrator:    orchestrator,
		EthSigner:       ethSigner,
		Signature:       signature,
	}
}