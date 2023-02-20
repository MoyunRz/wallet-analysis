package blocks

import (
	"time"
)

type BlockInfo struct {
	Id             int64     `xorm:"pk autoincr BIGINT"`
	Height         int64     `xorm:"not null BIGINT"`
	BlockHash      string    `xorm:"VARCHAR(255)"`
	Miner          string    `xorm:"VARCHAR(255)"`
	ParentHash     string    `xorm:"VARCHAR(255)"`
	ReceiptsRoot   string    `xorm:"VARCHAR(255)"`
	StateRoot      string    `xorm:"VARCHAR(255)"`
	BlockStatus    int       `xorm:"INT"`
	NextBlockHash  int       `xorm:"INT"`
	BlockTimestamp time.Time `xorm:"TIMESTAMP"`
	Transactions   int       `xorm:"INT"`
	CreatedAt      time.Time `xorm:"TIMESTAMP"`
	UpdatedAt      time.Time `xorm:"TIMESTAMP"`
	DeletedAt      time.Time `xorm:"TIMESTAMP"`
}

func (b *BlockInfo) TableName() string {
	return "block_info"
}
