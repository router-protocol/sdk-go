package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/pricefeed module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")

	ErrSymbolRequestNotFound   = sdkerrors.Register(ModuleName, 2, "symbol request not found")
	ErrPriceNotFound           = sdkerrors.Register(ModuleName, 3, "price not found")
	ErrResolveStatusNotSuccess = sdkerrors.Register(ModuleName, 4, "request is not resolved successfully")
	ErrEmptySymbolRequests     = sdkerrors.Register(ModuleName, 5, "submitted symbol requests are empty")
	ErrEmptySymbol             = sdkerrors.Register(ModuleName, 6, "symbol is empty")
	ErrInvalidOracleScriptID   = sdkerrors.Register(ModuleName, 7, "invalid oracle script id")

	ErrInvalidPricefeederName    = sdkerrors.Register(ModuleName, 8, "pricefeeder name is nil")
	ErrInvalidPricefeederAddress = sdkerrors.Register(ModuleName, 9, "invalid pricefeeder address")
)
