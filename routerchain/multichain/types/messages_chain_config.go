package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateChainConfig = "create_chain_config"
	TypeMsgUpdateChainConfig = "update_chain_config"
	TypeMsgDeleteChainConfig = "delete_chain_config"
)

var _ sdk.Msg = &MsgCreateChainConfig{}

func NewMsgCreateChainConfig(
	creator string,
	chainId string,
	chainName string,
	symbol string,
	chainType ChainType,
	confirmationsRequired uint64,
	gatewayContractAddress string,
	gatewayContractHeight uint64,
	RouterContractAddress string,
	lastObservedEventNonce uint64,
	lastObservedValsetNonce uint64,

) *MsgCreateChainConfig {
	return &MsgCreateChainConfig{
		Creator:                 creator,
		ChainId:                 chainId,
		ChainName:               chainName,
		Symbol:                  symbol,
		ChainType:               chainType,
		ConfirmationsRequired:   confirmationsRequired,
		GatewayContractAddress:  gatewayContractAddress,
		GatewayContractHeight:   gatewayContractHeight,
		RouterContractAddress:   RouterContractAddress,
		LastObservedEventNonce:  lastObservedEventNonce,
		LastObservedValsetNonce: lastObservedValsetNonce,
	}
}

func (msg *MsgCreateChainConfig) Route() string {
	return RouterKey
}

func (msg *MsgCreateChainConfig) Type() string {
	return TypeMsgCreateChainConfig
}

func (msg *MsgCreateChainConfig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateChainConfig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateChainConfig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateChainConfig{}

func NewMsgUpdateChainConfig(
	creator string,
	ChainId string,
	chainName string,
	symbol string,
	chainType ChainType,
	confirmationsRequired uint64,
	gatewayContractAddress string,
	gatewayContractHeight uint64,
	RouterContractAddress string,
	lastObservedEventNonce uint64,
	lastObservedValsetNonce uint64,

) *MsgUpdateChainConfig {
	return &MsgUpdateChainConfig{
		Creator:                 creator,
		ChainId:                 ChainId,
		ChainName:               chainName,
		Symbol:                  symbol,
		ChainType:               chainType,
		ConfirmationsRequired:   confirmationsRequired,
		GatewayContractAddress:  gatewayContractAddress,
		GatewayContractHeight:   gatewayContractHeight,
		RouterContractAddress:   RouterContractAddress,
		LastObservedEventNonce:  lastObservedEventNonce,
		LastObservedValsetNonce: lastObservedValsetNonce,
	}
}

func (msg *MsgUpdateChainConfig) Route() string {
	return RouterKey
}

func (msg *MsgUpdateChainConfig) Type() string {
	return TypeMsgUpdateChainConfig
}

func (msg *MsgUpdateChainConfig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateChainConfig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateChainConfig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteChainConfig{}

func NewMsgDeleteChainConfig(
	creator string,
	chainId string,
) *MsgDeleteChainConfig {
	return &MsgDeleteChainConfig{
		Creator: creator,
		ChainId: chainId,
	}
}
func (msg *MsgDeleteChainConfig) Route() string {
	return RouterKey
}

func (msg *MsgDeleteChainConfig) Type() string {
	return TypeMsgDeleteChainConfig
}

func (msg *MsgDeleteChainConfig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteChainConfig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteChainConfig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
