package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTokenPrices = "token_prices"

var _ sdk.Msg = &MsgTokenPrices{}

func NewMsgTokenPrices(tokenOracleProvider string, tokenPrice []TokenPrice) *MsgTokenPrices {
	return &MsgTokenPrices{
		TokenOracleProvider: tokenOracleProvider,
		TokenPrices:         tokenPrice,
	}
}

func (msg *MsgTokenPrices) Route() string {
	return RouterKey
}

func (msg *MsgTokenPrices) Type() string {
	return TypeMsgTokenPrices
}

func (msg *MsgTokenPrices) GetSigners() []sdk.AccAddress {
	tokenOracleProvider, err := sdk.AccAddressFromBech32(msg.TokenOracleProvider)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{tokenOracleProvider}
}

func (msg *MsgTokenPrices) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTokenPrices) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.TokenOracleProvider)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid tokenOracleProvider address (%s)", err)
	}
	return nil
}
