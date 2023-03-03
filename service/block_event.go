package service

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"wallet-analysis/common/log"
	"wallet-analysis/service/events"
	"wallet-analysis/utils"
)

type EventStore struct {
	Events   abi.Event
	CallFunc func(events.TopicsEvent)
}

var EventMap map[string]EventStore

func init() {
	utils.InitClient()
	EventMap = map[string]EventStore{
		"0x87b8ba4f1ba2e813af31d438ace9cf4fa3f0e82e86b679cd044ae1b07276c9c5": {
			utils.EventAbi.Events["MintLog"],
			events.UpdateMintTx,
		},
		"0x832711906223d7b1424466041e692f503b3467cdb4d5dbc5f746adfc531da26d": {
			utils.EventAbi.Events["TransferLog"],
			events.UpdateTransferLogTx,
		},
		"0x3a9276528c9c1b7064f942560fe085b661a55400887092ba3bc7063d492d5545": {
			utils.EventAbi.Events["BurnLog"],
			events.UpdateBurnLogTx,
		},
		"0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62": {
			utils.EventAbi.Events["TransferSingle"],
			events.UpdateTransferSingleTx,
		},
		"0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb": {
			utils.EventAbi.Events["TransferBatch"],
			events.UpdateTransferBatchTx,
		},
	}
}

func implEventByLogs(topics []string, decodedVData []byte, hash string, txIndex int) {
	switchErc1155Event(topics, decodedVData, hash, txIndex)
}

func switchErc1155Event(topics []string, decodedVData []byte, hash string, txIndex int) {
	topicsEvent, ok := EventMap[topics[0]]
	if ok {
		intr, err := topicsEvent.Events.Inputs.UnpackValues(decodedVData)
		if err != nil {
			log.Fatal(err)
		}
		topicsEvent.CallFunc(events.TopicsEvent{
			TxHash:  hash,
			Intr:    intr,
			TxIndex: txIndex,
			Topics:  topics,
		})
	}
}
