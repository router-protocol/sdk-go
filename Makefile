all:

copy-chain-types:
	cp ../router-chain/x/multichain/types/*.go routerchain/multichain/types
	rm -rf routerchain/multichain/types/*test.go  rm -rf routerchain/multichain/types/*gw.go
	
	cp ../router-chain/x/attestation/types/*.go routerchain/attestation/types
	rm -rf routerchain/attestation/types/*test.go  rm -rf routerchain/attestation/types/*gw.go

	cp ../router-chain/x/crosschain/types/*.go routerchain/crosschain/types
	rm -rf routerchain/crosschain/types/*test.go  rm -rf routerchain/crosschain/types/*gw.go
	
	cp ../router-chain/x/metastore/types/*.go routerchain/metastore/types
	rm -rf routerchain/metastore/types/*test.go  rm -rf routerchain/metastore/types/*gw.go

	cp ../router-chain/x/pricefeed/types/*.go routerchain/pricefeed/types
	rm -rf routerchain/pricefeed/types/*test.go  rm -rf routerchain/pricefeed/types/*gw.go

	cp ../router-chain/x/voyager/types/*.go routerchain/voyager/types
	rm -rf routerchain/voyager/types/*test.go  rm -rf routerchain/voyager/types/*gw.go

	cp ../router-chain/x/rwasm/types/*.go routerchain/rwasm/types
	rm -rf routerchain/rwasm/types/*test.go  rm -rf routerchain/rwasm/types/*gw.go
	
	cp -r ../router-chain/types/*.go routerchain/types
	rm -rf routerchain/types/*test.go

	cp ../router-chain/util/*.go routerchain/util
	cp -r ../router-chain/proto/* routerchain/proto/

	echo "👉 Replace router-chain/util with sdk-go/routerchain/util"
	echo "👉 Replace router-chain/x with sdk-go/routerchain"
	echo "👉 Replace router-chain/types with sdk-go/routerchain/types"