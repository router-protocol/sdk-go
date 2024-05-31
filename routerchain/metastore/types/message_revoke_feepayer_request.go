package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevokeFeepayerRequest = "revoke_feepayer_request"

var _ sdk.Msg = &MsgRevokeFeepayerRequest{}

func NewMsgRevokeFeepayerRequest(feepayer string, chainId string, dappAddress string) *MsgRevokeFeepayerRequest {
	return &MsgRevokeFeepayerRequest{
		Feepayer:    feepayer,
		ChainId:     chainId,
		DappAddress: dappAddress,
	}
}

func (msg *MsgRevokeFeepayerRequest) Route() string {
	return RouterKey
}

func (msg *MsgRevokeFeepayerRequest) Type() string {
	return TypeMsgRevokeFeepayerRequest
}

func (msg *MsgRevokeFeepayerRequest) GetSigners() []sdk.AccAddress {
	feepayer, err := sdk.AccAddressFromBech32(msg.Feepayer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{feepayer}
}

func (msg *MsgRevokeFeepayerRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeFeepayerRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Feepayer)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid feepayer address (%s)", err)
	}
	return nil
}
