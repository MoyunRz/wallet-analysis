package service

import (
	"wallet-analysis/common/conf"
	"wallet-analysis/models/blocks"
	"wallet-analysis/utils"
)

func init() {
	utils.InitClient()
}

func implEventByLogs(makeBlockLogs *blocks.BlockLogs) {
	v, e := makeBlockLogs.GetLogs()
	if e != nil {
		return
	}
	if v != nil && v.Id != 0 {
		return
	}

	err := makeBlockLogs.Insert()
	if err != nil {
		return
	}
	makeAbi := blocks.MakeContractAbi(nil)
	makeAbi.AbiAddress = makeBlockLogs.Address
	abis, err := makeAbi.GetAbis()
	if err != nil {
		return
	}
	if abis.Id == 0 {
		// save abi
		codes, err := utils.NewRpcClient(conf.Cfg.Host).GetCode(makeBlockLogs.Address)
		if err != nil {
			return
		}

		if codes != "" {
			makeAbi = blocks.MakeContractAbi(nil)
			makeAbi.AbiAddress = makeBlockLogs.Address
			makeAbi.AbiCode = codes
			err = makeAbi.Insert()
			if err != nil {
				return
			}
		}
	}
}
