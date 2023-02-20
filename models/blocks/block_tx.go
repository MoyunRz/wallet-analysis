package blocks

import (
	"time"
	"wallet-analysis/common/db"
)

type BlockTx struct {
	Id          int64     `xorm:"pk autoincr BIGINT"`
	TxHash      string    `xorm:"VARCHAR(255)"`
	FromAddress string    `xorm:"VARCHAR(255)"`
	ToAddress   string    `xorm:"VARCHAR(255)"`
	BlockHeight int       `xorm:"INT"`
	BlockHash   string    `xorm:"VARCHAR(255)"`
	Amount      string    `xorm:"DECIMAL(20,18)"`
	Fee         string    `xorm:"DECIMAL(20,18)"`
	TxStatus    string    `xorm:"VARCHAR(255)"`
	TxTimestamp time.Time `xorm:"TIMESTAMP"`
	CreatedAt   time.Time `xorm:"TIMESTAMP"`
	DeletedAt   time.Time `xorm:"TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"TIMESTAMP"`
}

func (b *BlockTx) TableName() string {
	return "block_tx"
}

func (b *BlockTx) Insert() error {
	_, err := db.SyncConn.Insert(b)
	if err != nil {
		return err
	}
	return nil
}

// GetTxByHashOrAddressOrHeight
// 根据 txhash 或者 用户地址 或者 区块高度获取交易
func (b *BlockTx) GetTxByHashOrAddressOrHeight(query string, height, limit, start int) (int64, []BlockTx, error) {

	blockList := make([]BlockTx, 0)
	querySql := db.SyncConn.Where("tx_hash=? or block_hash=? or from_Address=? or to_Address=? or block_height =?", query, query, query, query, height)
	total, err := querySql.Count(b)
	if err != nil {
		return total, blockList, err
	}
	err = querySql.Limit(limit, start).Find(&blockList)
	if err != nil {
		return total, blockList, err
	}

	return total, blockList, err
}
