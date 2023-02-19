package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"wallet-analysis/common/conf"
	"wallet-analysis/common/log"
	"wallet-analysis/utils"
)

var Rpc *utils.RpcClient

const ABI1155 = "erc1155.abi"
const ForwarderAbi = "forwarder.abi"

func init() {
	Rpc = utils.NewRpcClient(conf.Cfg.Host)
}

func ScanBlock() {
	lasterHeight, err := Rpc.BlockNumber()
	if err != nil {
		log.Fatal(err.Error())
	}
	// 循环解析
	for i := conf.Cfg.StartHeight; i < lasterHeight; i++ {
		// 根据区块高度获取区块
		block, err := GetBlockByHeight(i)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		// 处理区块内的交易
		GetTxInfoByHash(block.Transactions)
	}
}

// GetBlockByHeight
// 根据区块高度获取区块
func GetBlockByHeight(blockHeight int) (*utils.Block, error) {
	// 根据高度获取区块
	return Rpc.BlockByNumber(blockHeight)

}

// GetBlockByBlockHash
// 根据区块hash获取区块
func GetBlockByBlockHash(blockHash string) (*utils.Block, error) {
	return Rpc.BlockByHash(blockHash)
}

// GetTxInfoByHash
// 根据交易hash获取区块
func GetTxInfoByHash(transactions []*utils.Transaction) {
	// 遍历解析交易
	for i := 0; i < len(transactions); i++ {
		receipt, err := Rpc.TransactionReceipt(transactions[i].Hash)
		if err != nil {
			return
		}
		// 交易事件处理
		fmt.Println(receipt)
		EventHandle(receipt.Logs, receipt.TransactionHash)
	}
}

func ForwarderInputData(methodName string, res map[string]interface{}) {
	v := res["req"].(struct {
		From  common.Address "json:\"from\""
		To    common.Address "json:\"to\""
		Value *big.Int       "json:\"value\""
		Gas   *big.Int       "json:\"gas\""
		Nonce *big.Int       "json:\"nonce\""
		Data  []uint8        "json:\"data\""
	})
	fmt.Printf("解析 input 结果:%x \n", v.Data)
}

func EventHandle(vLog []*utils.Log, hash string) {
	fmt.Println("交易 vLog", vLog)
	fmt.Println("交易hash", hash)
	for j := 0; j < len(vLog); j++ {
		vlog := vLog[j]
		implEventByLogs(*vlog)
	}
}
