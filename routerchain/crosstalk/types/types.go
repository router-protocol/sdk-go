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
	ackType CrossTalkRequestAckType,
	ackGasLimit uint64,
	ackGasPrice uint64,
	requestTxOrigin string,
	isReadCall bool,
	feePayer []byte) *CrossTalkRequest {
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
		RequestTxOrigin:       requestTxOrigin,
		IsReadCall:            isReadCall,
		RequestNonce:          requestNonce,
		IsAtomic:              isAtomic,
		ExpiryTimestamp:       expiryTimestamp,
		DestContractAddresses: destContractAddresses,
		DestContractPayloads:  destContractPayloads,
		AckType:               ackType,
		AckGasLimit:           ackGasLimit,
		AckGasPrice:           ackGasPrice,
		FeePayer:              feePayer,
		Status:                CROSSTALK_REQUEST_CREATED,
	}
}

func NewCrossTalkRequestClaimHash(
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
	ackType CrossTalkRequestAckType,
	ackGasLimit uint64,
	ackGasPrice uint64,
	requestTxOrigin string,
	isReadCall bool,
	feePayer []byte) *CrossTalkRequestClaimHash {
	return &CrossTalkRequestClaimHash{
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
		RequestTxOrigin:       requestTxOrigin,
		IsReadCall:            isReadCall,
		RequestNonce:          requestNonce,
		IsAtomic:              isAtomic,
		ExpiryTimestamp:       expiryTimestamp,
		DestContractAddresses: destContractAddresses,
		DestContractPayloads:  destContractPayloads,
		AckType:               ackType,
		AckGasLimit:           ackGasLimit,
		AckGasPrice:           ackGasPrice,
		FeePayer:              feePayer,
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

func NewCrossTalkAckRequestClaimHash(
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
) *CrossTalkAckRequestClaimHash {
	return &CrossTalkAckRequestClaimHash{
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

func NewCrossTalkAckReceipt(
	chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	blockHeight uint64,
	relayerRouterAddress string,
	txHash string,
	eventIdentifier uint64,
) *CrossTalkAckReceipt {
	return &CrossTalkAckReceipt{
		ChainType:            chainType,
		ChainId:              chainId,
		EventNonce:           eventNonce,
		BlockHeight:          blockHeight,
		RelayerRouterAddress: relayerRouterAddress,
		TxHash:               txHash,
		EventIdentifier:      eventIdentifier,
		Status:               CROSSTALK_ACK_RECEIPT_CREATED,
	}
}

func NewCrossTalkAckReceiptClaimHash(
	chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	blockHeight uint64,
	relayerRouterAddress string,
	txHash string,
	eventIdentifier uint64,
) *CrossTalkAckReceiptClaimHash {
	return &CrossTalkAckReceiptClaimHash{
		ChainType:            chainType,
		ChainId:              chainId,
		EventNonce:           eventNonce,
		BlockHeight:          blockHeight,
		RelayerRouterAddress: relayerRouterAddress,
		TxHash:               txHash,
		EventIdentifier:      eventIdentifier,
	}
}
