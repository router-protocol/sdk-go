package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCrossTalkRequest{}, "crosstalk/CrossTalkRequest", nil)

	cdc.RegisterConcrete(&MsgCrossTalkAckRequest{}, "crosstalk/CrossTalkAckRequest", nil)
	cdc.RegisterConcrete(&MsgCrossTalkAckReceipt{}, "crosstalk/CrossTalkAckReceipt", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCrossTalkRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCrossTalkAckRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCrossTalkAckReceipt{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
