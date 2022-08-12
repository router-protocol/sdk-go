package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInboundRequest = "inbound_request"

var _ sdk.Msg = &MsgInboundRequest{}

func NewMsgInboundRequest(orchestrator string, eventNonce uint64, chainId uint64, blockHeight uint64, sender string, routerBridgeContract string, payload string) *MsgInboundRequest {
	return &MsgInboundRequest{
		Orchestrator:         orchestrator,
		EventNonce:           eventNonce,
		ChainId:              chainId,
		BlockHeight:          blockHeight,
		Sender:               sender,
		RouterBridgeContract: routerBridgeContract,
		Payload:              payload,
	}
}

func (msg *MsgInboundRequest) Route() string {
	return RouterKey
}

func (msg *MsgInboundRequest) Type() string {
	return TypeMsgInboundRequest
}

func (msg *MsgInboundRequest) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgInboundRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInboundRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}
	return nil
}
