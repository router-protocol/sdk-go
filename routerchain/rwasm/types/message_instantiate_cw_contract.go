package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInstantiateCwContract = "instantiate_cw_contract"

var _ sdk.Msg = &MsgInstantiateCwContract{}

func NewMsgInstantiateCwContract(sender string, admin string, codeId uint64, label string, msg string, funds string) *MsgInstantiateCwContract {
	return &MsgInstantiateCwContract{
		Sender: sender,
		Admin:  admin,
		CodeId: codeId,
		Label:  label,
		Msg:    msg,
		Funds:  funds,
	}
}

func (msg *MsgInstantiateCwContract) Route() string {
	return RouterKey
}

func (msg *MsgInstantiateCwContract) Type() string {
	return TypeMsgInstantiateCwContract
}

func (msg *MsgInstantiateCwContract) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgInstantiateCwContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInstantiateCwContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
