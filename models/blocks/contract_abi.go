package blocks

import (
	"time"
	"wallet-analysis/common/db"
	"xorm.io/xorm"
)

type ContractAbi struct {
	Id         int           `xorm:"not null pk autoincr INT"`
	AbiAddress string        `xorm:"VARCHAR(255)"`
	AbiCode    string        `xorm:"TEXT"`
	AbiJson    string        `xorm:"TEXT"`
	CreatedAt  time.Time     `xorm:"TIMESTAMP"`
	UpdatedAt  time.Time     `xorm:"TIMESTAMP"`
	DeletedAt  time.Time     `xorm:"TIMESTAMP"`
	Session    *xorm.Session `xorm:"-"`
}

func (b *ContractAbi) TableName() string {
	return "contract_abi"
}

func init() {
	err := db.ConnDB().Sync(new(ContractAbi))
	if err != nil {
		panic(err)
	}
}

func MakeContractAbi(session *xorm.Session) (b *ContractAbi) {
	b = new(ContractAbi)
	if session != nil {
		b.Session = session
	} else {
		b.Session = db.SyncConn.NewSession()
	}
	return b
}

func (b *ContractAbi) Insert() error {
	_, err := b.Session.Insert(b)
	if err != nil {
		return err
	}
	return nil
}

func (b *ContractAbi) GetAbis() (*ContractAbi, error) {
	nb := new(ContractAbi)
	_, err := b.Session.Where("abi_address =? ", b.AbiAddress).Get(nb)
	if err != nil {
		return nil, err
	}
	return nb, nil
}
