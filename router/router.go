package router

import (
	"github.com/gin-gonic/gin"
	"wallet-analysis/controller"
)

func InitRouters(group *gin.RouterGroup) {
	group.POST("/query_blockTxByHash", controller.QueryBlockTxByHash) // 校验地址
}
