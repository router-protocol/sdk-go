package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	attestationtypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAdhocRequest{}, "txlookup/CreateAdhocRequest", nil)
	cdc.RegisterConcrete(&MsgCreateTxDataRequest{}, "txlookup/CreateTxDataRequest", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { //nolint:staticcheck
	legacy.RegisterAminoMsg(cdc, &MsgCreateAdhocRequest{}, "txlookup/CreateAdhocRequest")
	legacy.RegisterAminoMsg(cdc, &MsgCreateTxDataRequest{}, "txlookup/CreateTxDataRequest")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAdhocRequest{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateTxDataRequest{},
	)

	registry.RegisterImplementations(
		(*attestationtypes.Claim)(nil),
		&MsgCreateTxDataRequest{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterLegacyAminoCodec(Amino)
	// RegisterCrypto(amino)
	Amino.Seal()
}
