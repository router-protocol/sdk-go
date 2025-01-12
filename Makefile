all:

copy-chain-types:
	mkdir -p routerchain/multichain/types
	cp ../router-chain/x/multichain/types/*.go routerchain/multichain/types
	rm -rf routerchain/multichain/types/*test.go  rm -rf routerchain/multichain/types/*gw.go
	
	mkdir -p routerchain/attestation/types
	cp ../router-chain/x/attestation/types/*.go routerchain/attestation/types
	rm -rf routerchain/attestation/types/*test.go  rm -rf routerchain/attestation/types/*gw.go

	mkdir -p routerchain/crosschain/types
	cp ../router-chain/x/crosschain/types/*.go routerchain/crosschain/types
	rm -rf routerchain/crosschain/types/*test.go  rm -rf routerchain/crosschain/types/*gw.go
	
	mkdir -p routerchain/metastore/types
	cp ../router-chain/x/metastore/types/*.go routerchain/metastore/types
	rm -rf routerchain/metastore/types/*test.go  rm -rf routerchain/metastore/types/*gw.go

	mkdir -p routerchain/pricefeed/types
	cp ../router-chain/x/pricefeed/types/*.go routerchain/pricefeed/types
	rm -rf routerchain/pricefeed/types/*test.go  rm -rf routerchain/pricefeed/types/*gw.go

	mkdir -p routerchain/voyager/types
	cp ../router-chain/x/voyager/types/*.go routerchain/voyager/types
	rm -rf routerchain/voyager/types/*test.go  rm -rf routerchain/voyager/types/*gw.go

	mkdir -p routerchain/rwasm/types
	cp ../router-chain/x/rwasm/types/*.go routerchain/rwasm/types
	rm -rf routerchain/rwasm/types/*test.go  rm -rf routerchain/rwasm/types/*gw.go

	mkdir -p routerchain/txlookup/types
	cp ../router-chain/x/txlookup/types/*.go routerchain/txlookup/types
	rm -rf routerchain/txlookup/types/*test.go  rm -rf routerchain/txlookup/types/*gw.go
	
	mkdir -p routerchain/types
	cp -r ../router-chain/types/*.go routerchain/types
	rm -rf routerchain/types/*test.go

	mkdir -p routerchain/util
	mkdir -p routerchain/proto
	cp ../router-chain/util/*.go routerchain/util
	cp -r ../router-chain/proto/* routerchain/proto/

	echo "👉 Replace router-chain/util with sdk-go/routerchain/util"
	echo "👉 Replace router-chain/x with sdk-go/routerchain"
	echo "👉 Replace router-chain/types with sdk-go/routerchain/types"