package router

import (
	"github.com/gin-gonic/gin"
	"wallet-analysis/common/middlewares"
	"wallet-analysis/controller"
)

func InitRouters(name string) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	group := r.Group(name)
	blockGroup(group)
	return r
}

func blockGroup(group *gin.RouterGroup) {
	g := group.Group("query")
	// 根据区块Hash查询区块
	g.POST("blockByHash", controller.QueryBlockByHash)
	// 根据交易Hash查询交易详情
	g.POST("blockTxByHash", controller.QueryBlockTxByHash)
	// 根据交易Hash查询合约交易详情
	g.POST("contractTxByHash", controller.QueryContractTxByHash)
	// 获取用户的资产Token
	g.POST("userAssetToken", controller.FindUserAssetToken)
	// 获取用户的所有合约
	g.POST("userAssetContract", controller.FindUserAssetContract)
}
