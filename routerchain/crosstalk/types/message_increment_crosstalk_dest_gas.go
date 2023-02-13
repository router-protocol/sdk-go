package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

const TypeMsgIncrementCrosstalkDestGas = "increment_crosstalk_dest_gas"

var _ sdk.Msg = &MsgIncrementCrosstalkDestGas{}

func NewMsgIncrementCrosstalkDestGas(feePayer string, chainType multichainTypes.ChainType, chainId string, eventNonce uint64, destGasLimit uint64, destGasPrice uint64) *MsgIncrementCrosstalkDestGas {
	return &MsgIncrementCrosstalkDestGas{
		FeePayer:     feePayer,
		ChainType:    chainType,
		ChainId:      chainId,
		EventNonce:   eventNonce,
		DestGasLimit: destGasLimit,
		DestGasPrice: destGasPrice,
	}
}

func (msg *MsgIncrementCrosstalkDestGas) Route() string {
	return RouterKey
}

func (msg *MsgIncrementCrosstalkDestGas) Type() string {
	return TypeMsgIncrementCrosstalkDestGas
}

func (msg *MsgIncrementCrosstalkDestGas) GetSigners() []sdk.AccAddress {
	feePayer, err := sdk.AccAddressFromBech32(msg.FeePayer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{feePayer}
}

func (msg *MsgIncrementCrosstalkDestGas) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIncrementCrosstalkDestGas) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FeePayer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid feePayer address (%s)", err)
	}
	return nil
}
