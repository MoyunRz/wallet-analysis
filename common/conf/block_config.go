package conf

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"wallet-analysis/common/log"
)

const EventAbiPath = "event.abi"

var EventAbi abi.ABI
var EthClient *ethclient.Client

func init() {
	var err error
	EventAbi, err = GetABI()
	if err != nil {
		log.Error("获取ABI对象 event path error:", err)
		log.Fatal(err)
		return
	}

	EthClient, err = ethclient.Dial("ws://192.168.19.66:7545")
	if err != nil {
		log.Error("连接私链失败:", err)
		log.Fatal(err)

		return
	}
}

// GetABI
// 获取ABI对象
func GetABI() (abi.ABI, error) {

	reader, err := os.Open(EventAbiPath)
	if err != nil {
		log.Fatal(err)
	}
	wrapABI, err := abi.JSON(reader)
	if err != nil {
		log.Fatal(err)
	}
	return wrapABI, err
}
