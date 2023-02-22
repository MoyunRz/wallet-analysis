package blocks

import (
	"time"
	"wallet-analysis/common/db"
	"xorm.io/xorm"
)

type BlockToken struct {
	Id               int           `xorm:"not null pk autoincr INT"`
	ContractAddress  string        `xorm:"VARCHAR(255)"`
	ContractName     string        `xorm:"VARCHAR(255)"`
	ContractFullName string        `xorm:"VARCHAR(255)"`
	ContractType     string        `xorm:"VARCHAR(255)"`
	CreatedAt        time.Time     `xorm:"TIMESTAMP"`
	UpdatedAt        time.Time     `xorm:"TIMESTAMP"`
	DeletedAt        time.Time     `xorm:"TIMESTAMP"`
	Session          *xorm.Session `xorm:"-"`
}

func (b *BlockToken) TableName() string {
	return "block_token"
}

func MakeBlockToken(session *xorm.Session) (b *BlockToken) {
	b = new(BlockToken)
	if session != nil {
		b.Session = session
	} else {
		b.Session = db.SyncConn.NewSession()
	}
	return b
}

func (b *BlockToken) Insert() error {
	_, err := b.Session.Insert(b)
	if err != nil {
		return err
	}
	return nil
}

func (b *BlockToken) FindAll() ([]BlockToken, error) {
	coinList := make([]BlockToken, 0)
	err := db.SyncConn.Find(&coinList)
	if err != nil {
		return nil, err
	}
	return coinList, nil
}

func (b *BlockToken) GetToken(tokenType string) (*BlockToken, error) {
	blockToken := new(BlockToken)
	_, err := db.SyncConn.Where("contract_type=?", tokenType).Get(&blockToken)
	if err != nil {
		return nil, err
	}
	return blockToken, nil
}
