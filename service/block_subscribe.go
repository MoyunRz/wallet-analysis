package service

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"wallet-analysis/common/log"
	"wallet-analysis/utils"
)

var logChan = make(chan utils.Log)

// StartSubscribe
// 开始区块事件订阅
func StartSubscribe(cAddress string) {
	cli := utils.EthClient
	if cli == nil {
		log.Error("开始区块事件订阅失败，无法进行socket连接eth")
		return
	}
	defer cli.Close()
	// 订阅的合约地址
	contractAddress := common.HexToAddress(cAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := cli.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:

			log.Infof("订阅事件 交易hash: %s", vLog.TxHash.String())
			var topics []string
			for i := 0; i < len(vLog.Topics); i++ {
				topics = append(topics, vLog.Topics[i].String())
			}

			implEventByLogs(topics, vLog.Data, int(vLog.Index))
		}
	}
}
