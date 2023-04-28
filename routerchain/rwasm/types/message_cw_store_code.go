package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCwStoreCode = "cw_store_code"

var _ sdk.Msg = &MsgCwStoreCode{}

func NewMsgCwStoreCode(sender string, wasmByteCode string) *MsgCwStoreCode {
	return &MsgCwStoreCode{
		Sender:       sender,
		WasmByteCode: wasmByteCode,
	}
}

func (msg *MsgCwStoreCode) Route() string {
	return RouterKey
}

func (msg *MsgCwStoreCode) Type() string {
	return TypeMsgCwStoreCode
}

func (msg *MsgCwStoreCode) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgCwStoreCode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCwStoreCode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
