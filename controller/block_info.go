package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet-analysis/models/blocks"
	"wallet-analysis/models/requests"
	"wallet-analysis/models/responses"
)

func resultOk(c *gin.Context, res interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"msg":    "success",
		"result": res,
	})
}

func resultMapOk(c *gin.Context, resMap map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"msg":    "success",
		"result": resMap,
	})
}

func resultError(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg":  msg,
	})
}

// QueryBlockByHash
// 根据区块Hash查询区块
func QueryBlockByHash(c *gin.Context) {
	var p requests.BlockQuery
	if c.ShouldBindJSON(&p) != nil {
		resultError(c, 400, "参数错误")
		return
	}
	makeBlockInfo := blocks.MakeBlockInfo(nil)
	isGet, err := makeBlockInfo.GetBlockByHash(p.BlockHash)
	if err != nil {
		resultError(c, 500, "查询出错,请稍后再试")
		return
	}
	if !isGet {
		// 查询成功
		resultOk(c, nil)
		return
	}

	resultOk(c, responses.BlockInfo{
		Id:             makeBlockInfo.Id,
		Height:         makeBlockInfo.Height,
		BlockHash:      makeBlockInfo.BlockHash,
		Miner:          makeBlockInfo.Miner,
		ParentHash:     makeBlockInfo.ParentHash,
		ReceiptsRoot:   makeBlockInfo.ReceiptsRoot,
		StateRoot:      makeBlockInfo.StateRoot,
		BlockStatus:    makeBlockInfo.BlockStatus,
		BlockTimestamp: makeBlockInfo.BlockTimestamp,
		Transactions:   makeBlockInfo.Transactions,
	})
	return
}

// QueryBlockTxByHash
// 根据交易Hash查询交易详情
func QueryBlockTxByHash(c *gin.Context) {
	var p requests.TxQuery
	if c.ShouldBindJSON(&p) != nil {
		resultError(c, 400, "参数错误")
		return
	}
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize <= 0 {
		p.PageNum = 10
	}
	makeBlockTx := blocks.MakeBlockTx(nil)
	makeContractTx := blocks.MakeContractTx(nil)
	txResponseList := make([]responses.TxResponse, 0)

	// 获取普通交易
	total, blockTxList, err := makeBlockTx.GetTxByHashOrAddressOrHeight(p.Query, p.PageSize, p.PageNum)
	if err != nil {
		resultError(c, 500, "查询出错,请稍后再试")
		return
	}

	for i := 0; i < len(blockTxList); i++ {
		txResponse := responses.TxResponse{}
		txResponse.Transaction = responses.ETHTransaction{
			Id:          blockTxList[i].Id,
			TxHash:      blockTxList[i].TxHash,
			FromAddress: blockTxList[i].FromAddress,
			ToAddress:   blockTxList[i].ToAddress,
			BlockHeight: blockTxList[i].BlockHeight,
			BlockHash:   blockTxList[i].BlockHash,
			Amount:      blockTxList[i].Amount,
			Fee:         blockTxList[i].Fee,
			TxStatus:    blockTxList[i].TxStatus,
			TxTimestamp: blockTxList[i].TxTimestamp,
		}
		// 	获取合约交易
		cTxList, err := makeContractTx.GetTxByHash(blockTxList[i].TxHash)
		if err != nil {
			return
		}
		for j := 0; j < len(cTxList); j++ {
			txResponse.ContractTransaction = append(txResponse.ContractTransaction, responses.ContractInfo{
				Id:            cTxList[j].Id,
				TxHash:        cTxList[j].TxHash,
				ContractId:    cTxList[j].ContractId,
				ContractEvent: cTxList[j].ContractEvent,
				FromAddress:   cTxList[j].FromAddress,
				ToAddress:     cTxList[j].ToAddress,
				TokenId:       cTxList[j].TokenId,
				Amount:        cTxList[j].Amount,
				LogIndex:      cTxList[j].LogIndex,
				TxNonce:       cTxList[j].TxNonce,
				ExtraData:     cTxList[j].ExtraData,
			})
		}
		txResponseList = append(txResponseList, txResponse)
	}
	resultMapOk(c, map[string]interface{}{
		"list":  txResponseList,
		"total": total,
	})
	return
}

// QueryContractTxByHash
// 根据交易Hash查询合约交易详情
func QueryContractTxByHash(c *gin.Context) {
	var p requests.TxQuery
	if c.ShouldBindJSON(&p) != nil {
		resultError(c, 400, "参数错误")
		return
	}
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize <= 0 {
		p.PageNum = 10
	}
	makeContractTx := blocks.MakeContractTx(nil)

	// 	获取合约交易
	total, cTxList, err := makeContractTx.GetTxByHashOrAddressOrHeight(p.Query, p.PageSize, p.PageNum)
	if err != nil {
		return
	}
	contractList := make([]responses.ContractInfo, 0)
	for j := 0; j < len(cTxList); j++ {
		contractList = append(contractList, responses.ContractInfo{
			Id:            cTxList[j].Id,
			TxHash:        cTxList[j].TxHash,
			ContractId:    cTxList[j].ContractId,
			ContractEvent: cTxList[j].ContractEvent,
			FromAddress:   cTxList[j].FromAddress,
			ToAddress:     cTxList[j].ToAddress,
			TokenId:       cTxList[j].TokenId,
			Amount:        cTxList[j].Amount,
			LogIndex:      cTxList[j].LogIndex,
			TxNonce:       cTxList[j].TxNonce,
			ExtraData:     cTxList[j].ExtraData,
		})
	}
	resultMapOk(c, map[string]interface{}{
		"list":  contractList,
		"total": total,
	})
	return
}

// FindUserAssetToken
// 获取用户的资产Token
func FindUserAssetToken(c *gin.Context) {
	var p requests.AssertsQuery
	if c.ShouldBindJSON(&p) != nil {
		resultError(c, 400, "参数错误")
		return
	}
	// 根据from 地址查询余额
	makeAssets := blocks.MakeAssets(nil)
	allTokens, err := makeAssets.FindAllTokenByAddress(p.Address, p.ContractAddress)
	if err != nil {
		resultError(c, 500, "查询错误")
		return
	}
	assetList := make([]responses.UserAsset, 0)
	for i := 0; i < len(allTokens); i++ {
		assetList = append(assetList, responses.UserAsset{
			Address:         allTokens[i].Address,
			ContractAddress: p.ContractAddress,
			TokenId:         allTokens[i].TokenId,
			TokenNums:       allTokens[i].TokenNums,
			TokenUrl:        allTokens[i].TokenUrl,
		})
	}
	resultMapOk(c, map[string]interface{}{
		"list":  assetList,
		"total": len(assetList),
	})
	return
}

// FindUserAssetContract
// 获取用户的所有合约
func FindUserAssetContract(c *gin.Context) {
	var p requests.AssertsQuery
	if c.ShouldBindJSON(&p) != nil {
		resultError(c, 400, "参数错误")
		return
	}
	// 根据from 地址查询余额
	makeAssets := blocks.MakeBlockToken(nil)
	allTokens, err := makeAssets.FindAllContractByAddress(p.Address)
	if err != nil {
		resultError(c, 500, "查询错误")
		return
	}
	tokenList := make([]responses.ContractToken, 0)

	for i := 0; i < len(allTokens); i++ {
		tokenList = append(tokenList, responses.ContractToken{
			Id:               allTokens[i].Id,
			ContractAddress:  p.ContractAddress,
			ContractName:     allTokens[i].ContractAddress,
			ContractFullName: allTokens[i].ContractName,
			ContractType:     allTokens[i].ContractFullName,
		})
	}
	resultMapOk(c, map[string]interface{}{
		"list":  tokenList,
		"total": len(tokenList),
	})
	return
}
