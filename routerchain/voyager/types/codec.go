package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgFundsDeposited{}, "voyager/FundsDeposited", nil)
	cdc.RegisterConcrete(&MsgFundsPaid{}, "voyager/FundsPaid", nil)
	cdc.RegisterConcrete(&MsgDepositInfoUpdated{}, "voyager/DepositInfoUpdated", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFundsDeposited{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFundsPaid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDepositInfoUpdated{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
