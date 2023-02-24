package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

type Block struct {
	Number       hexutil.Big    `json:"number"`
	Miner        string         `json:"miner"`
	Hash         string         `json:"hash"`
	ParentHash   string         `json:"parentHash"`
	ReceiptsRoot string         `json:"receiptsRoot"`
	StateRoot    string         `json:"stateRoot"`
	Difficulty   hexutil.Big    `json:"difficulty"`
	GasLimit     hexutil.Big    `json:"gasLimit"`
	GasUsed      hexutil.Big    `json:"gasUsed"`
	Timestamp    hexutil.Big    `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
}

type Transaction struct {
	Hash             string         `json:"hash"`
	BlockHash        string         `json:"blockHash"`
	BlockNumber      hexutil.Uint64 `json:"blockNumber"`
	From             string         `json:"from"`
	Gas              hexutil.Uint64 `json:"gas"`
	GasPrice         hexutil.Big    `json:"gasPrice"`
	Input            string         `json:"input"`
	Nonce            hexutil.Uint64 `json:"nonce"`
	To               string         `json:"to"`
	TransactionIndex hexutil.Uint64 `json:"transactionIndex"`
	Value            hexutil.Big    `json:"value"`
}

func (rpc *RpcClient) BlockNumber() (int64, error) {
	var result hexutil.Uint64
	err := rpc.CallNoAuth("eth_blockNumber", &result)
	return int64(result), err
}

func (rpc *RpcClient) BlockByNumber(h int64) (*Block, error) {
	var result Block
	err := rpc.CallNoAuth("eth_getBlockByNumber", &result, hexutil.Uint64(h).String(), true)
	return &result, err
}

func (rpc *RpcClient) BlockByHash(hash string) (*Block, error) {
	var result Block
	err := rpc.CallNoAuth("eth_getBlockByHash", &result, hash, true)
	return &result, err
}

type TransactionReceipt struct {
	TransactionHash   string         `json:"transactionHash"`
	BlockHash         string         `json:"blockHash"`
	BlockNumber       hexutil.Uint64 `json:"blockNumber"`
	CumulativeGasUsed hexutil.Big    `json:"cumulativeGasUsed"`
	From              string         `json:"from"`
	GasUsed           hexutil.Big    `json:"gasUsed"`
	To                string         `json:"to"`
	Status            hexutil.Uint   `json:"status"`
	Logs              []*Log         `json:"logs"`
}
type Log struct {
	Address     string         `json:"address"`
	Topics      []string       `json:"topics"`
	Data        string         `json:"data"`
	BlockNumber hexutil.Uint64 `json:"blockNumber"`
	LogIndex    hexutil.Uint   `json:"logIndex"`
}

func (rpc *RpcClient) EthBalanceByAddress(addr string) (decimal.Decimal, error) {
	var result string
	err := rpc.CallNoAuth("eth_getBalance", &result, addr, "latest")
	if err != nil {
		return decimal.Decimal{}, err
	}
	decodeBig, err := hexutil.DecodeBig(result)
	if err != nil {
		return decimal.Decimal{}, err
	}
	fromString := decimal.NewFromBigInt(decodeBig, 18)
	return fromString, err
}

func (rpc *RpcClient) TransactionReceipt(txhash string) (*TransactionReceipt, error) {
	var result TransactionReceipt
	err := rpc.CallNoAuth("eth_getTransactionReceipt", &result, txhash)
	return &result, err
}

func (rpc *RpcClient) TransactionByHash(txhash string) (*Transaction, error) {
	var result Transaction
	err := rpc.CallNoAuth("eth_getTransactionByHash", &result, txhash)
	//result.Value = hexutil.Big(*big.NewInt(1233333333))
	return &result, err
}

func (rpc *RpcClient) GetCode(address string) (bool, error) {
	var getCode string

	err := rpc.CallNoAuth("eth_getCode", &getCode, address, "latest")
	if err != nil {
		fmt.Println("错误:", err)
	}
	//fmt.Println("获得的代码:",getCode)
	if getCode == "0x" {
		return false, nil
	} else {
		return true, nil
	}
}
