package types

func NewCrosschainRequest(
	srcEventNonce uint64,
	srcBlockHeight uint64,
	srcChainId string,
	sourceTxHash string,
	srcTimestamp uint64,
	srcTxOrigin string,
	routeAmount string,
	routeRecipient string,
	destChainId string,
	requestSender string,
	requestMetadata []byte,
	requestPayload []byte,
	relayerFee string,
) *CrosschainRequest {
	return &CrosschainRequest{
		SrcChainId:      srcChainId,
		SrcEventNonce:   srcEventNonce,
		SrcBlockHeight:  srcBlockHeight,
		SourceTxHash:    sourceTxHash,
		SrcTimestamp:    srcTimestamp,
		SrcTxOrigin:     srcTxOrigin,
		RouteAmount:     routeAmount,
		RouteRecipient:  routeRecipient,
		DestChainId:     destChainId,
		RequestSender:   requestSender,
		RequestMetadata: requestMetadata,
		RequestPayload:  requestPayload,
		RelayerFee:      relayerFee,
		Status:          "created",
	}
}

func NewCrosschainRequestFromMsg(
	msg *MsgCrosschainRequest,
) *CrosschainRequest {
	return &CrosschainRequest{
		SrcChainId:      msg.SrcchainId,
		SrcEventNonce:   msg.EventNonce,
		SrcBlockHeight:  msg.BlockHeight,
		SourceTxHash:    msg.SourceTxHash,
		SrcTimestamp:    msg.SrcTimestamp,
		SrcTxOrigin:     msg.SrcTxOrigin,
		RouteAmount:     msg.RouteAmount,
		RouteRecipient:  msg.RouteRecipient,
		DestChainId:     msg.DestChainId,
		RequestSender:   msg.RequestSender,
		RequestMetadata: msg.RequestMetadata,
		RequestPayload:  msg.RequestPayload,
		RelayerFee:      msg.RelayerFee,
		Status:          "created",
	}
}

type Metadata interface{}
