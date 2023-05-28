package types

import (
	proto "github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/crypto/tmhash"
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
