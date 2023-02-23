package responses

import (
	"time"
)

type ETHTransaction struct {
	Id          int64     `json:"id"`
	TxHash      string    `json:"tx_hash"`
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	BlockHeight int64     `json:"block_height"`
	BlockHash   string    `json:"block_hash"`
	Amount      string    `json:"amount"`
	Fee         string    `json:"fee"`
	TxStatus    string    `json:"tx_status"`
	TxTimestamp time.Time `json:"tx_timestamp"`
}

type ContractInfo struct {
	Id            int64  `xorm:"pk autoincr BIGINT"`
	TxHash        string `xorm:"VARCHAR(255)"`
	ContractId    int    `xorm:"INT"`
	ContractEvent string `xorm:"VARCHAR(255)"`
	FromAddress   string `xorm:"VARCHAR(255)"`
	ToAddress     string `xorm:"VARCHAR(255)"`
	TokenId       string `xorm:"VARCHAR(255)"`
	Amount        string `xorm:"not null default 0.00 decimal(40,18)"`
	LogIndex      int    `xorm:"INT"`
	TxNonce       int    `xorm:"INT"`
	ExtraData     string `xorm:"TEXT"`
}

type TxResponse struct {
	Transaction         ETHTransaction `json:"transaction"`
	ContractTransaction []ContractInfo `json:"contract_transaction"`
}
