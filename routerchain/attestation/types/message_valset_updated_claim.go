package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgValsetUpdatedClaim = "valset_updated_claim"

var _ sdk.Msg = &MsgValsetUpdatedClaim{}

func NewMsgValsetUpdatedClaim(orchestrator string, eventNonce uint64, valsetNonce uint64, blockHeight uint64, members []*BridgeValidator) *MsgValsetUpdatedClaim {
	return &MsgValsetUpdatedClaim{
		Orchestrator: orchestrator,
		EventNonce:   eventNonce,
		ValsetNonce:  valsetNonce,
		BlockHeight:  blockHeight,
		Members:      members,
	}
}

func (msg *MsgValsetUpdatedClaim) Route() string {
	return RouterKey
}

func (msg *MsgValsetUpdatedClaim) Type() string {
	return TypeMsgValsetUpdatedClaim
}

func (msg *MsgValsetUpdatedClaim) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgValsetUpdatedClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgValsetUpdatedClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}
	return nil
}
