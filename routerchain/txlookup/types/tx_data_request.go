package types

import (
	"github.com/cometbft/cometbft/crypto/tmhash"
	proto "github.com/cosmos/gogoproto/proto"
)

func (msg *TxDataRequest) ClaimHash() ([]byte, error) {
	txDataRequestClaimHash := TxDataRequestClaimHash{
		ChainId:     msg.ChainId,
		TxHash:      msg.TxHash,
		TxData:      msg.TxData,
		BlockHeight: msg.BlockHeight,
	}

	out, err := proto.Marshal(&txDataRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}
