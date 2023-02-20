package blocks

import "time"

type AccountAssets struct {
	Id         int64     `xorm:"pk autoincr BIGINT"`
	ContractId int       `xorm:"INT"`
	Address    string    `xorm:"VARCHAR(255)"`
	TokenIds   string    `xorm:"VARCHAR(255)"`
	TokenNums  int64     `xorm:"BIGINT"`
	CreatedAt  time.Time `xorm:"TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"TIMESTAMP"`
	DeletedAt  time.Time `xorm:"TIMESTAMP"`
}

func (a *AccountAssets) TableName() string {
	return "account_assets"
}
