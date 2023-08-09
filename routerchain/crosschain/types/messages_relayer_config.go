package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateRelayerConfig = "create_relayer_config"
	TypeMsgUpdateRelayerConfig = "update_relayer_config"
	TypeMsgDeleteRelayerConfig = "delete_relayer_config"
)

var _ sdk.Msg = &MsgCreateRelayerConfig{}

func NewMsgCreateRelayerConfig(
	creator string,
	chainId string,
	relayerName string,
	channel string,
	relayerEnabled bool,

) *MsgCreateRelayerConfig {
	return &MsgCreateRelayerConfig{
		Creator:        creator,
		ChainId:        chainId,
		RelayerName:    relayerName,
		Channel:        channel,
		RelayerEnabled: relayerEnabled,
	}
}

func (msg *MsgCreateRelayerConfig) Route() string {
	return RouterKey
}

func (msg *MsgCreateRelayerConfig) Type() string {
	return TypeMsgCreateRelayerConfig
}

func (msg *MsgCreateRelayerConfig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRelayerConfig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRelayerConfig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRelayerConfig{}

func NewMsgUpdateRelayerConfig(
	creator string,
	chainId string,
	relayerName string,
	channel string,
	relayerEnabled bool,

) *MsgUpdateRelayerConfig {
	return &MsgUpdateRelayerConfig{
		Creator:        creator,
		ChainId:        chainId,
		RelayerName:    relayerName,
		Channel:        channel,
		RelayerEnabled: relayerEnabled,
	}
}

func (msg *MsgUpdateRelayerConfig) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRelayerConfig) Type() string {
	return TypeMsgUpdateRelayerConfig
}

func (msg *MsgUpdateRelayerConfig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRelayerConfig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRelayerConfig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRelayerConfig{}

func NewMsgDeleteRelayerConfig(
	creator string,
	chainId string,

) *MsgDeleteRelayerConfig {
	return &MsgDeleteRelayerConfig{
		Creator: creator,
		ChainId: chainId,
	}
}
func (msg *MsgDeleteRelayerConfig) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRelayerConfig) Type() string {
	return TypeMsgDeleteRelayerConfig
}

func (msg *MsgDeleteRelayerConfig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRelayerConfig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRelayerConfig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
