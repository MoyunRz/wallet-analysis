package utils

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"strings"
	"wallet-analysis/common/log"
)

var EventAbi abi.ABI
var EthClient *ethclient.Client

func InitClient() {
	var err error
	//EventAbi, err = GetABI()
	EventAbi, err = GetStringToABI(eventAbiStr)
	if err != nil {
		log.Error("获取ABI对象 event path error:", err)
		log.Fatal(err)
		return
	}
	EthClient, err = ethclient.Dial("ws://119.23.219.232:9944")
	if err != nil {
		log.Error("连接私链失败:", err)
		log.Fatal(err)
		return
	}
}

func MakeClient() (*ethclient.Client, error) {
	ethClient, err := ethclient.Dial("ws://119.23.219.232:9944")
	if err != nil {
		log.Error("连接私链失败:", err)
		log.Fatal(err)
		return nil, err
	}
	return ethClient, nil
}

// GetABI
// 获取ABI对象
func GetABI() (abi.ABI, error) {

	reader, err := os.Open("")
	if err != nil {
		log.Fatal(err)
	}
	wrapABI, err := abi.JSON(reader)
	if err != nil {
		log.Fatal(err)
	}
	return wrapABI, err
}

// GetStringToABI
// 获取ABI对象
func GetStringToABI(str string) (abi.ABI, error) {

	wrapABI, err := abi.JSON(strings.NewReader(str))
	if err != nil {
		return abi.ABI{}, err
	}
	return wrapABI, err
}

func String(d interface{}) string {
	str, _ := json.Marshal(d)
	return string(str)
}
