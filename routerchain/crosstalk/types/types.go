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
	requestSender []byte,
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

func NewCrossTalkAckRequest(
	eventNonce uint64,
	blockHeight uint64,
	relayerRouterAddress string,
	sourceChainType multichainTypes.ChainType,
	sourceChainId string,
	chainType multichainTypes.ChainType,
	chainId string,
	destinationTxHash string,
	eventIdentifier uint64,
	crosstalkRequestSender []byte,
	crosstalkNonce uint64,
	contractAckResponses []byte,
	exeCode uint64,
	execStatus bool,
	execFlags []bool,
	execData [][]byte,
) *CrossTalkAckRequest {
	return &CrossTalkAckRequest{
		EventNonce:             eventNonce,
		BlockHeight:            blockHeight,
		RelayerRouterAddress:   relayerRouterAddress,
		SourceChainType:        sourceChainType,
		SourceChainId:          sourceChainId,
		ChainType:              chainType,
		ChainId:                chainId,
		DestinationTxHash:      destinationTxHash,
		EventIdentifier:        eventIdentifier,
		CrosstalkRequestSender: crosstalkRequestSender,
		CrosstalkNonce:         crosstalkNonce,
		ContractAckResponses:   contractAckResponses,
		ExeCode:                exeCode,
		ExecStatus:             execStatus,
		ExecFlags:              execFlags,
		ExecData:               execData,
		Status:                 CROSSTALK_ACK_REQUEST_CREATED,
	}
}

func NewCrosstalkAckRequestConfirm(chainType multichainTypes.ChainType, chainId string, eventNonce uint64, claimHash []byte, orchestrator string, ethSigner string, signature string) *CrosstalkAckRequestConfirm {
	return &CrosstalkAckRequestConfirm{
		ChainType:    chainType,
		ChainId:      chainId,
		EventNonce:   eventNonce,
		ClaimHash:    claimHash,
		Orchestrator: orchestrator,
		EthSigner:    ethSigner,
		Signature:    signature,
	}
}
