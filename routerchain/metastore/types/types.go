package types

import (
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewMetadataInfo(chainType multichainTypes.ChainType, chainId string, daapAddress []byte, feePayer string) *MetaInfo {
	return &MetaInfo{
		ChainType:        chainType,
		ChainId:          chainId,
		DappAddress:      daapAddress,
		FeePayer:         feePayer,
		FeePayerApproved: false,
	}
}

func NewMetadataRequest(chainType multichainTypes.ChainType, chainId string, eventNonce uint64, blockHeight uint64, daapAddress []byte, feePayer string) *MetadataRequest {
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

func NewMetadataRequestClaimHash(chainType multichainTypes.ChainType, chainId string, eventNonce uint64, blockHeight uint64, daapAddress []byte, feePayer string) *MetadataRequestClaimHash {
	return &MetadataRequestClaimHash{
		ChainType:   chainType,
		ChainId:     chainId,
		EventNonce:  eventNonce,
		BlockHeight: blockHeight,
		DaapAddress: daapAddress,
		FeePayer:    feePayer,
	}
}
