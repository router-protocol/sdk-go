package types

import (
	"github.com/cometbft/cometbft/crypto/tmhash"
	proto "github.com/cosmos/gogoproto/proto"
)

func (msg *MetadataRequest) ClaimHash() ([]byte, error) {
	metadataRequestClaimHash := NewMetadataRequestClaimHash(
		msg.ChainId,
		msg.Contract,
		msg.EventNonce,
		msg.BlockHeight,
		msg.DaapAddress,
		msg.FeePayer,
	)

	out, err := proto.Marshal(metadataRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}
