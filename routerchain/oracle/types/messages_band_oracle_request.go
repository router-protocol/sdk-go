package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateBandOracleRequest = "create_band_oracle_request"
	TypeMsgUpdateBandOracleRequest = "update_band_oracle_request"
	TypeMsgDeleteBandOracleRequest = "delete_band_oracle_request"
)

var _ sdk.Msg = &MsgCreateBandOracleRequest{}

func NewMsgCreateBandOracleRequest(
	creator string,
	requestId uint64,
	oracleScriptId uint64,
	symbols []string,
	askCount uint64,
	minCount uint64,
	prepareGas uint64,
	executeGas uint64,

) *MsgCreateBandOracleRequest {
	return &MsgCreateBandOracleRequest{
		Creator:        creator,
		RequestId:      requestId,
		OracleScriptId: oracleScriptId,
		Symbols:        symbols,
		AskCount:       askCount,
		MinCount:       minCount,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
	}
}

func (msg *MsgCreateBandOracleRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateBandOracleRequest) Type() string {
	return TypeMsgCreateBandOracleRequest
}

func (msg *MsgCreateBandOracleRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBandOracleRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBandOracleRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBandOracleRequest{}

func NewMsgUpdateBandOracleRequest(
	creator string,
	requestId uint64,
	oracleScriptId uint64,
	symbols []string,
	askCount uint64,
	minCount uint64,
	prepareGas uint64,
	executeGas uint64,

) *MsgUpdateBandOracleRequest {
	return &MsgUpdateBandOracleRequest{
		Creator:        creator,
		RequestId:      requestId,
		OracleScriptId: oracleScriptId,
		Symbols:        symbols,
		AskCount:       askCount,
		MinCount:       minCount,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
	}
}

func (msg *MsgUpdateBandOracleRequest) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBandOracleRequest) Type() string {
	return TypeMsgUpdateBandOracleRequest
}

func (msg *MsgUpdateBandOracleRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBandOracleRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBandOracleRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteBandOracleRequest{}

func NewMsgDeleteBandOracleRequest(
	creator string,
	requestId uint64,

) *MsgDeleteBandOracleRequest {
	return &MsgDeleteBandOracleRequest{
		Creator:   creator,
		RequestId: requestId,
	}
}
func (msg *MsgDeleteBandOracleRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBandOracleRequest) Type() string {
	return TypeMsgDeleteBandOracleRequest
}

func (msg *MsgDeleteBandOracleRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBandOracleRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBandOracleRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
