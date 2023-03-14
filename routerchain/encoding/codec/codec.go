// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/router-protocol/router-chain/blob/main/LICENSE
package codec

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ethermint "github.com/evmos/ethermint/types"
	cryptocodec "github.com/router-protocol/sdk-go/routerchain/crypto/codec"

	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	crosstalkTypes "github.com/router-protocol/sdk-go/routerchain/crosstalk/types"
	inboundTypes "github.com/router-protocol/sdk-go/routerchain/inbound/types"
	outboundTypes "github.com/router-protocol/sdk-go/routerchain/outbound/types"
)

// RegisterLegacyAminoCodec registers Interfaces from types, crypto, and SDK std.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
}

// RegisterInterfaces registers Interfaces from types, crypto, and SDK std.
func RegisterInterfaces(interfaceRegistry codectypes.InterfaceRegistry) {

	std.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	fmt.Println("Register encoding config")
	ethermint.RegisterInterfaces(interfaceRegistry)

	interfaceRegistry.RegisterInterface(
		"attestation.Claim",
		(*attestationTypes.Claim)(nil),
		&inboundTypes.MsgInboundRequest{},
		&crosstalkTypes.MsgCrossTalkRequest{},
		&crosstalkTypes.MsgCrossTalkAckRequest{},
		&crosstalkTypes.MsgCrossTalkAckReceipt{},
		&attestationTypes.MsgValsetUpdatedClaim{},
		&outboundTypes.MsgOutboundAckRequest{},
	)
}
