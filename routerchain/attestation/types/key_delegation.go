package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var (
	// EthAddressByValidatorKey indexes cosmos validator account addresses
	// i.e. gravity1ahx7f8wyertuus9r20284ej0asrs085ceqtfnm
	// [0x1248a4405201cc3a00ab515ce9c4dd47]
	EthAddressByValidatorKey = util.HashString("EthAddressValidatorKey")

	// ValidatorByEthAddressKey indexes ethereum addresses
	// i.e. 0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B
	// [0xbfe41763f372108317ed982a4cd1b4a8]
	ValidatorByEthAddressKey = util.HashString("ValidatorByEthAddressKey")

	// KeyOrchestratorAddress indexes the validator keys for an orchestrator
	// [0x391e8708521fb085676169e8fb232cda]
	KeyOrchestratorAddress = util.HashString("KeyOrchestratorAddress")
)

// GetOrchestratorAddressKey returns the following key format
// prefix 				orchestrator address
// [0x0][gravity1ahx7f8wyertuus9r20284ej0asrs085ceqtfnm]
func GetOrchestratorAddressKey(orc sdk.AccAddress) []byte {
	if err := sdk.VerifyAddressFormat(orc); err != nil {
		panic(sdkerrors.Wrap(err, "invalid orchestrator address"))
	}
	return util.AppendBytes(KeyOrchestratorAddress, orc.Bytes())
}

// GetEthAddressByValidatorKey returns the following key format
// prefix              cosmos-validator
// [0x0][gravityvaloper1ahx7f8wyertuus9r20284ej0asrs085ceqtfnm]
func GetEthAddressByValidatorKey(validator sdk.ValAddress) []byte {
	if err := sdk.VerifyAddressFormat(validator); err != nil {
		panic(sdkerrors.Wrap(err, "invalid validator address"))
	}
	return util.AppendBytes(EthAddressByValidatorKey, validator.Bytes())
}

// GetValidatorByEthAddressKey returns the following key format
// prefix              cosmos-validator
// [0x0][0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B]
func GetValidatorByEthAddressKey(ethAddress EthAddress) []byte {
	return util.AppendBytes(ValidatorByEthAddressKey, ethAddress.GetAddress().Bytes())
}
