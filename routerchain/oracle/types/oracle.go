package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	// OracleScriptID is the type-safe unique identifier type for oracle scripts.
	OracleScriptID uint64

	// OracleRequestID is the type-safe unique identifier type for data requests.
	OracleRequestID int64
)

// int64ToBytes convert int64 to a byte slice
func int64ToBytes(num int64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, uint64(num))
	return result
}

func NewPriceState(price sdk.Dec, timestamp int64) *TokenPriceState {
	return &TokenPriceState{
		Price:           price,
		CumulativePrice: sdk.ZeroDec(),
		Timestamp:       timestamp,
	}
}

func (p *TokenPriceState) UpdatePrice(price sdk.Dec, timestamp int64) {
	cumulativePriceDelta := sdk.NewDec(timestamp - p.Timestamp).Mul(p.Price)
	p.CumulativePrice = p.CumulativePrice.Add(cumulativePriceDelta)
	p.Timestamp = timestamp
	p.Price = price
}
