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
	cdc.RegisterConcrete(&MsgSetCrosstalkFeePayer{}, "crosstalk/SetCrosstalkFeePayer", nil)
	cdc.RegisterConcrete(&MsgIncrementCrosstalkDestGas{}, "crosstalk/IncrementCrosstalkDestGas", nil)
	cdc.RegisterConcrete(&MsgIncrementCrosstalkAckGas{}, "crosstalk/IncrementCrosstalkAckGas", nil)
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetCrosstalkFeePayer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIncrementCrosstalkDestGas{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIncrementCrosstalkAckGas{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
