package blocks

import (
	"time"
	"wallet-analysis/common/db"
)

type BlockCoin struct {
	Id               int       `xorm:"not null pk autoincr INT"`
	ContractAddress  string    `xorm:"VARCHAR(255)"`
	ContractName     string    `xorm:"VARCHAR(255)"`
	ContractFullName string    `xorm:"VARCHAR(255)"`
	ContractType     string    `xorm:"VARCHAR(255)"`
	CreatedAt        time.Time `xorm:"TIMESTAMP"`
	UpdatedAt        time.Time `xorm:"TIMESTAMP"`
	DeletedAt        time.Time `xorm:"TIMESTAMP"`
}

func (b *BlockCoin) TableName() string {
	return "block_coin"
}

func (b *BlockCoin) Insert() error {
	_, err := db.SyncConn.Insert(b)
	if err != nil {
		return err
	}
	return nil
}

func (b *BlockCoin) FindAll() ([]BlockCoin, error) {
	coinList := make([]BlockCoin, 0)
	err := db.SyncConn.Find(&coinList)
	if err != nil {
		return nil, err
	}
	return coinList, nil
}
