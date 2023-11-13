package blocks

import (
	"time"
	"wallet-analysis/common/db"
	"xorm.io/xorm"
)

type BlockLogs struct {
	Id          int           `xorm:"not null pk autoincr INT"`
	Address     string        `xorm:"VARCHAR(255)"`
	Topics      string        `xorm:"TEXT"`
	Data        string        `xorm:"VARCHAR(255)"`
	BlockHash   string        `xorm:"VARCHAR(255)"`
	BlockNumber int           `xorm:"INT"`
	LogIndex    int           `xorm:"INT"`
	CreatedAt   time.Time     `xorm:"TIMESTAMP"`
	UpdatedAt   time.Time     `xorm:"TIMESTAMP"`
	DeletedAt   time.Time     `xorm:"TIMESTAMP"`
	Session     *xorm.Session `xorm:"-"`
}

func (b *BlockLogs) TableName() string {
	return "block_logs"
}

func init() {
	err := db.ConnDB().Sync(new(BlockLogs))
	if err != nil {
		panic(err)
	}
}

func MakeBlockEvents(session *xorm.Session) (b *BlockLogs) {
	b = new(BlockLogs)
	if session != nil {
		b.Session = session
	} else {
		b.Session = db.SyncConn.NewSession()
	}
	return b
}

func (b *BlockLogs) Insert() error {
	_, err := b.Session.Insert(b)
	if err != nil {
		return err
	}
	return nil
}

func (b *BlockLogs) GetLogs() (*BlockLogs, error) {
	nb := new(BlockLogs)
	_, err := b.Session.Where("block_hash =? and log_index =? and address=?", b.BlockHash, b.BlockHash, b.Address).Get(nb)
	if err != nil {
		return nil, err
	}
	return nb, nil
}
