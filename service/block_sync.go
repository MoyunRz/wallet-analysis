package service

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
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
	// 设6秒一个块
	// (结束的时间 - 处理的开始时间)/12 = 出块的个数
	// 出块的个数 > 12 则不休眠
	// 出块的个数 < 12 则：(12 - 出块的个数) * 12 = 休眠时间
	sleepTime := 1
	scanHeight := int64(1)
	blockInfo := blocks.MakeBlockInfo(nil)
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
		if blockInfo.Height != 0 {
			scanHeight = blockInfo.Height
		}
		log.Info("根据数据库高度来扫描 高度 ===> ", scanHeight)
	}

	for {
		endHeight := getNewHeight()
		scanNum := endHeight - 12
		// 循环解析
		for ; scanHeight < scanNum; scanHeight++ {
			log.Infof("当前扫描高度:%d,截止扫描的高度:%d", scanHeight, endHeight-12)
			// 根据区块高度获取区块
			block, err := GetBlockByHeight(scanHeight)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
			if len(block.Transactions) != 0 {
				GetTxInfoByHash(block)
			}
			// 处理完成 写入数据库
			writeBlockToDB(block)
		}
		newBlocks := getNewHeight() - scanHeight
		if newBlocks >= 12 {
			sleepTime = 12
		} else {
			sleepTime = int(4 * (12 - newBlocks))
		}
		log.Info("休眠计算：", newBlocks)
		time.Sleep(time.Second * time.Duration(sleepTime))
	}
}

func getNewHeight() int64 {
	newHeight, err := Rpc.BlockNumber()
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return newHeight
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
	var wg sync.WaitGroup
	ch := make(chan struct{}, 6)
	// 遍历解析交易
	for i := 0; i < len(transactions); i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(height int) {
			// 处理区块内的交易
			defer wg.Done()
			receipt, err := Rpc.TransactionReceipt(transactions[height].Hash)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
			// 交易事件处理
			EventHandle(receipt)
			// 插入数据库
			trans, err := Rpc.TransactionByHash(transactions[height].Hash)
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
					TxTimestamp: time.Unix(block.Timestamp.ToInt().Int64(), 0),
				})
				db.RollbackSession(session, err)
				UpdateETHAssets([]string{receipt.From, receipt.To})
			}
			log.Info("协程 Sleep 0.5 秒")
			time.Sleep(time.Duration(500))
			<-ch
		}(i)
	}
	wg.Wait()
	db.RollbackSession(session, session.Commit())
	time.Sleep(time.Duration(500))
}

func EventHandle(receipts *utils.TransactionReceipt) {
	fmt.Println("扫链 交易hash", receipts.BlockHash)
	for j := 0; j < len(receipts.Logs); j++ {
		vlog := receipts.Logs[j]
		jsonData, err := json.Marshal(vlog.Topics)
		if err != nil {
			fmt.Println("转换为JSON时出错:", err)
			return
		}
		makeBlockLogs := blocks.MakeBlockEvents(nil)
		makeBlockLogs.Address = vlog.Address
		makeBlockLogs.BlockNumber = int(receipts.BlockNumber)
		makeBlockLogs.BlockHash = receipts.BlockHash
		makeBlockLogs.LogIndex = int(vlog.LogIndex)
		makeBlockLogs.Topics = string(jsonData)
		makeBlockLogs.Data = vlog.Data[2:]
		//这步是对监听到的DATA数据进行解析
		implEventByLogs(makeBlockLogs)
	}
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
	isGet, err := blockInfo.GetBlockByHash(block.Hash)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	blockTimestamp := time.Unix(block.Timestamp.ToInt().Int64(), 0)
	blockInfo.BlockStatus = 1
	blockInfo.BlockHash = block.Hash
	blockInfo.Height = block.Number.ToInt().Int64()
	blockInfo.Miner = block.Miner
	blockInfo.ReceiptsRoot = block.ReceiptsRoot
	blockInfo.StateRoot = block.StateRoot
	blockInfo.ParentHash = block.ParentHash
	blockInfo.BlockTimestamp = blockTimestamp
	blockInfo.Transactions = len(block.Transactions)
	if isGet {
		err = blockInfo.UpdateBlockInfo()
	} else {
		err = blockInfo.Insert()
	}
	db.RollbackSession(session, err)
	db.RollbackSession(session, session.Commit())
}

func UpdateETHAssets(addrList []string) {
	session := db.SyncConn.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	assets := blocks.MakeAssets(session)
	isUp := false
	for i := 0; i < len(addrList); i++ {
		if !common.IsHexAddress(addrList[i]) || strings.Contains(addrList[i], "0x0000000000000000000000000000000000000000") {
			continue
		}
		isUp = false
		balance, err := GetEthBalance(addrList[i])
		if err != nil {
			log.Error("查询资产失败")
			log.Fatal(err.Error())
			return
		}
		asset, err := assets.GetAssetsByAddr(addrList[i])
		if err != nil {
			log.Error("查询资产失败")
			log.Fatal(err.Error())
			return
		}
		asset.Address = addrList[i]
		asset.ContractId = 0
		asset.TokenNums = balance.String()
		if asset.Id == 0 {
			err = assets.Insert(asset)
		} else {
			err = assets.UpdateAssets(asset)
		}
		db.RollbackSession(session, err)
		isUp = true
	}
	if isUp {
		db.RollbackSession(session, session.Commit())
	}
}

func GetEthBalance(addr string) (decimal.Decimal, error) {
	return Rpc.EthBalanceByAddress(addr)
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
