package types

func NewMetadataInfo(chainId string, daapAddress string, feePayer string) *MetaInfo {
	return &MetaInfo{
		ChainId:          chainId,
		DappAddress:      daapAddress,
		FeePayer:         feePayer,
		FeePayerApproved: false,
	}
}

func NewMetadataRequest(chainId string, eventNonce uint64, blockHeight uint64, daapAddress string, feePayer string) *MetadataRequest {
	return &MetadataRequest{
		ChainId:     chainId,
		EventNonce:  eventNonce,
		BlockHeight: blockHeight,
		DaapAddress: daapAddress,
		FeePayer:    feePayer,
		Status:      META_TX_CREATED,
	}
}

func NewMetadataRequestClaimHash(chainId string, eventNonce uint64, blockHeight uint64, daapAddress string, feePayer string) *MetadataRequestClaimHash {
	return &MetadataRequestClaimHash{
		ChainId:     chainId,
		EventNonce:  eventNonce,
		BlockHeight: blockHeight,
		DaapAddress: daapAddress,
		FeePayer:    feePayer,
	}
}
