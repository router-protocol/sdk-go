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
	// CrosschainRequestCheckpointABIJSON checks the ETH ABI for compatibility of the CrossTalkRequest message
	CrosschainRequestCheckpointABIJSON = `[{
		"name": "checkpoint",
		"stateMutability": "nonpayable",
		"type": "function",
		"inputs":  [			
			{ "internalType": "bytes32","name": "_methodName","type": "bytes32"},
			{"internalType": "uint256","name": "_routeAmount","type": "uint256"},
			{"internalType": "uint256","name": "_requestIdentifier","type": "uint256"},
			{"internalType": "uint256","name": "_requestTimestamp","type": "uint256"},
			{ "internalType": "string", "name": "_srcChainId","type": "string" },			
			{ "internalType": "address", "name": "_routeRecipient", "type": "address"},
			{"internalType": "string","name": "_destChainId","type": "string"},
			{ "internalType": "address", "name": "_asmAddress", "type": "address"},						
			{ "internalType": "string","name": "_requestSender","type": "string"},		
			{ "internalType": "address", "name": "_handlerAddress", "type": "address"},
			{ "internalType": "bytes","name": "_packet","type": "bytes"},
			{"internalType": "bool","name": "_isReadCall","type": "bool"}			
		],
		"outputs": [
			{ "internalType": "bytes32", "name": "", "type": "bytes32" }
		]
	}]`

	// CrosschainRequestNearCheckpointABIJSON checks the ETH ABI for compatibility of the CrossTalkRequest message
	CrosschainRequestNearCheckpointABIJSON = `[{
		"name": "checkpoint",
		"stateMutability": "nonpayable",
		"type": "function",
		"inputs":  [			
			{ "internalType": "bytes32","name": "_methodName","type": "bytes32"},
			{"internalType": "uint256","name": "_routeAmount","type": "uint256"},
			{"internalType": "uint256","name": "_requestIdentifier","type": "uint256"},
			{"internalType": "uint256","name": "_requestTimestamp","type": "uint256"},
			{ "internalType": "string", "name": "_srcChainId","type": "string" },			
			{ "internalType": "string", "name": "_routeRecipient", "type": "string"},
			{"internalType": "string","name": "_destChainId","type": "string"},
			{ "internalType": "string", "name": "_asmAddress", "type": "string"},						
			{ "internalType": "string","name": "_requestSender","type": "string"},		
			{ "internalType": "string", "name": "_handlerAddress", "type": "string"},
			{ "internalType": "bytes","name": "_packet","type": "bytes"},
			{"internalType": "bool","name": "_isReadCall","type": "bool"}			
		],
		"outputs": [
			{ "internalType": "bytes32", "name": "", "type": "bytes32" }
		]
	}]`

	// CrosschainRequestAlephZeroCheckpointABIJSON checks the ETH ABI for compatibility of the CrossTalkRequest message
	CrosschainRequestAlephZeroCheckpointABIJSON = `[{
			"name": "checkpoint",
			"stateMutability": "nonpayable",
			"type": "function",
			"inputs":  [			
				{ "internalType": "bytes32","name": "_methodName","type": "bytes32"},
				{"internalType": "uint256","name": "_routeAmount","type": "uint256"},
				{"internalType": "uint256","name": "_requestIdentifier","type": "uint256"},
				{"internalType": "uint256","name": "_requestTimestamp","type": "uint256"},
				{ "internalType": "string", "name": "_srcChainId","type": "string" },			
				{ "internalType": "bytes32", "name": "_routeRecipient", "type": "bytes32"},
				{"internalType": "string","name": "_destChainId","type": "string"},
				{ "internalType": "bytes32", "name": "_asmAddress", "type": "bytes32"},						
				{ "internalType": "string","name": "_requestSender","type": "string"},		
				{ "internalType": "bytes32", "name": "_handlerAddress", "type": "bytes32"},
				{ "internalType": "bytes","name": "_packet","type": "bytes"},
				{"internalType": "bool","name": "_isReadCall","type": "bool"}			
			],
			"outputs": [
				{ "internalType": "bytes32", "name": "", "type": "bytes32" }
			]
		}]`

	// CrosschainAckRequestCheckpointABIJSON checks the ETH ABI for compatibility of the CrossTalkAckRequest message
	CrosschainAckRequestCheckpointABIJSON = `[{
		"name": "checkpoint",
		"stateMutability": "nonpayable",
		"type": "function",
		"inputs":  [			
			{ "internalType": "bytes32","name": "_methodName","type": "bytes32"},
			{ "internalType": "string", "name": "_chainId","type": "string" },
			{"internalType": "uint256","name": "_requestIdentifier","type": "uint256"},
			{"internalType": "uint256","name": "_ackRequestIdentifier","type": "uint256"},
			{ "internalType": "string", "name": "_destChainId","type": "string" },	
			{ "internalType": "address","name": "_requestSender","type": "address"},
			{"internalType": "bytes","name": "_execData","type": "bytes"},	
			{"internalType": "bool","name": "execFlag", "type": "bool"}				
		],
		"outputs": [
			{ "internalType": "bytes32", "name": "", "type": "bytes32" }
		]
	}]`

	// CrosschainAckRequestNearCheckpointABIJSON checks the ETH ABI for compatibility of the CrossTalkAckRequest message
	CrosschainAckRequestNearCheckpointABIJSON = `[{
		"name": "checkpoint",
		"stateMutability": "nonpayable",
		"type": "function",
		"inputs":  [			
			{ "internalType": "bytes32","name": "_methodName","type": "bytes32"},
			{ "internalType": "string", "name": "_chainId","type": "string" },
			{"internalType": "uint256","name": "_requestIdentifier","type": "uint256"},
			{"internalType": "uint256","name": "_ackRequestIdentifier","type": "uint256"},
			{ "internalType": "string", "name": "_destChainId","type": "string" },	
			{ "internalType": "string","name": "_requestSender","type": "string"},
			{"internalType": "bytes","name": "_execData","type": "bytes"},	
			{"internalType": "bool","name": "execFlag", "type": "bool"}				
		],
		"outputs": [
			{ "internalType": "bytes32", "name": "", "type": "bytes32" }
		]
	}]`

	// CrosschainAckRequestNearCheckpointABIJSON checks the ETH ABI for compatibility of the CrossTalkAckRequest message
	CrosschainAckRequestAlephZeroCheckpointABIJSON = `[{
			"name": "checkpoint",
			"stateMutability": "nonpayable",
			"type": "function",
			"inputs":  [			
				{ "internalType": "bytes32","name": "_methodName","type": "bytes32"},
				{ "internalType": "string", "name": "_chainId","type": "string" },
				{"internalType": "uint256","name": "_requestIdentifier","type": "uint256"},
				{"internalType": "uint256","name": "_ackRequestIdentifier","type": "uint256"},
				{ "internalType": "string", "name": "_destChainId","type": "string" },	
				{ "internalType": "bytes32","name": "_requestSender","type": "bytes32"},
				{"internalType": "bytes","name": "_execData","type": "bytes"},	
				{"internalType": "bool","name": "execFlag", "type": "bool"}				
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
			{ "internalType": "bytes32","name": "_methodName","type": "bytes32"},			
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

	// OutboundBatchRequestEventABIJSON checks the ETH ABI for compatibility of the Outbound batch event
	OutboundBatchRequestEventABIJSON = `[
		{
			"anonymous": false,
	"inputs": [
			  {"indexed": false,"internalType": "uint64","name": "destinationChainType","type": "uint64"},
			  {"indexed": false,"internalType": "string","name": "destinationChainId","type": "string"},
			  {"indexed": false,"internalType": "bytes[]","name": "handlers","type": "bytes[]"},
			  {"indexed": false,"internalType": "bytes[]","name": "payloads","type": "bytes[]"},
			  {"indexed": false,"internalType": "uint256","name": "relayerFee","type": "uint256"},
			  {"indexed": false,"internalType": "uint256","name": "gasLimit","type": "uint256"},
			  {"indexed": false,"internalType": "uint256","name": "gasPrice","type": "uint256"},
			  {"indexed": false,"internalType": "uint64","name": "outbound_ack_gas_limit","type": "uint64"},
			  {"indexed": false,"internalType": "bool","name": "isAtomic","type": "bool"},
			  {"indexed": false,"internalType": "uint256","name": "expTimestamp","type": "uint256"},
			  {"indexed": false,"internalType": "uint256","name": "routeAmount","type": "uint256"},
			  {"indexed": false,"internalType": "bytes","name": "routeRecipient","type": "bytes"},
			  {"indexed": false,"internalType": "bytes","name": "asmAddress","type": "bytes"}
			],
			"name": "OutboundBatchRequest",
			"type": "event"
		  }
	]`

	// InboundSudoRequestABIJSON checks the ETH ABI for compatibility of the Inbound Batch request
	InboundSudoRequestABIJSON = `[{
		"inputs": [
			{"internalType": "bytes","name": "_sender","type": "bytes"},
			{"internalType": "string","name": "_srcChainId","type": "string"},
			{"internalType": "uint64","name": "_srcChainType","type": "uint64"},
			{"internalType": "uint64","name": "_eventNonce","type": "uint64"},
			{"internalType": "bytes","name": "_payload","type": "bytes"},
			{"internalType": "bytes","name": "_asmAddress","type": "bytes"}
			
		],
		"name": "inboundSudo",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}]`

	// OutboundAckSudoRequestABIJSON checks the ETH ABI for compatibility of the Outbound Ack request
	OutboundAckSudoRequestABIJSON = `[{
		"inputs": [
		  {"internalType": "bytes","name": "outboundTxRequestedBy","type": "bytes"},
		  {"internalType": "string","name": "destChainId","type": "string"},
		  {"internalType": "uint64","name": "destChainType","type": "uint64"},
		  {"internalType": "uint64","name": "outboundBatchNonce","type": "uint64"},
		  {"internalType": "uint64","name": "execCode","type": "uint64"},
		  {"internalType": "bool","name": "execStatus","type": "bool"},
		  {"internalType": "bool[]","name": "execFlags","type": "bool[]"},
		  {"internalType": "bytes[]","name": "execData","type": "bytes[]"},
		  {"internalType": "uint256","name": "refund_amount","type": "uint256"}
		],
		"name": "outboundAckSudo",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	  }]`
)
