package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet-analysis/models/blocks"
	"wallet-analysis/models/requests"
	"wallet-analysis/models/responses"
)

func resultOk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
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
	if c.ShouldBindQuery(&p) != nil {
		resultError(c, 400, "参数错误")
		return
	}
	makeBlockInfo := blocks.MakeBlockInfo(nil)
	hash, err := makeBlockInfo.GetBlockByHash(p.BlockHash)
	if err != nil {
		resultError(c, 500, "查询出错,请稍后再试")
		return
	}

	if hash {
		// 查询成功

	}

}

// QueryBlockTxByHash
// 根据交易Hash查询交易详情
func QueryBlockTxByHash(c *gin.Context) {
	var p requests.TxQuery
	if c.ShouldBindQuery(&p) != nil {
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
	if c.ShouldBindQuery(&p) != nil {
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
	var p requests.TxQuery
	if c.ShouldBindQuery(&p) != nil {
		resultError(c, 400, "参数错误")
		return
	}

}
