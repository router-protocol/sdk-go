package main

func main() {
	// network := common.LoadNetwork("devnet-alpha", "k8s")
	// tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("Network", network)
	// senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
	// 	os.Getenv("HOME")+"/.routerd",
	// 	"routerd",
	// 	"file",
	// 	"gaurav",
	// 	"12345678",
	// 	"35F804E20869F9150029FC2A70AAB602B5E4606251420083B3D36200A9B10C52", // keyring will be used if pk not provided
	// 	false,
	// )

	// if err != nil {
	// 	panic(err)
	// }

	// // initialize grpc client
	// clientCtx, err := chainclient.NewClientContext(network.ChainId, senderAddress.String(), cosmosKeyring)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	// ctx, _ := context.WithTimeout(context.Background(), time.Minute)

	// routerChainClient, err := chainclient.NewChainClient(
	// 	clientCtx,
	// 	network.ChainGrpcEndpoint,
	// 	// common.OptionTLSCert(network.ChainTlsCert),
	// 	// common.OptionGasPrices("1000000000000000route"),
	// 	common.OptionGasPrices("500000000route"),
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	// allMetaInfo, err := routerChainClient.GetAllMetaInfo(ctx)
	// if err != nil {
	// 	fmt.Println("Error fetching all meta info")
	// 	panic(err)
	// }

	// fmt.Println("allMetaInfo", allMetaInfo)
	// for _, metaInfo := range allMetaInfo.MetaInfo {

	// 	if metaInfo.FeePayerApproved {
	// 		continue
	// 	}

	// 	fmt.Println("metaInfo.FeePayer", metaInfo.FeePayer, "sender", senderAddress.String())
	// 	// prepare tx msg
	// 	msg := metastoreTypes.NewMsgApproveFeepayerRequest(senderAddress.String(), metaInfo.ChainId, metaInfo.DappAddress)

	// 	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	// 	res, err := routerChainClient.SyncBroadcastMsg(msg)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Println(res.TxResponse.TxHash)

	// 	time.Sleep(time.Second * 5)
	// 	gasFee, err := routerChainClient.GetGasFee()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println("gas fee:", gasFee)
	// }
}
