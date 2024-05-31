package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConfirmCrosschainAckRequest = "confirm_crosschain_ack_request"

var _ sdk.Msg = &MsgConfirmCrosschainAckRequest{}

func NewMsgConfirmCrosschainAckRequest(orchestrator string, ackSrcChainId string, ackRequestIdentifier uint64, claimHash []byte, ethSigner string, signature string) *MsgConfirmCrosschainAckRequest {
	return &MsgConfirmCrosschainAckRequest{
		Orchestrator:         orchestrator,
		AckSrcChainId:        ackSrcChainId,
		AckRequestIdentifier: ackRequestIdentifier,
		ClaimHash:            claimHash,
		EthSigner:            ethSigner,
		Signature:            signature,
	}
}

func (msg *MsgConfirmCrosschainAckRequest) Route() string {
	return RouterKey
}

func (msg *MsgConfirmCrosschainAckRequest) Type() string {
	return TypeMsgConfirmCrosschainAckRequest
}

func (msg *MsgConfirmCrosschainAckRequest) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgConfirmCrosschainAckRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConfirmCrosschainAckRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}
	return nil
}
