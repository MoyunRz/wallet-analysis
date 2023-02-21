package service

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
	"wallet-analysis/common/conf"
	"wallet-analysis/common/db"
	"wallet-analysis/common/log"
	"wallet-analysis/models/blocks"
	"wallet-analysis/utils"
	"xorm.io/xorm"
)

var Rpc *utils.RpcClient

func init() {
	Rpc = utils.NewRpcClient(conf.Cfg.Host)
}

func ScanBlock() {
	scanHeight := int64(0)
	blockInfo := new(blocks.BlockInfo)
	// 获取数据库的区块高度
	if conf.Cfg.IsReStartScan {
		scanHeight = conf.Cfg.StartHeight
	} else {
		err := blockInfo.GetMaxHeight()
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		scanHeight = blockInfo.Height
	}

	lasterHeight, err := Rpc.BlockNumber()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	// 循环解析
	for i := scanHeight; i < (lasterHeight - 12); i++ {
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
		rollbackSession(session, err)
		btx.Session = session
		if btx == nil {
			// 插入
			err = btx.Insert(&blocks.BlockTx{
				TxHash:      receipt.TransactionHash,
				FromAddress: receipt.From,
				ToAddress:   receipt.To,
				BlockHeight: block.Number.ToInt().Int64(),
				BlockHash:   receipt.BlockHash,
				Amount:      trans.Value.ToInt().String(),
				Fee:         trans.Gas.String(),
				TxStatus:    receipt.Status.String(),
				TxTimestamp: time.Unix(int64(block.Timestamp), 0),
			})
			rollbackSession(session, err)
		}
	}
	err = session.Commit()
	rollbackSession(session, err)
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
		implEventByLogs(vlog.Topics, decodedVData, int(vlog.LogIndex))
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
	rollbackSession(session, err)
	// add Commit() after all actions
	err = session.Commit()
	rollbackSession(session, err)
}

func rollbackSession(session *xorm.Session, err error) {

	if err != nil {
		err = session.Rollback()
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		log.Fatal(err.Error())
		return
	}
}
