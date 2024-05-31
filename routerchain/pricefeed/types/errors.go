package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/pricefeed module sentinel errors
var (
	ErrSample               = errorsmod.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = errorsmod.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = errorsmod.Register(ModuleName, 1501, "invalid version")

	ErrSymbolRequestNotFound   = errorsmod.Register(ModuleName, 2, "symbol request not found")
	ErrPriceNotFound           = errorsmod.Register(ModuleName, 3, "price not found")
	ErrResolveStatusNotSuccess = errorsmod.Register(ModuleName, 4, "request is not resolved successfully")
	ErrEmptySymbolRequests     = errorsmod.Register(ModuleName, 5, "submitted symbol requests are empty")
	ErrEmptySymbol             = errorsmod.Register(ModuleName, 6, "symbol is empty")
	ErrInvalidOracleScriptID   = errorsmod.Register(ModuleName, 7, "invalid oracle script id")

	ErrInvalidPricefeederName    = errorsmod.Register(ModuleName, 8, "pricefeeder name is nil")
	ErrInvalidPricefeederAddress = errorsmod.Register(ModuleName, 9, "invalid pricefeeder address")
)
