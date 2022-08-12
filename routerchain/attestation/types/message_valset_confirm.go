package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgValsetConfirm = "valset_confirm"

var _ sdk.Msg = &MsgValsetConfirm{}

func NewMsgValsetConfirm(orchestrator string, valsetNonce uint64, ethAddress string, signature string) *MsgValsetConfirm {
	return &MsgValsetConfirm{
		Orchestrator: orchestrator,
		ValsetNonce:  valsetNonce,
		EthAddress:   ethAddress,
		Signature:    signature,
	}
}

func (msg *MsgValsetConfirm) Route() string {
	return RouterKey
}

func (msg *MsgValsetConfirm) Type() string {
	return TypeMsgValsetConfirm
}

func (msg *MsgValsetConfirm) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgValsetConfirm) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgValsetConfirm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}
	return nil
}
