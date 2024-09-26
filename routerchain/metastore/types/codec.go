package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	attestationtypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
)

var (
	Amino = codec.NewLegacyAmino()
	// ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())

	// AminoCdc is a amino codec created to support amino JSON compatible msgs.
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateMetadataRequest{}, "metastore/CreateMetadataRequest", nil)
	cdc.RegisterConcrete(&MsgApproveFeepayerRequest{}, "metastore/ApproveFeepayerRequest", nil)
	cdc.RegisterConcrete(&MsgRevokeFeepayerRequest{}, "metastore/RevokeFeepayerRequest", nil)
	cdc.RegisterConcrete(&MsgBurnTokensRequest{}, "metastore/BurnTokensRequest", nil)

	// this line is used by starport scaffolding # 2
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { //nolint:staticcheck
	legacy.RegisterAminoMsg(cdc, &MsgCreateMetadataRequest{}, "metastore/CreateMetadataRequest")
	legacy.RegisterAminoMsg(cdc, &MsgApproveFeepayerRequest{}, "metastore/ApproveFeepayerRequest")
	legacy.RegisterAminoMsg(cdc, &MsgRevokeFeepayerRequest{}, "metastore/RevokeFeepayerRequest")
	legacy.RegisterAminoMsg(cdc, &MsgBurnTokensRequest{}, "metastore/BurnTokensRequest")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*attestationtypes.Claim)(nil),
		&MsgCreateMetadataRequest{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMetadataRequest{},
		&MsgApproveFeepayerRequest{},
		&MsgRevokeFeepayerRequest{},
		&MsgBurnTokensRequest{},
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
