package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClearAdmin = "clear_admin"

var _ sdk.Msg = &MsgClearAdmin{}

func NewMsgClearAdmin(sender string, contract string) *MsgClearAdmin {
	return &MsgClearAdmin{
		Sender:   sender,
		Contract: contract,
	}
}

func (msg *MsgClearAdmin) Route() string {
	return RouterKey
}

func (msg *MsgClearAdmin) Type() string {
	return TypeMsgClearAdmin
}

func (msg *MsgClearAdmin) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgClearAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClearAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
