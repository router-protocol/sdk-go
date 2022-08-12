package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
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

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
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
