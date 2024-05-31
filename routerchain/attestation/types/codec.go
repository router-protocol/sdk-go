package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetOrchestratorAddress{}, "attestation/SetOrchestratorAddress", nil)
	cdc.RegisterConcrete(&MsgValsetConfirm{}, "attestation/ValsetConfirm", nil)
	cdc.RegisterConcrete(&MsgValsetUpdatedClaim{}, "attestation/ValsetUpdatedClaim", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { //nolint:staticcheck
	legacy.RegisterAminoMsg(cdc, &MsgSetOrchestratorAddress{}, "attestation/SetOrchestratorAddress")
	legacy.RegisterAminoMsg(cdc, &MsgValsetConfirm{}, "attestation/ValsetConfirm")
	legacy.RegisterAminoMsg(cdc, &MsgValsetUpdatedClaim{}, "attestation/ValsetUpdatedClaim")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*Claim)(nil),
		&MsgValsetUpdatedClaim{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetOrchestratorAddress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgValsetConfirm{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgValsetUpdatedClaim{},
	)

	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterLegacyAminoCodec(Amino)
	// RegisterCrypto(amino)
	Amino.Seal()
}
