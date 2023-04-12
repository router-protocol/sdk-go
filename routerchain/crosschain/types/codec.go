package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCrosschainRequest{}, "crosschain/CrosschainRequest", nil)
	cdc.RegisterConcrete(&MsgConfirmCrosschainRequest{}, "crosschain/ConfirmCrosschainRequest", nil)
	cdc.RegisterConcrete(&MsgCrosschainAckRequest{}, "crosschain/CrosschainAckRequest", nil)
	cdc.RegisterConcrete(&MsgConfirmCrosschainAckRequest{}, "crosschain/ConfirmCrosschainAckRequest", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCrosschainRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConfirmCrosschainRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCrosschainAckRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConfirmCrosschainAckRequest{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
