package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	multichaintypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

const TypeMsgOutboundBatchRequest = "outbound_batch_request"

var _ sdk.Msg = &MsgOutboundBatchRequest{}

func NewMsgOutboundBatchRequest(sender string, destinationChainType multichaintypes.ChainType, destinationChainId string, contractCalls []ContractCall, isAtomic bool, relayerFee sdk.Coin, outgoingTxFee sdk.Coin) *MsgOutboundBatchRequest {
	return &MsgOutboundBatchRequest{
		Sender:               sender,
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		ContractCalls:        contractCalls,
		IsAtomic:             isAtomic,
		RelayerFee:           relayerFee,
		OutgoingTxFee:        outgoingTxFee,
	}
}

func (msg *MsgOutboundBatchRequest) Route() string {
	return RouterKey
}

func (msg *MsgOutboundBatchRequest) Type() string {
	return TypeMsgOutboundBatchRequest
}

func (msg *MsgOutboundBatchRequest) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgOutboundBatchRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOutboundBatchRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
