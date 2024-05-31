package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	Amino = codec.NewLegacyAmino()
	// ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgExecuteCwContract{}, "rwasm/ExecuteCwContract", nil)
	cdc.RegisterConcrete(&MsgInstantiateCwContract{}, "rwasm/InstantiateCwContract", nil)
	cdc.RegisterConcrete(&MsgCwStoreCode{}, "rwasm/CwStoreCode", nil)
	cdc.RegisterConcrete(&MsgMigrateCwContract{}, "rwasm/MigrateCwContract", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { //nolint:staticcheck
	cdc.RegisterConcrete(&MsgExecuteCwContract{}, "rwasm/ExecuteCwContract", nil)
	cdc.RegisterConcrete(&MsgInstantiateCwContract{}, "rwasm/InstantiateCwContract", nil)
	cdc.RegisterConcrete(&MsgCwStoreCode{}, "rwasm/CwStoreCode", nil)
	cdc.RegisterConcrete(&MsgMigrateCwContract{}, "rwasm/MigrateCwContract", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExecuteCwContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInstantiateCwContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCwStoreCode{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMigrateCwContract{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterLegacyAminoCodec(Amino)
	// RegisterCrypto(amino)
	Amino.Seal()
}
