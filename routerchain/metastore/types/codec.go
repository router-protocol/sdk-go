package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino = codec.NewLegacyAmino()
	// ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())

	// AminoCdc is a amino codec created to support amino JSON compatible msgs.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateMetadataRequest{}, "metastore/CreateMetadataRequest", nil)
	cdc.RegisterConcrete(&MsgApproveFeepayerRequest{}, "metastore/ApproveFeepayerRequest", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { //nolint:staticcheck
	cdc.RegisterConcrete(&MsgCreateMetadataRequest{}, "metastore/CreateMetadataRequest", nil)
	cdc.RegisterConcrete(&MsgApproveFeepayerRequest{}, "metastore/ApproveFeepayerRequest", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMetadataRequest{},
		&MsgApproveFeepayerRequest{},
	)

	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterLegacyAminoCodec(amino)
	// RegisterCrypto(amino)
	amino.Seal()
}
