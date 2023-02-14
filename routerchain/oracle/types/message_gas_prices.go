package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGasPrices = "gas_prices"

var _ sdk.Msg = &MsgGasPrices{}

func NewMsgGasPrices(gasOracleProvider string, gasPrices []GasPriceState) *MsgGasPrices {
	return &MsgGasPrices{
		GasOracleProvider: gasOracleProvider,
		GasPrices:         gasPrices,
	}
}

func (msg *MsgGasPrices) Route() string {
	return RouterKey
}

func (msg *MsgGasPrices) Type() string {
	return TypeMsgGasPrices
}

func (msg *MsgGasPrices) GetSigners() []sdk.AccAddress {
	gasOracleProvider, err := sdk.AccAddressFromBech32(msg.GasOracleProvider)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{gasOracleProvider}
}

func (msg *MsgGasPrices) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGasPrices) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GasOracleProvider)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid gasOracleProvider address (%s)", err)
	}
	return nil
}
