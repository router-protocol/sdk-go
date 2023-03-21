package types

import (
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewMetadataRequest(chainType multichainTypes.ChainType, chainId string, eventNonce uint64, blockHeight uint64, daapAddress string, feePayer string) *MetadataRequest {
	return &MetadataRequest{
		ChainType:   chainType,
		ChainId:     chainId,
		EventNonce:  eventNonce,
		BlockHeight: blockHeight,
		DaapAddress: daapAddress,
		FeePayer:    feePayer,
		Status:      META_TX_CREATED,
	}
}

func NewMetadataRequestClaimHash(chainType multichainTypes.ChainType, chainId string, eventNonce uint64, blockHeight uint64, daapAddress string, feePayer string) *MetadataRequestClaimHash {
	return &MetadataRequestClaimHash{
		ChainType:   chainType,
		ChainId:     chainId,
		EventNonce:  eventNonce,
		BlockHeight: blockHeight,
		DaapAddress: daapAddress,
		FeePayer:    feePayer,
	}
}
