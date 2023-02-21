package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"wallet-analysis/common/db"
	"wallet-analysis/common/log"
	"wallet-analysis/models/blocks"
	"wallet-analysis/utils"
	"xorm.io/xorm"
)

func init() {
	utils.InitClient()
}
func implEventByLogs(topics []string, decodedVData []byte, hash string, txIndex int) {
	eventAbi := utils.EventAbi
	session := db.SyncConn.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// 进行事件匹配
	switch topics[0] {
	//铸造事件 MintLog
	case "0x87b8ba4f1ba2e813af31d438ace9cf4fa3f0e82e86b679cd044ae1b07276c9c5":
		log.Info("铸造事件 event MintLog")
		//这步是对监听到的DATA数据进行解析
		intr, err := eventAbi.Events["MintLog"].Inputs.UnpackValues(decodedVData)
		if err != nil {
			log.Fatal(err)
		}
		//打印监听到的参数
		updateMintTx(session, hash, intr, txIndex)
		break
	// 交易事件 TransferLog
	case "0x832711906223d7b1424466041e692f503b3467cdb4d5dbc5f746adfc531da26d":
		log.Info("交易事件 event TransferLog")
		//这步是对监听到的DATA数据进行解析
		intr, err := eventAbi.Events["TransferLog"].Inputs.UnpackValues(decodedVData)
		if err != nil {
			log.Fatal(err)
		}
		var list []interface{}
		for _, v := range intr {
			list = append(list, v)
		}

		//打印监听到的参数
		log.Info(list)
		break
	// 销毁事件 BurnLog
	case "0x3a9276528c9c1b7064f942560fe085b661a55400887092ba3bc7063d492d5545":
		log.Info("销毁事件 event BurnLog")
		//这步是对监听到的DATA数据进行解析
		intr, err := eventAbi.Events["BurnLog"].Inputs.UnpackValues(decodedVData)
		if err != nil {
			log.Fatal(err)
		}
		list := []interface{}{}
		for _, v := range intr {
			list = append(list, v)
		}
		//打印监听到的参数
		log.Info(list)
		break
	// 单体转账 TransferSingle
	case "0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62":
		//这步是对监听到的DATA数据进行解析
		intr, err := eventAbi.Events["TransferSingle"].Inputs.UnpackValues(decodedVData)
		if err != nil {
			log.Fatal(err)
		}
		var list []interface{}
		list = append(list, GetIndexedAddress(topics[1]), GetIndexedAddress(topics[2]), GetIndexedAddress(topics[3]))
		for _, v := range intr {
			list = append(list, v)
		}
		//打印监听到的参数
		log.Info(list)
		break
	// a->b 批量转账 TransferBatch
	case "0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb":
		log.Info("批量转账 event TransferBatch ")
		//这步是对监听到的DATA数据进行解析
		intr, err := eventAbi.Events["TransferBatch"].Inputs.UnpackValues(decodedVData)
		if err != nil {
			log.Fatal(err)
		}
		var list []interface{}
		list = append(list, GetIndexedAddress(topics[1]), GetIndexedAddress(topics[2]), GetIndexedAddress(topics[3]))
		for _, v := range intr {
			list = append(list, v)
		}

		//打印监听到的参数
		log.Info(list)
		break
	}
}

func GetIndexedAddress(topics string) string {
	return strings.Replace(topics, "0x000000000000000000000000", "0x", 1)
}

// updateMintTx
// 更新合约铸造交易
func updateMintTx(session *xorm.Session, txHash string, intr []interface{}, txIndex int) {
	makeContractTx := blocks.MakeContractTx(session)
	addrList := intr[0].([]common.Address)
	tokenIds := intr[1].([]*big.Int)
	amounts := intr[2].([]*big.Int)
	// txType := intr[3].(*big.Int)
	datas := intr[3].([]uint8)
	// 合并相同交易组
	txMap := mergingTx(addrList, tokenIds, amounts)

	for addr, v := range txMap {
		for tokenId, amount := range v {
			tx, err := makeContractTx.GetTxByHashAndAddress(
				txHash,
				"0x0000000000000000000000000000000000000000000000000000000000000000",
				addr,
				tokenId,
				int64(txIndex),
			)
			rollbackSession(session, err)
			if tx == nil {
				err = makeContractTx.Insert(&blocks.ContractTx{
					TxHash:        txHash,
					ContractId:    1,
					ContractEvent: "MintLog",
					FromAddress:   "0x0000000000000000000000000000000000000000000000000000000000000000",
					ToAddress:     addr,
					TokenId:       fmt.Sprintf("%d", tokenId),
					Amount:        fmt.Sprintf("%d", amount),
					LogIndex:      txIndex,
					ExtraData:     fmt.Sprintf("%s", datas[:]),
				})
				rollbackSession(session, err)
			}
		}
	}
}

// updateTransferSingleTx
// 更新合约单笔转账交易
func updateTransferSingleTx(session *xorm.Session, txHash string, list []interface{}, txIndex int) {
	makeContractTx := blocks.MakeContractTx(session)
	//sender := list[0].(string)
	from := list[1].(string)
	to := list[2].(string)
	tokenId := list[3].(int)
	amount := list[3].(int)
	tx, err := makeContractTx.GetTxByHashAndAddress(
		txHash,
		from,
		to,
		int64(tokenId),
		int64(txIndex),
	)
	rollbackSession(session, err)

	if tx == nil {
		err = makeContractTx.Insert(&blocks.ContractTx{
			TxHash:        txHash,
			ContractId:    1,
			ContractEvent: "TransferSingle",
			FromAddress:   from,
			ToAddress:     to,
			TokenId:       fmt.Sprintf("%d", tokenId),
			Amount:        fmt.Sprintf("%d", amount),
			LogIndex:      txIndex,
			ExtraData:     "",
		})
		rollbackSession(session, err)
	}
}

func mergingTx(addrList []common.Address, tokenIds, amounts []*big.Int) map[string]map[int64]int64 {

	txMap := map[string]map[int64]int64{}

	// 合并相同交易组
	for i := 0; i < len(addrList); i++ {
		t := txMap[addrList[i].String()]
		if t != nil {
			tm := t[tokenIds[i].Int64()]
			if tm != 0 {
				tm += amounts[i].Int64()
				var n = map[int64]int64{}
				n[tokenIds[i].Int64()] = tm
				txMap[addrList[i].String()] = n
			}
		} else {
			var n = map[int64]int64{}
			n[tokenIds[i].Int64()] = amounts[i].Int64()
			txMap[addrList[i].String()] = n
		}
	}

	return txMap
}
