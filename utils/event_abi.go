package utils

const eventAbiStr = `[
    {
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "account",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "operator",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "bool",
				"name": "approved",
				"type": "bool"
			}
		],
		"name": "ApprovalForAll",
		"type": "event"
	},
	{
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "bytes32",
                "name": "role",
                "type": "bytes32"
            },
            {
                "indexed": true,
                "internalType": "bytes32",
                "name": "previousAdminRole",
                "type": "bytes32"
            },
            {
                "indexed": true,
                "internalType": "bytes32",
                "name": "newAdminRole",
                "type": "bytes32"
            }
        ],
        "name": "RoleAdminChanged",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "bytes32",
                "name": "role",
                "type": "bytes32"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "account",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "RoleGranted",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "bytes32",
                "name": "role",
                "type": "bytes32"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "account",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "RoleRevoked",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "operator",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256[]",
                "name": "ids",
                "type": "uint256[]"
            },
            {
                "indexed": false,
                "internalType": "uint256[]",
                "name": "values",
                "type": "uint256[]"
            }
        ],
        "name": "TransferBatch",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "operator",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "from",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "to",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "id",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "TransferSingle",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "string",
                "name": "value",
                "type": "string"
            },
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "id",
                "type": "uint256"
            }
        ],
        "name": "URI",
        "type": "event"
    },
    {
		"anonymous": false,
		"inputs": [
			{
				"components": [
					{
						"internalType": "address",
						"name": "from",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "to",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "value",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gas",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "nonce",
						"type": "uint256"
					},
					{
						"internalType": "bytes",
						"name": "data",
						"type": "bytes"
					}
				],
				"indexed": false,
				"internalType": "struct MetaTxForwarder.ForwardRequest",
				"name": "",
				"type": "tuple"
			}
		],
		"name": "ProxyLog",
		"type": "event"
	},
	{
    		"anonymous": false,
    		"inputs": [
    			{
    				"indexed": false,
    				"internalType": "address",
    				"name": "",
    				"type": "address"
    			},
    			{
    				"indexed": false,
    				"internalType": "uint256[]",
    				"name": "",
    				"type": "uint256[]"
    			},
    			{
    				"indexed": false,
    				"internalType": "uint256[]",
    				"name": "",
    				"type": "uint256[]"
    			},
    			{
    				"indexed": false,
    				"internalType": "enum XunWenGe.TxEnum",
    				"name": "",
    				"type": "uint8"
    			}
    		],
    		"name": "BurnLog",
    		"type": "event"
    	},
    	{
    		"anonymous": false,
    		"inputs": [
    			{
    				"indexed": false,
    				"internalType": "address[]",
    				"name": "",
    				"type": "address[]"
    			},
    			{
    				"indexed": false,
    				"internalType": "uint256[]",
    				"name": "",
    				"type": "uint256[]"
    			},
    			{
    				"indexed": false,
    				"internalType": "uint256[]",
    				"name": "",
    				"type": "uint256[]"
    			},
    			{
    				"indexed": false,
    				"internalType": "enum XunWenGe.TxEnum",
    				"name": "",
    				"type": "uint8"
    			},
    			{
    				"indexed": false,
    				"internalType": "bytes",
    				"name": "",
    				"type": "bytes"
    			}
    		],
    		"name": "MintLog",
    		"type": "event"
    	},
	{
    		"anonymous": false,
    		"inputs": [
    			{
    				"indexed": false,
    				"internalType": "address",
    				"name": "",
    				"type": "address"
    			},
    			{
    				"indexed": false,
    				"internalType": "address[]",
    				"name": "",
    				"type": "address[]"
    			},
    			{
    				"indexed": false,
    				"internalType": "uint256[]",
    				"name": "",
    				"type": "uint256[]"
    			},
    			{
    				"indexed": false,
    				"internalType": "uint256[]",
    				"name": "",
    				"type": "uint256[]"
    			},
    			{
    				"indexed": false,
    				"internalType": "enum XunWenGe.TxEnum",
    				"name": "",
    				"type": "uint8"
    			},
    			{
    				"indexed": false,
    				"internalType": "bytes",
    				"name": "data",
    				"type": "bytes"
    			}
    		],
    		"name": "TransferLog",
    		"type": "event"
    	}
]
`
