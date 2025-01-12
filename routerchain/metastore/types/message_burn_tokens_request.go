package types

import (
	errorsmod "cosmossdk.io/errors"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types "github.com/router-protocol/sdk-go/routerchain/types"
)

const TypeMsgBurnTokensRequest = "burn_tokens_request"

var _ sdk.Msg = &MsgBurnTokensRequest{}

func NewMsgBurnTokensRequest(sender string, amount sdkmath.Int) *MsgBurnTokensRequest {
	return &MsgBurnTokensRequest{
		Sender: sender,
		Amount: types.NewRouterCoin(amount),
	}
}

func (msg *MsgBurnTokensRequest) Route() string {
	return RouterKey
}

func (msg *MsgBurnTokensRequest) Type() string {
	return TypeMsgBurnTokensRequest
}

func (msg *MsgBurnTokensRequest) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgBurnTokensRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnTokensRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	if !msg.Amount.IsPositive() {
		return errorsmod.Wrapf(ErrInvalidAmount, "negative amount (%s)", msg.Amount)
	}
	return nil
}
