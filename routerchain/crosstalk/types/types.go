package types

import multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"

func NewCrossTalkRequest(eventNonce uint64, blockHeight uint64,
	sourceChainType multichainTypes.ChainType, sourceChainId string, sourceTxHash string,
	destinationChainType multichainTypes.ChainType, destinationChainId string,
	requestSender string, requestNonce uint64, isAtomic bool,
	gasLimit uint64, gasPrice uint64, expiryTimestamp uint64, ethSigner string,
	signature string, contractCalls []CrosstalkContractCall) *CrossTalkRequest {
	return &CrossTalkRequest{
		EventNonce:           eventNonce,
		BlockHeight:          blockHeight,
		SourceChainType:      sourceChainType,
		SourceChainId:        sourceChainId,
		SourceTxHash:         sourceTxHash,
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		RequestSender:        requestSender,
		RequestNonce:         requestNonce,
		IsAtomic:             isAtomic,
		GasLimit:             gasLimit,
		GasPrice:             gasPrice,
		ExpiryTimestamp:      expiryTimestamp,
		EthSigner:            ethSigner,
		Signature:            signature,
		ContractCalls:        contractCalls,
		Status:               CROSSTALK_REQUEST_CREATED,
	}

}
