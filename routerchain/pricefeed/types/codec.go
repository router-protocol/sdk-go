package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgTokenPrices{}, "pricefeed/TokenPrices", nil)
	cdc.RegisterConcrete(&MsgGasPrices{}, "pricefeed/GasPrices", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTokenPrices{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGasPrices{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&UpdateSymbolRequestProposal{},
		&UpdatePriceFeederInfoProposal{},
		&CreateWhitelistedIBCChannelProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
