package blocks

import (
	"time"
)

type ContractTx struct {
	Id            int64     `xorm:"pk autoincr BIGINT"`
	TxHash        string    `xorm:"VARCHAR(255)"`
	ContractId    int       `xorm:"INT"`
	ContractEvent string    `xorm:"VARCHAR(255)"`
	FromAddress   string    `xorm:"VARCHAR(255)"`
	ToAddress     string    `xorm:"VARCHAR(255)"`
	TokenId       string    `xorm:"VARCHAR(255)"`
	Amount        string    `xorm:"DECIMAL(20,18)"`
	CreatedAt     time.Time `xorm:"TIMESTAMP"`
	UpdatedAt     time.Time `xorm:"TIMESTAMP"`
	DeletedAt     time.Time `xorm:"TIMESTAMP"`
	ExtraData     string    `xorm:"TEXT"`
}

func (b *ContractTx) TableName() string {
	return "contract_tx"
}
