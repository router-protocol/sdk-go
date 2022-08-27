package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBandTokenPricesData = "band_token_prices_data"

var (
	_ sdk.Msg = &MsgBandTokenPricesData{}

	// BandTokenPricesResultStoreKeyPrefix is a prefix for storing result
	BandTokenPricesResultStoreKeyPrefix = "band_token_prices_result"

	// LastBandTokenPricesIDKey is the key for the last request id
	LastBandTokenPricesIDKey = "band_token_prices_last_id"

	// BandTokenPricesClientIDKey is query request identifier
	BandTokenPricesClientIDKey = "band_token_prices_id"
)

// NewMsgBandTokenPricesData creates a new BandTokenPrices message
func NewMsgBandTokenPricesData(
	creator string,
	oracleScriptID OracleScriptID,
	sourceChannel string,
	calldata *BandTokenPricesCallData,
	askCount uint64,
	minCount uint64,
	feeLimit sdk.Coins,
	prepareGas uint64,
	executeGas uint64,
) *MsgBandTokenPricesData {
	return &MsgBandTokenPricesData{
		ClientID:       BandTokenPricesClientIDKey,
		Creator:        creator,
		OracleScriptID: uint64(oracleScriptID),
		SourceChannel:  sourceChannel,
		Calldata:       calldata,
		AskCount:       askCount,
		MinCount:       minCount,
		FeeLimit:       feeLimit,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
	}
}

// Route returns the message route
func (m *MsgBandTokenPricesData) Route() string {
	return RouterKey
}

// Type returns the message type
func (m *MsgBandTokenPricesData) Type() string {
	return TypeMsgBandTokenPricesData
}

// GetSigners returns the message signers
func (m *MsgBandTokenPricesData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes returns the signed bytes from the message
func (m *MsgBandTokenPricesData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic check the basic message validation
func (m *MsgBandTokenPricesData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if m.SourceChannel == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid source channel")
	}
	return nil
}

// BandTokenPricesResultStoreKey is a function to generate key for each result in store
func BandTokenPricesResultStoreKey(requestID OracleRequestID) []byte {
	return append(KeyPrefix(BandTokenPricesResultStoreKeyPrefix), int64ToBytes(int64(requestID))...)
}
