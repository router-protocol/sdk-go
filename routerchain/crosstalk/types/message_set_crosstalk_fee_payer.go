package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

const TypeMsgSetCrosstalkFeePayer = "set_crosstalk_fee_payer"

var _ sdk.Msg = &MsgSetCrosstalkFeePayer{}

func NewMsgSetCrosstalkFeePayer(feePayer string, chainType multichainTypes.ChainType, chainId string, eventNonce uint64) *MsgSetCrosstalkFeePayer {
	return &MsgSetCrosstalkFeePayer{
		FeePayer:   feePayer,
		ChainType:  chainType,
		ChainId:    chainId,
		EventNonce: eventNonce,
	}
}

func (msg *MsgSetCrosstalkFeePayer) Route() string {
	return RouterKey
}

func (msg *MsgSetCrosstalkFeePayer) Type() string {
	return TypeMsgSetCrosstalkFeePayer
}

func (msg *MsgSetCrosstalkFeePayer) GetSigners() []sdk.AccAddress {
	feePayer, err := sdk.AccAddressFromBech32(msg.FeePayer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{feePayer}
}

func (msg *MsgSetCrosstalkFeePayer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetCrosstalkFeePayer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FeePayer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid feePayer address (%s)", err)
	}
	return nil
}
