package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ValsetKeyPrefix is the prefix to retrieve all Valset
	ValsetKeyPrefix = "Valset/value/"
)

var (
	// LatestValsetNonce indexes the latest valset nonce
	// [0xba0fa05da166611b656bac7739a6e7d3]
	LatestValsetNonce = util.HashString("LatestValsetNonce")

	// LastUnBondingBlockHeight indexes the last validator unbonding block height
	// [0x06a6b30651341e80276e0d2e19449250]
	LastUnBondingBlockHeight = util.HashString("LastUnBondingBlockHeight")

	// LastEventNonceByValidatorKey indexes lateset event nonce by validator
	// [0xeefcb999cc3d7b80b052b55106a6ba5e]
	LastEventNonceByValidatorKey = util.HashString("LastEventNonceByValidatorKey")
)

// ValsetKey returns the store key to retrieve a Valset from the index fields
func ValsetKey(
	valsetNonce uint64,
) []byte {
	return util.AppendBytes(util.HashString(ValsetKeyPrefix), util.UInt64Bytes(valsetNonce))
}

// GetLastEventNonceByValidatorKey indexes latest event nonce by validator
// GetLastEventNonceByValidatorKey returns the following key format
// prefix              cosmos-validator
// [0x0][gravity1ahx7f8wyertuus9r20284ej0asrs085ceqtfnm]
func GetLastEventNonceByValidatorKey(validator sdk.ValAddress, chainType multichainTypes.ChainType, chainId string) []byte {
	if err := sdk.VerifyAddressFormat(validator); err != nil {
		panic(sdkerrors.Wrap(err, "invalid validator address"))
	}
	return util.AppendBytes(LastEventNonceByValidatorKey, util.UInt64Bytes(uint64(chainType)), []byte(chainId), validator.Bytes())
}
