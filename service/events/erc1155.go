package events

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"wallet-analysis/common/db"
	"wallet-analysis/common/log"
	"wallet-analysis/models/blocks"
	"wallet-analysis/service/abis"
	"wallet-analysis/utils"
)

var contractAddress = ""
var contractId = 0

func init() {

	blockCoin := new(blocks.BlockToken)
	token, err := blockCoin.GetToken("ERC1155")
	if err != nil {
		log.Fatal(err)
		return
	}
	contractAddress = token.ContractAddress
	contractId = token.Id
}

func GetIndexedAddress(topics string) string {
	return strings.Replace(topics, "0x000000000000000000000000", "0x", 1)
}

func Update1155Assets(addrList []string, tokenAddress string, tokenId int64) {
	contract := common.HexToAddress(tokenAddress)
	makeClient, err := utils.MakeClient()
	if err != nil {
		log.Error("链接 eth socket error", err)
		log.Fatal(err.Error())
		return
	}
	defer makeClient.Close()
	//创建合约对象
	xunWenGeToken, err := abis.NewXunWenGe(contract, makeClient)
	if err != nil {
		log.Error("NewXunWenGe error", err)
		log.Fatal(err.Error())
		return
	}
	session := db.SyncConn.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	assets := blocks.MakeAssets(session)
	tId := fmt.Sprintf("%d", tokenId)
	isUp := false
	for i := 0; i < len(addrList); i++ {
		isUp = false
		balance, err := xunWenGeToken.BalanceOf(&bind.CallOpts{}, common.HexToAddress(addrList[i]), big.NewInt(tokenId))
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Balance: %s,toBalance: %s\n", balance)

		err = assets.GetAssets(contractId, tId, addrList[i])
		if err != nil {
			log.Error("查询资产失败")
			log.Fatal(err.Error())
			return
		}
		assets.ContractId = contractId
		assets.Address = addrList[i]
		assets.TokenId = tId
		assets.TokenNums = balance.String()
		if assets.Id == 0 {
			err = assets.Insert()
		} else {
			err = assets.UpdateAssets()
		}
		db.RollbackSession(session, err)
		isUp = true
	}
	if isUp {
		db.RollbackSession(session, session.Commit())
	}
}

// UpdateMintTx
// 更新合约铸造交易
func UpdateMintTx(txHash string, intr []interface{}, txIndex int) {
	session := db.SyncConn.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	makeContractTx := blocks.MakeContractTx(session)
	addrList := intr[0].([]common.Address)
	tokenIds := intr[1].([]*big.Int)
	amounts := intr[2].([]*big.Int)
	// txType := intr[3].(*big.Int)
	datas := intr[3].(uint8)
	// 合并相同交易组
	txMap := mergingTx(addrList, tokenIds, amounts)
	isUp := false
	for addr, v := range txMap {
		log.Info(addr, "铸造")
		for tokenId, amount := range v {
			log.Info("铸造", tokenId, amount)
			tx, err := makeContractTx.GetTxByHashAndAddress(
				txHash,
				"0x0000000000000000000000000000000000000000000000000000000000000000",
				addr,
				tokenId,
				int64(txIndex),
			)
			db.RollbackSession(session, err)
			if tx == nil {
				log.Info("Insert 铸造")
				err = makeContractTx.Insert(&blocks.ContractTx{
					TxHash:        txHash,
					ContractId:    contractId,
					ContractEvent: "MintLog",
					FromAddress:   "0x0000000000000000000000000000000000000000000000000000000000000000",
					ToAddress:     addr,
					TokenId:       fmt.Sprintf("%d", tokenId),
					Amount:        fmt.Sprintf("%d", amount),
					LogIndex:      txIndex,
					ExtraData:     fmt.Sprintf("%x", datas),
				})
				db.RollbackSession(session, err)
				isUp = true
			}
			// 更新用户资产
			go func(addr string, tid int64) {
				Update1155Assets([]string{addr}, contractAddress, tid)
			}(addr, tokenId)
		}
	}
	if isUp {
		// 交易提交
		log.Info("交易提交")
		db.RollbackSession(session, session.Commit())
	}
}

// UpdateTransferLogTx
// 更新合约单笔转账交易
func UpdateTransferLogTx(txHash string, intr []interface{}, txIndex int) {
	session := db.SyncConn.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	from := intr[0].(common.Address)
	addrList := intr[1].([]common.Address)
	tokenIds := intr[2].([]*big.Int)
	amounts := intr[3].([]*big.Int)
	// txType := intr[3].(*big.Int)
	datas := intr[4]

	// 合并相同交易组
	txMap := mergingTx(addrList, tokenIds, amounts)
	makeContractTx := blocks.MakeContractTx(session)

	isUp := false
	for addr, v := range txMap {
		for tokenId, amount := range v {
			tx, err := makeContractTx.GetTxByHashAndAddress(
				txHash,
				from.String(),
				addr,
				tokenId,
				int64(txIndex),
			)
			db.RollbackSession(session, err)
			if tx == nil {
				err = makeContractTx.Insert(&blocks.ContractTx{
					TxHash:        txHash,
					ContractId:    contractId,
					ContractEvent: "TransferLog",
					FromAddress:   from.String(),
					ToAddress:     addr,
					TokenId:       fmt.Sprintf("%d", tokenId),
					Amount:        fmt.Sprintf("%d", amount),
					LogIndex:      txIndex,
					ExtraData:     fmt.Sprintf("%x", datas),
				})
				db.RollbackSession(session, err)
				isUp = true
			}
			// 更新用户资产
			go func(addr string, tid int64) {
				Update1155Assets([]string{addr}, contractAddress, tid)
			}(addr, tokenId)
		}
	}
	if isUp {
		db.RollbackSession(session, session.Commit())
	}
}

// UpdateBurnLogTx
// 更新合约单笔转账交易
func UpdateBurnLogTx(txHash string, intr []interface{}, txIndex int) {

	session := db.SyncConn.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	from := intr[0].(common.Address)
	tokenIds := intr[1].([]*big.Int)
	amounts := intr[3].([]*big.Int)
	// txType := intr[3].(*big.Int)

	// 合并相同交易组
	txMap := mergingTx([]common.Address{from}, tokenIds, amounts)
	makeContractTx := blocks.MakeContractTx(session)

	isUp := false
	for addr, v := range txMap {
		log.Info("销毁 token")
		for tokenId, amount := range v {
			tx, err := makeContractTx.GetTxByHashAndAddress(
				txHash,
				from.String(),
				"0x0000000000000000000000000000000000000000000000000000000000000000",
				tokenId,
				int64(txIndex),
			)
			db.RollbackSession(session, err)
			if tx == nil {
				err = makeContractTx.Insert(&blocks.ContractTx{
					TxHash:        txHash,
					ContractId:    contractId,
					ContractEvent: "BurnLog",
					FromAddress:   from.String(),
					ToAddress:     "0x0000000000000000000000000000000000000000000000000000000000000000",
					TokenId:       fmt.Sprintf("%d", tokenId),
					Amount:        fmt.Sprintf("%d", amount),
					LogIndex:      txIndex,
					ExtraData:     "",
				})
				db.RollbackSession(session, err)
				isUp = true
			}
			// 更新用户资产
			go func(addr string, tid int64) {
				Update1155Assets([]string{addr}, contractAddress, tid)
			}(addr, tokenId)
		}
	}
	if isUp {
		db.RollbackSession(session, session.Commit())
	}
}

// UpdateTransferSingleTx
// 更新合约单笔转账交易
func UpdateTransferSingleTx(txHash string, intr []interface{}, txIndex int) {
	session := db.SyncConn.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	from := intr[1].(string)
	to := intr[2].(string)
	tokenId := intr[3].(*big.Int)
	amount := intr[4].(*big.Int)
	// 合并相同交易组
	makeContractTx := blocks.MakeContractTx(session)

	isUp := false
	tx, err := makeContractTx.GetTxByHashAndAddress(
		txHash,
		from,
		to,
		tokenId.Int64(),
		int64(txIndex),
	)
	db.RollbackSession(session, err)
	if tx == nil {
		err = makeContractTx.Insert(&blocks.ContractTx{
			TxHash:        txHash,
			ContractId:    contractId,
			ContractEvent: "TransferSingle",
			FromAddress:   from,
			ToAddress:     to,
			TokenId:       fmt.Sprintf("%d", tokenId),
			Amount:        fmt.Sprintf("%d", amount),
			LogIndex:      txIndex,
			ExtraData:     "",
		})
		db.RollbackSession(session, err)
		isUp = true
	}
	if isUp {
		db.RollbackSession(session, session.Commit())
	}
	// 更新用户资产
	go func(lst []string, tid int64) {
		Update1155Assets(lst, contractAddress, tid)
	}([]string{from, to}, tokenId.Int64())
}

// UpdateTransferBatchTx
// 更新合约多笔转账交易
func UpdateTransferBatchTx(txHash string, intr []interface{}, txIndex int) {
	session := db.SyncConn.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	from := intr[1].(string)
	to := intr[2].(string)
	tokenIds := intr[3].([]*big.Int)
	amounts := intr[4].([]*big.Int)
	makeContractTx := blocks.MakeContractTx(session)
	isUp := false
	for i := 0; i < len(tokenIds); i++ {
		tx, err := makeContractTx.GetTxByHashAndAddress(
			txHash,
			from,
			to,
			tokenIds[i].Int64(),
			int64(txIndex),
		)
		db.RollbackSession(session, err)
		if tx == nil {
			err = makeContractTx.Insert(&blocks.ContractTx{
				TxHash:        txHash,
				ContractId:    contractId,
				ContractEvent: "TransferBatch",
				FromAddress:   from,
				ToAddress:     to,
				TokenId:       fmt.Sprintf("%d", tokenIds[i]),
				Amount:        fmt.Sprintf("%d", amounts[i]),
				LogIndex:      txIndex,
				ExtraData:     "",
			})
			db.RollbackSession(session, err)
			isUp = true
		}
		if isUp {
			db.RollbackSession(session, session.Commit())
		}
		// 更新用户资产
		go func(lst []string, tid int64) {
			Update1155Assets(lst, contractAddress, tid)
		}([]string{from, to}, tokenIds[i].Int64())
	}
}

func mergingTx(addrList []common.Address, tokenIds, amounts []*big.Int) map[string]map[int64]int64 {
	txMap := map[string]map[int64]int64{}
	log.Info("合并相同交易组")
	// 合并相同交易组
	for i := 0; i < len(addrList); i++ {
		key := addrList[i].String()
		t := txMap[key]
		if t != nil {
			tm := t[tokenIds[i].Int64()]
			if tm != 0 {
				tm += amounts[i].Int64()
				var n = map[int64]int64{}
				n[tokenIds[i].Int64()] = tm
				txMap[key] = n
			}
		} else {
			var n = map[int64]int64{}
			n[tokenIds[i].Int64()] = amounts[i].Int64()
			txMap[key] = n
		}
	}
	log.Info("合并相同交易组 txMap", txMap)
	return txMap
}
