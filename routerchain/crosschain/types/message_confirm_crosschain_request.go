package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConfirmCrosschainRequest = "confirm_crosschain_request"

var _ sdk.Msg = &MsgConfirmCrosschainRequest{}

func NewMsgConfirmCrosschainRequest(orchestrator string, sourceChainId string, requestIdentifier uint64, claimHash []byte, ethSigner string, signature string) *MsgConfirmCrosschainRequest {
	return &MsgConfirmCrosschainRequest{
		Orchestrator:      orchestrator,
		SourceChainId:     sourceChainId,
		RequestIdentifier: requestIdentifier,
		ClaimHash:         claimHash,
		EthSigner:         ethSigner,
		Signature:         signature,
	}
}

func (msg *MsgConfirmCrosschainRequest) Route() string {
	return RouterKey
}

func (msg *MsgConfirmCrosschainRequest) Type() string {
	return TypeMsgConfirmCrosschainRequest
}

func (msg *MsgConfirmCrosschainRequest) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgConfirmCrosschainRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConfirmCrosschainRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}
	return nil
}
