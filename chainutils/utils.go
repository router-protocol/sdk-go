package chainutils

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func DecodeBytesToAddress(txnBytes []byte, chainType multichainTypes.ChainType) (string, error) {
	if txnBytes == nil {
		return "", errors.New("transaction empty")
	}
	txnHash := string(txnBytes)
	if chainType == multichainTypes.CHAIN_TYPE_EVM {
		base64Data, _ := base64.StdEncoding.DecodeString(txnHash)
		txnHash = fmt.Sprintf("0x%s", hex.EncodeToString(base64Data))
	}
	return txnHash, nil
}
