package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInstantiateCwContract2 = "instantiate_cw_contract2"

var _ sdk.Msg = &MsgInstantiateCwContract2{}

func NewMsgInstantiateCwContract2(sender string, admin string, codeId uint64, label string, msg string, funds string, salt string, fixMsg bool) *MsgInstantiateCwContract2 {
	return &MsgInstantiateCwContract2{
		Sender: sender,
		Admin:  admin,
		CodeId: codeId,
		Label:  label,
		Msg:    msg,
		Funds:  funds,
		Salt:   salt,
	}
}

func (msg *MsgInstantiateCwContract2) Route() string {
	return RouterKey
}

func (msg *MsgInstantiateCwContract2) Type() string {
	return TypeMsgInstantiateCwContract2
}

func (msg *MsgInstantiateCwContract2) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgInstantiateCwContract2) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInstantiateCwContract2) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
