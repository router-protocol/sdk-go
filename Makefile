all:

copy-chain-types:
	cp ../router-chain/x/multichain/types/*.go routerchain/multichain/types
	rm -rf routerchain/multichain/types/*test.go  rm -rf routerchain/multichain/types/*gw.go
	
	cp ../router-chain/x/attestation/types/*.go routerchain/attestation/types
	rm -rf routerchain/attestation/types/*test.go  rm -rf routerchain/attestation/types/*gw.go
	
	cp ../router-chain/x/inbound/types/*.go routerchain/inbound/types
	rm -rf routerchain/inbound/types/*test.go  rm -rf routerchain/inbound/types/*gw.go
	
	cp -r ../router-chain/types/*.go routerchain/types
	rm -rf routerchain/types/*test.go

	cp ../router-chain/util/*.go routerchain/util
	cp -r ../router-chain/proto/* routerchain/proto/

	echo "👉 Replace router-chain/util with sdk-go/routerchain/util
	echo "👉 Replace router-protocol/router-chain/x with sdk-go/routerchain