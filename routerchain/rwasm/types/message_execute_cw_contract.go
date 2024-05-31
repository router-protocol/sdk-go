package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgExecuteCwContract = "execute_cw_contract"

var _ sdk.Msg = &MsgExecuteCwContract{}

func NewMsgExecuteCwContract(sender string, contract string, msg string, funds string) *MsgExecuteCwContract {
	return &MsgExecuteCwContract{
		Sender:   sender,
		Contract: contract,
		Msg:      msg,
		Funds:    funds,
	}
}

func (msg *MsgExecuteCwContract) Route() string {
	return RouterKey
}

func (msg *MsgExecuteCwContract) Type() string {
	return TypeMsgExecuteCwContract
}

func (msg *MsgExecuteCwContract) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgExecuteCwContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExecuteCwContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
