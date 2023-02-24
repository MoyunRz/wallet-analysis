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
	Id            int64  `json:"id"`
	TxHash        string `json:"tx_hash"`
	ContractId    int    `json:"contract_id"`
	ContractEvent string `json:"contract_event"`
	FromAddress   string `json:"from_address"`
	ToAddress     string `json:"to_address"`
	TokenId       string `json:"token_id"`
	Amount        string `json:"amount"`
	LogIndex      int    `json:"log_index"`
	TxNonce       int    `json:"tx_nonce"`
	ExtraData     string `json:"extra_data"`
}

type UserAsset struct {
	Address         string `json:"address"`
	ContractAddress string `json:"contract_address"`
	TokenId         string `json:"token_id"`
	TokenName       string `json:"token_name"`
	TokenNums       string `json:"token_nums"`
	TokenUrl        string `json:"token_url"`
}

type ContractToken struct {
	Id               int    `json:"address"`
	ContractAddress  string `json:"contract_address"`
	ContractName     string `json:"contract_name"`
	ContractFullName string `json:"contract_full_name"`
	ContractType     string `json:"contract_type"`
}

type TxResponse struct {
	Transaction         []ETHTransaction `json:"eth_transaction"`
	ContractTransaction []ContractInfo   `json:"contract_transaction"`
}

type BlockInfo struct {
	Id             int64     `json:"id"`
	Height         int64     `json:"height"`
	BlockHash      string    `json:"block_hash"`
	Miner          string    `json:"miner"`
	ParentHash     string    `json:"parent_hash"`
	ReceiptsRoot   string    `json:"receipts_root"`
	StateRoot      string    `json:"state_root"`
	BlockStatus    int       `json:"block_status"`
	BlockTimestamp time.Time `json:"block_timestamp"`
	Transactions   int       `json:"transactions"`
}
