package types

import (
	"bytes"
	"encoding/hex"
	"fmt"

	gethcommon "github.com/ethereum/go-ethereum/common"

	errorsmod "cosmossdk.io/errors"
)

const (
	// ETHContractAddressLen is the length of contract address strings
	ETHContractAddressLen = 42

	// ZeroAddress is an EthAddress containing the zero ethereum address
	ZeroAddressString = "0x0000000000000000000000000000000000000000"
)

// Regular EthAddress
type EthAddress struct {
	address gethcommon.Address
}

// Returns the contained address as a string
func (ea EthAddress) GetAddress() gethcommon.Address {
	return ea.address
}

// Sets the contained address, performing validation before updating the value
func (ea *EthAddress) SetAddress(address string) error {
	if err := ValidateEthAddress(address); err != nil {
		return err
	}
	ea.address = gethcommon.HexToAddress(address)
	return nil
}

func NewEthAddressFromBytes(address []byte) (*EthAddress, error) {
	if err := ValidateEthAddress(hex.EncodeToString(address)); err != nil {
		return nil, errorsmod.Wrap(err, "invalid input address")
	}

	addr := EthAddress{gethcommon.BytesToAddress(address)}
	return &addr, nil
}

// Creates a new EthAddress from a string, performing validation and returning any validation errors
func NewEthAddress(address string) (*EthAddress, error) {
	if err := ValidateEthAddress(address); err != nil {
		return nil, errorsmod.Wrap(err, "invalid input address")
	}

	addr := EthAddress{gethcommon.HexToAddress(address)}
	return &addr, nil
}

// Returns a new EthAddress with 0x0000000000000000000000000000000000000000 as the wrapped address
func ZeroAddress() EthAddress {
	return EthAddress{gethcommon.HexToAddress(ZeroAddressString)}
}

// Validates the input string as an Ethereum Address
// Addresses must not be empty, have 42 character length, start with 0x and have 40 remaining characters in [0-9a-fA-F]
func ValidateEthAddress(address string) error {
	if address == "" {
		return fmt.Errorf("empty")
	}

	// An auditor recommended we should check the error of hex.DecodeString, given that geth's HexToAddress ignores it

	if has0xPrefix(address) {
		address = address[2:]
	}

	if _, err := hex.DecodeString(address); err != nil {
		return fmt.Errorf("invalid hex with error: %s", err)
	}

	if !gethcommon.IsHexAddress(address) {
		return fmt.Errorf("address(%s) doesn't pass format validation", address)
	}

	return nil
}

// Performs validation on the wrapped string
func (ea EthAddress) ValidateBasic() error {
	return ValidateEthAddress(ea.address.Hex())
}

// EthAddrLessThan migrates the Ethereum address less than function
func EthAddrLessThan(e EthAddress, o EthAddress) bool {
	return bytes.Compare([]byte(e.GetAddress().Hex()), []byte(o.GetAddress().Hex())) == -1
}

func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}
