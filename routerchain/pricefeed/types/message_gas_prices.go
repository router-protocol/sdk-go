package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGasPrices = "gas_prices"

var _ sdk.Msg = &MsgGasPrices{}

func NewMsgGasPrices(priceFeederAddress string, priceFeederName string, gasPrices []GasPrice) *MsgGasPrices {
	return &MsgGasPrices{
		PriceFeederAddress: priceFeederAddress,
		PriceFeederName:    priceFeederAddress,
		GasPrices:          gasPrices,
	}
}

func (msg *MsgGasPrices) Route() string {
	return RouterKey
}

func (msg *MsgGasPrices) Type() string {
	return TypeMsgGasPrices
}

func (msg *MsgGasPrices) GetSigners() []sdk.AccAddress {
	priceFeeder, err := sdk.AccAddressFromBech32(msg.PriceFeederAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{priceFeeder}
}

func (msg *MsgGasPrices) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGasPrices) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.PriceFeederAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid priceFeeder address (%s)", err)
	}
	return nil
}
