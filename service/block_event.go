package service

import (
	"wallet-analysis/common/log"
	"wallet-analysis/service/events"
	"wallet-analysis/utils"
)

func init() {
	utils.InitClient()
}
func implEventByLogs(topics []string, decodedVData []byte, hash string, txIndex int) {
	switchErc1155Event(topics, decodedVData, hash, txIndex)
}

func switchErc1155Event(topics []string, decodedVData []byte, hash string, txIndex int) {
	eventAbi := utils.EventAbi
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
		// 进行事件处理
		events.UpdateMintTx(hash, intr, txIndex)
		break
	// 交易事件 TransferLog
	case "0x832711906223d7b1424466041e692f503b3467cdb4d5dbc5f746adfc531da26d":
		log.Info("交易事件 event TransferLog")
		//这步是对监听到的DATA数据进行解析
		intr, err := eventAbi.Events["TransferLog"].Inputs.UnpackValues(decodedVData)
		if err != nil {
			log.Fatal(err)
		}
		// 进行事件处理
		events.UpdateTransferLogTx(hash, intr, txIndex)
		break
	// 销毁事件 BurnLog
	case "0x3a9276528c9c1b7064f942560fe085b661a55400887092ba3bc7063d492d5545":
		log.Info("销毁事件 event BurnLog")
		//这步是对监听到的DATA数据进行解析
		intr, err := eventAbi.Events["BurnLog"].Inputs.UnpackValues(decodedVData)
		if err != nil {
			log.Fatal(err)
		}
		// 进行事件处理
		events.UpdateBurnLogTx(hash, intr, txIndex)
		break
	// 单体转账 TransferSingle
	case "0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62":
		//这步是对监听到的DATA数据进行解析
		intr, err := eventAbi.Events["TransferSingle"].Inputs.UnpackValues(decodedVData)
		if err != nil {
			log.Fatal(err)
		}
		var list []interface{}
		list = append(list, events.GetIndexedAddress(topics[1]), events.GetIndexedAddress(topics[2]), events.GetIndexedAddress(topics[3]))
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
		list = append(list, events.GetIndexedAddress(topics[2]), events.GetIndexedAddress(topics[3]))
		for _, v := range intr {
			list = append(list, v)
		}
		//打印监听到的参数
		log.Info(list)
		break
	}
}
