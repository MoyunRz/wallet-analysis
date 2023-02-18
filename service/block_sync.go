package service

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
	"strings"
	"wallet-analysis/common/conf"
	"wallet-analysis/common/log"
	"wallet-analysis/utils"
)

var Rpc *utils.RpcClient

const ABI1155 = "erc1155.abi"
const ForwarderAbi = "forwarder.abi"

var logs = make(chan utils.Log)

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
		fmt.Println(receipt)
		EventDea(receipt.Logs, receipt.TransactionHash)

		//code, err := Rpc.GetCode(transactions[i].To)
		//if err != nil {
		//	log.Fatal(err.Error())
		//	return
		//}
		//if code {
		//	// 合约交易
		//	//ResolveTxInput(transactions[i].To, transactions[i].Input)
		//	fmt.Println("")
		//} else {
		//	// 普通交易
		//
		//}
	}
}

// ResolveTxInput
// 解析交易的Input数据
func ResolveTxInput(contractAddress, encodedData string) {
	contractType := ""
	fpath := ""
	contractAddress = strings.ToLower(contractAddress)
	if contractAddress == strings.ToLower("0x645F483704B557625893f18c4ceb8ECdCdF7094F") {
		fpath = ForwarderAbi
		contractType = "Forwarder"
	}
	if contractAddress == strings.ToLower("0x6Cf015d91f18ec8E5bC5915366EA5e560Cbb6B31") {
		fpath = ABI1155
		contractType = "ABI1155"
	}
	fmt.Println(contractType)
	reader, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	tokenAbi, err := abi.JSON(reader)
	if err != nil {
		log.Fatal(err)
	}

	decodedSig, err := hex.DecodeString(encodedData[2:10])
	if err != nil {
		log.Fatal(err)
	}
	// recover Method from signature and ABI
	method, err := tokenAbi.MethodById(decodedSig)
	if err != nil {
		log.Fatal(err)
	}
	// decode txInput Payload
	decodedData, err := hex.DecodeString(encodedData[10:])
	if err != nil {
		log.Fatal(err)
	}

	var s = make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(s, decodedData)
	if err != nil {
		return
	}
	// 进行单独解析
	if contractType == "ABI1155" {
		ERC1155InputData(method.Name, s)
	}
	// 进行单独解析
	if contractType == "Forwarder" {
		ForwarderInputData(method.Name, s)
	}
}

func ERC1155InputData(methodName string, res map[string]interface{}) {

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
