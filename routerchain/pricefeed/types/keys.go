package types

const (
	// ModuleName defines the module name
	ModuleName = "pricefeed"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_pricefeed"

	// Version defines the current version the IBC module supports
	Version = "bandchain-1"

	// PortID is the default port id that module binds to
	PortID = "pricefeed"

	ROUTER_PRICE_FEEDER = "router_price_feeder"
	BAND_PRICE_FEEDER   = "band_feeder"
)

// PortKey defines the key to store the port ID in store
var PortKey = KeyPrefix("pricefeed-port-")

func KeyPrefix(p string) []byte {
	return []byte(p)
}
