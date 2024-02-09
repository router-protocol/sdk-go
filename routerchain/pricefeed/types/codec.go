package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
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

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { //nolint:staticcheck
	legacy.RegisterAminoMsg(cdc, &MsgTokenPrices{}, "pricefeed/TokenPrices")
	legacy.RegisterAminoMsg(cdc, &MsgGasPrices{}, "pricefeed/GasPrices")
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
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterLegacyAminoCodec(amino)
	// RegisterCrypto(amino)
	amino.Seal()
}
