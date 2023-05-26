package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTokenPrices = "token_prices"

var _ sdk.Msg = &MsgTokenPrices{}

func NewMsgTokenPrices(priceFeederAddress string, priceFeederName string, tokenPrices []Price) *MsgTokenPrices {
	return &MsgTokenPrices{
		PriceFeederAddress: priceFeederAddress,
		PriceFeederName:    priceFeederName,
		TokenPrices:        tokenPrices,
	}
}

func (msg *MsgTokenPrices) Route() string {
	return RouterKey
}

func (msg *MsgTokenPrices) Type() string {
	return TypeMsgTokenPrices
}

func (msg *MsgTokenPrices) GetSigners() []sdk.AccAddress {
	priceFeeder, err := sdk.AccAddressFromBech32(msg.PriceFeederAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{priceFeeder}
}

func (msg *MsgTokenPrices) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTokenPrices) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.PriceFeederAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid priceFeeder address (%s)", err)
	}
	return nil
}
