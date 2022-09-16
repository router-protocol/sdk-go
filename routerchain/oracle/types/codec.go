package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBandTokenPricesData{}, "oracle/BandTokenPricesData", nil)
	cdc.RegisterConcrete(&MsgCreateBandOracleRequest{}, "oracle/CreateBandOracleRequest", nil)
	cdc.RegisterConcrete(&MsgUpdateBandOracleRequest{}, "oracle/UpdateBandOracleRequest", nil)
	cdc.RegisterConcrete(&MsgDeleteBandOracleRequest{}, "oracle/DeleteBandOracleRequest", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBandTokenPricesData{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBandOracleRequest{},
		&MsgUpdateBandOracleRequest{},
		&MsgDeleteBandOracleRequest{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
