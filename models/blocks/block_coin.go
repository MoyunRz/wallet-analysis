package blocks

import (
	"time"
)

type BlockCoin struct {
	Id               int       `xorm:"not null pk autoincr INT"`
	ContractAddress  string    `xorm:"VARCHAR(255)"`
	ContractName     string    `xorm:"VARCHAR(255)"`
	ContractFullName string    `xorm:"VARCHAR(255)"`
	CreatedAt        time.Time `xorm:"TIMESTAMP"`
	UpdatedAt        time.Time `xorm:"TIMESTAMP"`
	DeletedAt        time.Time `xorm:"TIMESTAMP"`
}

func (a *BlockCoin) TableName() string {
	return "block_coin"
}
