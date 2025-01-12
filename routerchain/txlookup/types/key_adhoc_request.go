package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// AdhocRequestKeyPrefix is the prefix to retrieve all adhoc requests
	AdhocRequestKeyPrefix = "AdhocRequest/value/"
	// AdhocRequestNonceKey is the key to retrieve the nonce of the latest adhoc request
	AdhocRequestNonceKey = "AdhocRequestNonce/value/"
)

// AdhocRequestKey returns the store key to retrieve a adhoc request from the index fields
func AdhocRequestKey(
	middlewareContractAddress string,
	chainIds []string,
	hashes []string,
	metadata [][]byte,
) []byte {
	var parts [][]byte

	// Add the middleware contract address to parts
	parts = append(parts, []byte(middlewareContractAddress))

	// Convert each chainId to bytes and add it to parts
	for _, chainId := range chainIds {
		parts = append(parts, []byte(chainId))
	}

	// Convert each hash to bytes and add it to parts
	for _, hash := range hashes {
		parts = append(parts, []byte(hash))
	}

	// Add metadata directly to parts
	parts = append(parts, metadata...)

	// Use AppendBytes to combine all parts into a single byte slice
	return util.AppendBytes(parts...)
}
