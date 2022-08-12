package util

import (
	"crypto/md5"
	"encoding/binary"
	fmt "fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// UInt64FromBytes create uint from binary big endian representation
func UInt64FromBytes(s []byte) uint64 {
	return binary.BigEndian.Uint64(s)
}

// UInt64Bytes uses the SDK byte marshaling to encode a uint64
func UInt64Bytes(n uint64) []byte {
	return sdk.Uint64ToBigEndian(n)
}

// UInt64FromString to parse out a uint64 for a nonce
func UInt64FromString(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

// GetPrefixFromBech32 returns the human readable part of a bech32 string (excluding the 1 byte)
// Returns an error on too short input or when the 1 byte cannot be found
// Note: This is an excerpt from the Decode function for bech32 strings
func GetPrefixFromBech32(bech32str string) (string, error) {
	if len(bech32str) < 8 {
		return "", fmt.Errorf("invalid bech32 string length %d",
			len(bech32str))
	}
	one := strings.LastIndexByte(bech32str, '1')
	if one < 1 || one+7 > len(bech32str) {
		return "", fmt.Errorf("invalid index of 1")
	}

	return bech32str[:one], nil
}

// Hashing string using cryptographic MD5 function
// returns 128bit(16byte) value
func HashString(input string) []byte {
	md5 := md5.New()
	md5.Write([]byte(input))
	return md5.Sum(nil)
}

func AppendBytes(args ...[]byte) []byte {
	length := 0
	for _, v := range args {
		length += len(v)
	}

	res := make([]byte, length)

	length = 0
	for _, v := range args {
		copy(res[length:length+len(v)], v)
		length += len(v)
	}

	return res
}

func StrToFixByteArray(s string) ([32]byte, error) {
	var out [32]byte
	if len([]byte(s)) > 32 {
		return out, fmt.Errorf("string too long")
	}
	copy(out[:], s)
	return out, nil
}
