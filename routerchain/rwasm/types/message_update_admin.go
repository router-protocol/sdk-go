package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateAdmin = "update_admin"

var _ sdk.Msg = &MsgUpdateAdmin{}

func NewMsgUpdateAdmin(sender string, contract string, admin string) *MsgUpdateAdmin {
	return &MsgUpdateAdmin{
		Sender:   sender,
		Contract: contract,
		NewAdmin: admin,
	}
}

func (msg *MsgUpdateAdmin) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAdmin) Type() string {
	return TypeMsgUpdateAdmin
}

func (msg *MsgUpdateAdmin) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgUpdateAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
