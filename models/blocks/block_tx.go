package blocks

import (
	"time"
	"wallet-analysis/common/db"
	"xorm.io/xorm"
)

type BlockTx struct {
	Id          int64     `xorm:"pk autoincr BIGINT"`
	TxHash      string    `xorm:"VARCHAR(255)"`
	FromAddress string    `xorm:"VARCHAR(255)"`
	ToAddress   string    `xorm:"VARCHAR(255)"`
	BlockHeight int64     `xorm:"INT"`
	BlockHash   string    `xorm:"VARCHAR(255)"`
	Amount      string    `xorm:"DECIMAL(20,18)"`
	Fee         string    `xorm:"DECIMAL(20,18)"`
	TxStatus    string    `xorm:"VARCHAR(255)"`
	TxTimestamp time.Time `xorm:"TIMESTAMP"`
	CreatedAt   time.Time `xorm:"TIMESTAMP"`
	DeletedAt   time.Time `xorm:"TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"TIMESTAMP"`
	Session     *xorm.Session
}

func (b *BlockTx) TableName() string {
	return "block_tx"
}

func MakeBlockTx(session *xorm.Session) (b *BlockTx) {
	b = new(BlockTx)
	if session != nil {
		b.Session = session
	} else {
		b.Session = db.SyncConn.NewSession()
	}
	return b
}

func (b *BlockTx) Insert(newTx *BlockTx) error {
	_, err := b.Session.Insert(newTx)
	if err != nil {
		return err
	}
	return nil
}

// UpdateBlockTx
// 更新区块信息
func (b *BlockTx) UpdateBlockTx() error {
	_, err := b.Session.Where("id=? ", b.Id).Update(b)

	if err != nil {
		return err
	}
	return nil
}

// GetTxByHash
// 根据 txhash 或者 用户地址 或者 区块高度获取交易
func (b *BlockTx) GetTxByHash(query string) ([]BlockTx, error) {

	blockList := make([]BlockTx, 0)
	querySql := db.SyncConn.Where("tx_hash=?", query)
	err := querySql.Find(&blockList)
	if err != nil {
		return nil, err
	}
	return blockList, nil
}

// GetTxByHashOrAddressOrHeight
// 根据 txhash 或者 用户地址 或者 区块高度获取交易
func (b *BlockTx) GetTxByHashOrAddressOrHeight(query string, height, limit, start int) (int64, []BlockTx, error) {

	blockList := make([]BlockTx, 0)
	querySql := db.SyncConn.Where("tx_hash=? or block_hash=? or from_address=? or to_address=? or block_height =?", query, query, query, query, height)
	total, err := querySql.Count(b)
	if err != nil {
		return 0, nil, err
	}
	err = querySql.Limit(limit, start).Find(&blockList)
	if err != nil {
		return 0, nil, err
	}
	return total, blockList, nil
}

// GetTxByHashAndAddress
// 根据 txhash 或者 用户地址 或者 区块高度获取交易
func (b *BlockTx) GetTxByHashAndAddress(txHash string, from, to string) (*BlockTx, error) {
	blockTx := new(BlockTx)
	_, err := db.SyncConn.Where("tx_hash=? and from_address=? and to_address=?", txHash, from, to).Get(blockTx)
	if err != nil {
		return nil, err
	}
	return blockTx, nil
}
