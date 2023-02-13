package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

const TypeMsgIncrementCrosstalkAckGas = "increment_crosstalk_ack_gas"

var _ sdk.Msg = &MsgIncrementCrosstalkAckGas{}

func NewMsgIncrementCrosstalkAckGas(feePayer string, chainType multichainTypes.ChainType, chainId string, eventNonce uint64, ackGasLimit uint64, ackGasPrice uint64) *MsgIncrementCrosstalkAckGas {
	return &MsgIncrementCrosstalkAckGas{
		FeePayer:    feePayer,
		ChainType:   chainType,
		ChainId:     chainId,
		EventNonce:  eventNonce,
		AckGasLimit: ackGasLimit,
		AckGasPrice: ackGasPrice,
	}
}

func (msg *MsgIncrementCrosstalkAckGas) Route() string {
	return RouterKey
}

func (msg *MsgIncrementCrosstalkAckGas) Type() string {
	return TypeMsgIncrementCrosstalkAckGas
}

func (msg *MsgIncrementCrosstalkAckGas) GetSigners() []sdk.AccAddress {
	feePayer, err := sdk.AccAddressFromBech32(msg.FeePayer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{feePayer}
}

func (msg *MsgIncrementCrosstalkAckGas) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIncrementCrosstalkAckGas) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FeePayer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid feePayer address (%s)", err)
	}
	return nil
}
