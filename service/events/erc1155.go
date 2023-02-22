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
	"xorm.io/xorm"
)

var contractAddress = ""

func init() {
	blockCoin := new(blocks.BlockToken)
	token, err := blockCoin.GetToken("ERC1155")
	if err != nil {
		log.Fatal(err)
		return
	}
	contractAddress = token.ContractAddress
}

func GetIndexedAddress(topics string) string {
	return strings.Replace(topics, "0x000000000000000000000000", "0x", 1)
}

func Update1155Assets(account, contract string, tokenId int64) {
	tokenAddress := common.HexToAddress(contract)
	//创建合约对象
	xunWenGeToken, err := abis.NewXunWenGe(tokenAddress, utils.EthClient)
	if err != nil {
		fmt.Println("newChttoken error", err)
	}

	bal, err := xunWenGeToken.BalanceOf(&bind.CallOpts{}, common.HexToAddress(account), big.NewInt(tokenId))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("balance: %s\n", bal) // "wei: 74605500647408739782407023"
}

// UpdateMintTx
// 更新合约铸造交易
func UpdateMintTx(session *xorm.Session, txHash string, intr []interface{}, txIndex int) {
	makeContractTx := blocks.MakeContractTx(session)
	addrList := intr[0].([]common.Address)
	tokenIds := intr[1].([]*big.Int)
	amounts := intr[2].([]*big.Int)
	// txType := intr[3].(*big.Int)
	datas := intr[3].(uint8)
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
			db.RollbackSession(session, err)
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
					ExtraData:     fmt.Sprintf("%s", datas),
				})
				db.RollbackSession(session, err)
			}
			// 更新用户资产
			go func(addr string, tid int64) {
				Update1155Assets(addr, contractAddress, tid)
			}(addr, tokenId)
		}
	}
}

// UpdateTransferSingleTx
// 更新合约单笔转账交易
func UpdateTransferSingleTx(session *xorm.Session, txHash string, list []interface{}, txIndex int) {
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
	db.RollbackSession(session, err)

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
		db.RollbackSession(session, err)
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
