package types

import (
	"github.com/bandprotocol/bandchain-packet/obi"
)

const BandPriceMultiplier uint64 = 1000000000 // 1e9

// GetCalldata gets the Band IBC request call data based on the symbols and multiplier.
func (r *BandOracleRequest) GetCalldata() []byte {

	calldata := &BandTokenPricesCallData{
		Symbols:    r.Symbols,
		Multiplier: BandPriceMultiplier,
	}

	callData := obi.MustEncode(calldata)

	return callData
}
