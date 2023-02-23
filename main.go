package main

import (
	"fmt"
	"wallet-analysis/common/conf"
	"wallet-analysis/router"
)

func main() {
	//go func() {
	//	service.ScanBlock()
	//}()
	//go func() {
	//	service.StartSubscribe()
	//}()
	r := router.InitRouters("eth")
	err := r.Run(fmt.Sprintf(":%d", conf.Cfg.ServerPort))
	if err != nil {
		return
	}
}
