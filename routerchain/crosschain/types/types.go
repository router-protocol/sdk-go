package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewCrosschainRequest(
	srcChainId string,
	srcEventNonce uint64,
	srcBlockHeight uint64,
	sourceTxHash string,
	srcTimestamp uint64,
	srcTxOrigin string,
	routeAmount sdk.Int,
	routeRecipient string,
	destChainId string,
	requestSender string,
	requestMetadata []byte,
	requestPayload []byte,
) *CrosschainRequest {
	return &CrosschainRequest{
		SrcChainId:      srcChainId,
		EventNonce:      srcEventNonce,
		BlockHeight:     srcBlockHeight,
		SourceTxHash:    sourceTxHash,
		SrcTimestamp:    srcTimestamp,
		SrcTxOrigin:     srcTxOrigin,
		RouteAmount:     routeAmount,
		RouteRecipient:  routeRecipient,
		DestChainId:     destChainId,
		RequestSender:   requestSender,
		RequestMetadata: requestMetadata,
		RequestPayload:  requestPayload,
		Status:          "created",
	}
}

func NewCrosschainRequestFromMsg(
	msg *MsgCrosschainRequest,
) *CrosschainRequest {
	return &CrosschainRequest{
		SrcChainId:      msg.SrcChainId,
		EventNonce:      msg.EventNonce,
		BlockHeight:     msg.BlockHeight,
		SourceTxHash:    msg.SourceTxHash,
		SrcTimestamp:    msg.SrcTimestamp,
		SrcTxOrigin:     msg.SrcTxOrigin,
		RouteAmount:     msg.RouteAmount,
		RouteRecipient:  msg.RouteRecipient,
		DestChainId:     msg.DestChainId,
		RequestSender:   msg.RequestSender,
		RequestMetadata: msg.RequestMetadata,
		RequestPayload:  msg.RequestPayload,
		Status:          "created",
	}
}

type Metadata interface{}
