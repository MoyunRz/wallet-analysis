package service

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"
	"time"
	"wallet-analysis/common/conf"
	"wallet-analysis/common/db"
	"wallet-analysis/common/log"
	"wallet-analysis/models/blocks"
	"wallet-analysis/utils"
)

var Rpc *utils.RpcClient

func init() {
	Rpc = utils.NewRpcClient(conf.Cfg.Host)
}

func ScanBlock() {
	// 计算间隔休眠时间
	// 设12秒一个块
	// (结束的时间 - 处理的开始时间)/12 = 出块的个数
	// 出块的个数 > 12 则不休眠
	// 出块的个数 < 12 则：(12 - 出块的个数) * 12 = 休眠时间
	sleepTime := 1

	for {
		scanHeight := int64(0)
		blockInfo := blocks.MakeBlockInfo(nil)
		startTime := time.Now().Unix()
		// 获取数据库的区块高度
		if conf.Cfg.IsReStartScan {
			log.Info("根据配置文件高度来扫描 高度 ===> ", conf.Cfg.StartHeight)
			scanHeight = conf.Cfg.StartHeight
		} else {
			err := blockInfo.GetMaxHeight()
			if err != nil {
				log.Fatal(err.Error())
				return
			}
			log.Info("根据数据库高度来扫描 高度 ===> ", blockInfo.Height)
			scanHeight = blockInfo.Height
		}

		lasterHeight, err := Rpc.BlockNumber()
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		var wg sync.WaitGroup
		ch := make(chan struct{}, 12)
		// 循环解析
		for i := scanHeight; i < (lasterHeight - 12); i++ {
			ch <- struct{}{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				log.Info("扫描高度 ---> ", i)
				// 根据区块高度获取区块
				block, err := GetBlockByHeight(i)
				if err != nil {
					log.Fatal(err.Error())
					return
				}
				// 处理区块内的交易
				GetTxInfoByHash(block)
				// 处理完成 写入数据库
				writeBlockToDB(block)
				time.Sleep(time.Duration(500))
				<-ch
			}()
		}
		wg.Wait()
		endTime := time.Now().Unix()
		newBlocks := (endTime - startTime) / 12
		if newBlocks > 12 {
			sleepTime = 1
		} else {
			sleepTime = int(12 * (12 - newBlocks))
		}
		time.Sleep(time.Second * time.Duration(newBlocks))
	}
}

// GetBlockByHeight
// 根据区块高度获取区块
func GetBlockByHeight(blockHeight int64) (*utils.Block, error) {
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
func GetTxInfoByHash(block *utils.Block) {
	transactions := block.Transactions
	session := db.SyncConn.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	blockTx := blocks.MakeBlockTx(session)

	// 遍历解析交易
	for i := 0; i < len(transactions); i++ {
		receipt, err := Rpc.TransactionReceipt(transactions[i].Hash)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		// 交易事件处理
		EventHandle(receipt.Logs, receipt.TransactionHash)

		// 插入数据库
		trans, err := Rpc.TransactionByHash(transactions[i].Hash)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		// 添加交易
		btx, err := blockTx.GetTxByHashAndAddress(receipt.TransactionHash, receipt.From, receipt.To)
		db.RollbackSession(session, err)
		if btx == nil {
			log.Info("添加ETH交易")
			err = blockTx.Insert(&blocks.BlockTx{
				TxHash:      receipt.TransactionHash,
				FromAddress: receipt.From,
				ToAddress:   receipt.To,
				BlockHeight: block.Number.ToInt().Int64(),
				BlockHash:   receipt.BlockHash,
				Amount:      trans.Value.ToInt().String(),
				Fee:         fmt.Sprintf("%d", int64(trans.Gas)),
				TxStatus:    fmt.Sprintf("%d", int(receipt.Status)),
				TxTimestamp: time.Unix(int64(block.Timestamp), 0),
			})
			db.RollbackSession(session, err)
		}
	}
	db.RollbackSession(session, session.Commit())
}

func EventHandle(vLog []*utils.Log, hash string) {
	fmt.Println("扫链 交易hash", hash)
	for j := 0; j < len(vLog); j++ {
		vlog := vLog[j]
		//这步是对监听到的DATA数据进行解析
		decodedVData, err := hex.DecodeString(vlog.Data[2:])
		if err != nil {
			log.Fatal("对监听到的DATA数据进行解析，发生错误", err)
			return
		}
		//这步是对监听到的DATA数据进行解析
		implEventByLogs(vlog.Topics, decodedVData, hash, int(vlog.LogIndex))
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

func writeBlockToDB(block *utils.Block) {

	session := db.SyncConn.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	blockInfo := blocks.MakeBlockInfo(session)
	isGet, err := blockInfo.GetTxByHash(block.Hash)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	blockInfo.BlockStatus = 1
	blockInfo.BlockHash = block.Hash
	blockInfo.Height = block.Number.ToInt().Int64()
	blockInfo.Miner = block.Miner
	blockInfo.ReceiptsRoot = block.ReceiptsRoot
	blockInfo.StateRoot = block.StateRoot
	blockInfo.ParentHash = block.ParentHash
	blockInfo.BlockTimestamp = time.Unix(int64(block.Timestamp), 0)
	blockInfo.Transactions = len(block.Transactions)
	if isGet {
		err = blockInfo.UpdateBlockInfo()
	} else {
		err = blockInfo.Insert()
	}
	db.RollbackSession(session, err)
	// add Commit() after all actions
	err = session.Commit()
	db.RollbackSession(session, err)
}
