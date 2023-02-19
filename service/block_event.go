package service

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"wallet-analysis/common/conf"
	"wallet-analysis/common/log"
	"wallet-analysis/utils"
)

func implEventByLogs(vlog utils.Log) {
	eventAbi := conf.EventAbi
	switch vlog.Topics[0] {
	//铸造事件 MintLog
	case "0x87b8ba4f1ba2e813af31d438ace9cf4fa3f0e82e86b679cd044ae1b07276c9c5":
		//这步是对监听到的DATA数据进行解析
		decodedVData, err := hex.DecodeString(vlog.Data[2:])
		if err != nil {
			log.Fatal(err)
		}
		intr, err := eventAbi.Events["MintLog"].Inputs.UnpackValues(decodedVData)
		if err != nil {
			log.Fatal(err)
		}
		//打印监听到的参数
		fmt.Println(intr[0].([]common.Address), intr[1].([]*big.Int), intr[2], intr[3], intr[4])
		break
	// 交易事件 TransferLog
	case "0x832711906223d7b1424466041e692f503b3467cdb4d5dbc5f746adfc531da26d":
		//这步是对监听到的DATA数据进行解析
		intr, err := eventAbi.Events["TransferLog"].Inputs.UnpackValues([]byte(vlog.Data))
		if err != nil {
			log.Fatal(err)
		}
		list := []interface{}{}
		for _, v := range intr {
			list = append(list, v)
		}
		//打印监听到的参数
		fmt.Println(list)
		break
	// 销毁事件 BurnLog
	case "0x3a9276528c9c1b7064f942560fe085b661a55400887092ba3bc7063d492d5545":

		//这步是对监听到的DATA数据进行解析
		intr, err := eventAbi.Events["BurnLog"].Inputs.UnpackValues([]byte(vlog.Data))
		if err != nil {
			log.Fatal(err)
		}
		list := []interface{}{}
		for _, v := range intr {
			list = append(list, v)
		}
		//打印监听到的参数
		fmt.Println(list)
		break
	}
}
