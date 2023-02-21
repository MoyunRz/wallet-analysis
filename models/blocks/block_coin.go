package blocks

import (
	"time"
	"wallet-analysis/common/db"
	"xorm.io/xorm"
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
	Session          *xorm.Session
}

func (b *BlockCoin) TableName() string {
	return "block_coin"
}

func MakeBlockCoin(session *xorm.Session) (b *BlockCoin) {
	b = new(BlockCoin)
	if session != nil {
		b.Session = session
	} else {
		b.Session = db.SyncConn.NewSession()
	}
	return b
}

func (b *BlockCoin) Insert() error {
	_, err := b.Session.Insert(b)
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
