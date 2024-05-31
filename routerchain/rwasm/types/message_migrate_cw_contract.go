package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMigrateCwContract = "migrate_cw_contract"

var _ sdk.Msg = &MsgMigrateCwContract{}

func NewMsgMigrateCwContract(sender string, contract string, codeId uint64, msg string) *MsgMigrateCwContract {
	return &MsgMigrateCwContract{
		Sender:   sender,
		Contract: contract,
		CodeId:   codeId,
		Msg:      msg,
	}
}

func (msg *MsgMigrateCwContract) Route() string {
	return RouterKey
}

func (msg *MsgMigrateCwContract) Type() string {
	return TypeMsgMigrateCwContract
}

func (msg *MsgMigrateCwContract) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgMigrateCwContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMigrateCwContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
