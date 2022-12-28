package util

// The go-ethereum ABI encoder *only* encodes function calls and then it only encodes
// function calls for which you provide an ABI json just like you would get out of the
// solidity compiler with your compiled contract.
// You are supposed to compile your contract, use abigen to generate an ABI , import
// this generated go module and then use for that for all testing and development.
// This abstraction layer is more trouble than it's worth, because we don't want to
// encode a function call at all, but instead we want to emulate a Solidity encode operation
// which has no equal available from go-ethereum.
//
// In order to work around this absurd series of problems we have to manually write the below
// 'function specification' that will encode the same arguments into a function call. We can then
// truncate the first several bytes where the call name is encoded to finally get the equal of the

const (
	// OutgoingBatchTxCheckpointABIJSON checks the ETH ABI for compatibility of the OutgoingBatchTx message
	OutgoingBatchTxCheckpointABIJSON = `[{
		"name": "checkpoint",
		"stateMutability": "nonpayable",
		"type": "function",
		"inputs":  [			
			{ "internalType": "bytes32","name": "_methodName","type": "bytes32"},
			{ "internalType": "uint256","name": "_chainType","type": "uint256" },
			{ "internalType": "string", "name": "_chainId","type": "string" },
			{ "internalType": "string", "name": "_sender","type": "string" },
			{"internalType": "uint256","name": "_nonce","type": "uint256"},
			{ "internalType": "uint256","name": "_relayerFee","type": "uint256"},
			{"internalType": "uint256","name": "_outgoingTxFee","type": "uint256"},
			{"internalType": "bool","name": "_isAtomic","type": "bool"},			
			{"internalType": "uint256","name": "_expTimestamp","type": "uint256"},
			{"internalType": "bytes[]","name": "_handlers","type": "bytes[]"},			
			{"internalType": "bytes[]","name": "_payloads","type": "bytes[]"}			
		],
		"outputs": [
			{ "internalType": "bytes32", "name": "", "type": "bytes32" }
		]
	}]`

	// CrossTalkRequestCheckpointABIJSON checks the ETH ABI for compatibility of the CrossTalkRequest message
	CrossTalkRequestCheckpointABIJSON = `[{
			"name": "checkpoint",
			"stateMutability": "nonpayable",
			"type": "function",
			"inputs":  [			
				{ "internalType": "bytes32","name": "_methodName","type": "bytes32"},
				{"internalType": "uint256","name": "_eventIdentifier","type": "uint256"},
				{"internalType": "uint256","name": "_crossTalkNonce","type": "uint256"},
				{ "internalType": "uint256","name": "_chainType","type": "uint256" },
				{ "internalType": "string", "name": "_chainId","type": "string" },
				{ "internalType": "string", "name": "_srcChainId","type": "string" },
				{ "internalType": "uint256","name": "_srcChainType","type": "uint256" },
				{ "internalType": "bytes","name": "_caller","type": "bytes"},
				{"internalType": "bool","name": "_isAtomic","type": "bool"},	
				{"internalType": "uint256","name": "_expTimestamp","type": "uint256"},
				{"internalType": "bytes[]","name": "_handlers","type": "bytes[]"},			
				{"internalType": "bytes[]","name": "_payloads","type": "bytes[]"}			
			],
			"outputs": [
				{ "internalType": "bytes32", "name": "", "type": "bytes32" }
			]
		}]`

	// CrossTalkAckRequestCheckpointABIJSON checks the ETH ABI for compatibility of the CrossTalkRequest message
	CrossTalkAckRequestCheckpointABIJSON = `[{
		"name": "checkpoint",
		"stateMutability": "nonpayable",
		"type": "function",
		"inputs":  [			
			{ "internalType": "bytes32","name": "_methodName","type": "bytes32"},
			{"internalType": "uint256","name": "_eventIdentifier","type": "uint256"},
			{"internalType": "uint256","name": "_crossTalkNonce","type": "uint256"},
			{ "internalType": "uint256","name": "_chainType","type": "uint256" },
			{ "internalType": "string", "name": "_chainId","type": "string" },
			{ "internalType": "uint256","name": "_destChainType","type": "uint256" },
			{ "internalType": "string", "name": "_destChainId","type": "string" },			
			{ "internalType": "bytes","name": "_caller","type": "bytes"},
			{"internalType": "bool[]","name": "_execFlags","type": "bool[]"},				
			{"internalType": "bytes[]","name": "_execData","type": "bytes[]"}					
		],
		"outputs": [
			{ "internalType": "bytes32", "name": "", "type": "bytes32" }
		]
	}]`

	// ValsetCheckpointABIJSON checks the ETH ABI for compatibility of the Valset update message
	ValsetCheckpointABIJSON = `[{
		"name": "checkpoint",
		"stateMutability": "pure",
		"type": "function",
		"inputs": [
			{ "internalType": "bytes32",   "name": "_routerId",   "type": "bytes32"   },
			{ "internalType": "bytes32",   "name": "_checkpoint",  "type": "bytes32"   },
			{ "internalType": "uint256",   "name": "_valsetNonce", "type": "uint256"   },
			{ "internalType": "address[]", "name": "_validators",  "type": "address[]" },
			{ "internalType": "uint256[]", "name": "_powers",      "type": "uint256[]" }			
		],
		"outputs": [
			{ "internalType": "bytes32", "name": "", "type": "bytes32" }
		]
	}]`

	// OutgoingLogicCallABIJSON checks the ETH ABI for compatibility of the logic call message
	OutgoingLogicCallABIJSON = `[{
	  "name": "checkpoint",
      "outputs": [],
      "stateMutability": "pure",
      "type": "function",
      "inputs": [
			{ "internalType": "bytes32",   "name": "_gravityId",                "type": "bytes32"   },
			{ "internalType": "bytes32",   "name": "_methodName",             "type": "bytes32"   },
			{ "internalType": "uint256[]", "name": "_transferAmounts",        "type": "uint256[]" },
			{ "internalType": "address[]", "name": "_transferTokenContracts", "type": "address[]" },
			{ "internalType": "uint256[]", "name": "_feeAmounts",             "type": "uint256[]" },
			{ "internalType": "address[]", "name": "_feeTokenContracts",      "type": "address[]" },
			{ "internalType": "address",   "name": "_logicContractAddress",   "type": "address"   },
			{ "internalType": "bytes",     "name": "_payload",                "type": "bytes"     },
			{ "internalType": "uint256",   "name": "_timeout",                "type": "uint256"   },
			{ "internalType": "bytes32",   "name": "_invalidationId",         "type": "bytes32"   },
			{ "internalType": "uint256",   "name": "_invalidationNonce",      "type": "uint256"   }
      ]
    }]`
)
