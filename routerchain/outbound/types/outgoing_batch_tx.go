package types

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

func (i OutgoingBatchTx) ValidateBasic() error {
	//TODO: Validate id?
	//TODO: Validate cosmos sender?

	return nil
}

// GetCheckpoint gets the checkpoint signature from the given outgoing tx batch
func (outgoingBatchTx OutgoingBatchTx) GetCheckpoint(routerIDstring string) []byte {

	abi, err := abi.JSON(strings.NewReader(util.OutgoingBatchTxCheckpointABIJSON))
	if err != nil {
		panic("Bad ABI constant!")
	}

	// the contract argument is not a arbitrary length array but a fixed length 32 byte
	// array, therefore we have to utf8 encode the string (the default in this case) and
	// then copy the variable length encoded data into a fixed length array. This function
	// will panic if gravityId is too long to fit in 32 bytes
	// routerID, err := util.StrToFixByteArray(routerIDstring)
	// if err != nil {
	// 	panic(err)
	// }

	// Create the methodName argument which salts the signature
	methodNameBytes := []uint8("requestFromRouter")
	var batchMethodName [32]uint8
	copy(batchMethodName[:], methodNameBytes)

	// Run through the elements of the batch and serialize them
	payloads := make([][]byte, len(outgoingBatchTx.ContractCalls))
	handlers := make([][]byte, len(outgoingBatchTx.ContractCalls))

	for j, contractCal := range outgoingBatchTx.ContractCalls {
		handlers[j] = contractCal.DestinationContractAddress
		payloads[j] = contractCal.Payload
	}

	chainType := &big.Int{}
	chainType.SetUint64(uint64(outgoingBatchTx.DestinationChainType))

	batchNonce := &big.Int{}
	batchNonce.SetUint64(outgoingBatchTx.Nonce)

	// the methodName needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	abiEncodedBatch, err := abi.Pack("checkpoint",
		batchMethodName,
		chainType,
		outgoingBatchTx.DestinationChainId,
		outgoingBatchTx.SourceAddress,
		batchNonce,
		outgoingBatchTx.RelayerFee.Amount.BigInt(),
		outgoingBatchTx.OutgoingTxFee.Amount.BigInt(),
		outgoingBatchTx.IsAtomic,
		handlers,
		payloads,
	)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if err != nil {
		panic(fmt.Sprintf("Error packing checkpoint! %s/n", err))
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you were to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	return crypto.Keccak256Hash(abiEncodedBatch[4:]).Bytes()
}
