package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wallet-analysis/common/conf"
	"wallet-analysis/router"
	"wallet-analysis/service"
)

func main() {
	go func() {
		service.ScanBlock()
	}()
	go func() {
		service.StartSubscribe()
	}()
	r := gin.Default()
	router.InitRouters(initHttp(r))
	err := r.Run(fmt.Sprintf(":%d", conf.Cfg.ServerPort))
	if err != nil {
		return
	}
}

func initHttp(defaultGin *gin.Engine) *gin.RouterGroup {
	group := defaultGin.Group("eth")
	return group
}
