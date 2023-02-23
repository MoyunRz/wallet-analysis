package requests

type BlockQuery struct {
	BlockHash string `json:"block_hash" form:"block_hash"`
	Height    int    `json:"height" form:"height"`
	From      string `json:"from" form:"from"`
	To        string `json:"to" form:"to"`
	PageNum   int    `json:"page_num" form:"page_num"`
	PageSize  int    `json:"page_size" form:"page_size"`
}
type TxQuery struct {
	Query    string `json:"query" form:"query"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}
