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
	cdc.RegisterConcrete(&MsgFundsDeposited{}, "voyager/FundsDeposited", nil)
	cdc.RegisterConcrete(&MsgFundsPaid{}, "voyager/FundsPaid", nil)
	cdc.RegisterConcrete(&MsgDepositInfoUpdated{}, "voyager/DepositInfoUpdated", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { //nolint:staticcheck
	legacy.RegisterAminoMsg(cdc, &MsgFundsDeposited{}, "voyager/FundsDeposited")
	legacy.RegisterAminoMsg(cdc, &MsgFundsPaid{}, "voyager/FundsPaid")
	legacy.RegisterAminoMsg(cdc, &MsgDepositInfoUpdated{}, "voyager/DepositInfoUpdated")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*attestationtypes.Claim)(nil),
		&MsgFundsDeposited{}, &MsgFundsPaid{}, &MsgDepositInfoUpdated{},
	)

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

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterLegacyAminoCodec(Amino)
	// RegisterCrypto(amino)
	Amino.Seal()
}
