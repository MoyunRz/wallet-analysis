package blocks

import (
	"time"
	"wallet-analysis/common/db"
	"xorm.io/xorm"
)

type BlockInfo struct {
	Id             int64         `xorm:"pk autoincr BIGINT"`
	Height         int64         `xorm:"not null BIGINT"`
	BlockHash      string        `xorm:"VARCHAR(255)"`
	Miner          string        `xorm:"VARCHAR(255)"`
	ParentHash     string        `xorm:"VARCHAR(255)"`
	ReceiptsRoot   string        `xorm:"VARCHAR(255)"`
	StateRoot      string        `xorm:"VARCHAR(255)"`
	BlockStatus    int           `xorm:"INT"`
	BlockTimestamp time.Time     `xorm:"TIMESTAMP"`
	Transactions   int           `xorm:"INT"`
	CreatedAt      time.Time     `xorm:"TIMESTAMP"`
	UpdatedAt      time.Time     `xorm:"TIMESTAMP"`
	DeletedAt      time.Time     `xorm:"TIMESTAMP"`
	Session        *xorm.Session `xorm:"-"`
}

func (b *BlockInfo) TableName() string {
	return "block_info"
}
func init() {
	err := db.ConnDB().Sync(new(BlockInfo))
	if err != nil {
		panic(err)
	}
}
func MakeBlockInfo(session *xorm.Session) (b *BlockInfo) {
	b = new(BlockInfo)
	if session != nil {
		b.Session = session
	} else {
		b.Session = db.SyncConn.NewSession()
	}
	return b
}

func (b *BlockInfo) Insert() error {
	_, err := b.Session.Insert(b)
	if err != nil {
		return err
	}
	return nil
}

// UpdateBlockInfo
// 更新区块信息
func (b *BlockInfo) UpdateBlockInfo() error {
	_, err := b.Session.Where("id=? ", b.Id).Update(b)

	if err != nil {
		return err
	}
	return nil
}

func (b *BlockInfo) GetMaxHeight() error {
	isGet, err := db.SyncConn.Desc("height").Get(b)
	if err != nil {
		return err
	}
	if !isGet {
		b.Height = 0
	}
	return nil
}

// GetBlockByHash
// 根据 txhash 或者 用户地址 或者 区块高度获取交易
func (b *BlockInfo) GetBlockByHash(blockHash string) (bool, error) {
	isGet, err := db.SyncConn.Where("block_hash=?", blockHash).Get(b)

	if err != nil {
		return isGet, err
	}
	return isGet, nil
}

// GetTxByHashOrAddressOrHeight
// 根据 txhash 或者 用户地址 或者 区块高度获取交易
func (b *BlockInfo) GetTxByHashOrAddressOrHeight(blockHash string, height int) (bool, error) {

	return db.SyncConn.Where("block_hash=? or height=? ", blockHash, height).Get(b)
}
