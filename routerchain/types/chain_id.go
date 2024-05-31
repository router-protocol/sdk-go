package types

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"

	tmrand "github.com/cometbft/cometbft/libs/rand"

	errorsmod "cosmossdk.io/errors"
)

var (
	regexChainID   = `[a-z]*`
	regexSeparator = `-{1}`
	regexEpoch     = `[1-9][0-9]*`
	routerChainID  = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, regexChainID, regexSeparator, regexEpoch))
)

// IsValidChainID returns false if the given chain identifier is incorrectly formatted.
func IsValidChainID(chainID string) bool {
	if len(chainID) > 48 {
		return false
	}

	return routerChainID.MatchString(chainID)
}

// ParseChainID parses a string chain identifier's epoch to an Ethereum-compatible
// chain-id in *big.Int format. The function returns an error if the chain-id has an invalid format
func ParseChainID(chainID string) (*big.Int, error) {
	chainID = strings.TrimSpace(chainID)
	if len(chainID) > 48 {
		return nil, errorsmod.Wrapf(ErrInvalidChainID, "chain-id '%s' cannot exceed 48 chars", chainID)
	}

	matches := routerChainID.FindStringSubmatch(chainID)
	if matches == nil || len(matches) != 3 || matches[1] == "" {
		return nil, errorsmod.Wrap(ErrInvalidChainID, chainID)
	}

	// verify that the chain-id entered is a base 10 integer
	chainIDInt, ok := new(big.Int).SetString(matches[2], 10)
	if !ok {
		return nil, errorsmod.Wrapf(ErrInvalidChainID, "epoch %s must be base-10 integer format", matches[2])
	}

	return chainIDInt, nil
}

// GenerateRandomChainID returns a random chain-id in the valid format.
func GenerateRandomChainID() string {
	return fmt.Sprintf("router-%d", 10+tmrand.Intn(10000))
}
