package blocks

import (
	"time"
	"wallet-analysis/common/db"
)

type AccountAssets struct {
	Id         int64     `xorm:"pk autoincr BIGINT"`
	ContractId int       `xorm:"INT"`
	Address    string    `xorm:"VARCHAR(255)"`
	TokenId    string    `xorm:"VARCHAR(255)"`
	TokenNums  int64     `xorm:"BIGINT"`
	CreatedAt  time.Time `xorm:"TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"TIMESTAMP"`
	DeletedAt  time.Time `xorm:"TIMESTAMP"`
}

func (a *AccountAssets) TableName() string {
	return "account_assets"
}

// Insert
// 插入
func (a *AccountAssets) Insert() error {
	_, err := db.SyncConn.Insert(a)
	if err != nil {
		return err
	}
	return nil
}

// UpdateAssets
// 更新用户资产
func (a *AccountAssets) UpdateAssets() error {
	_, err := db.SyncConn.Where("id=? ", a.Id).Update(a)

	if err != nil {
		return err
	}
	return nil
}

func (a *AccountAssets) GetAssets(cid, tokenId int, address string) error {
	_, err := db.SyncConn.
		Where("contract_id=? and token_id and address =?", cid, tokenId, address).
		Get(a)
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountAssets) GetAcoountAllAssets() ([]AccountAssets, error) {
	assets := make([]AccountAssets, 0)
	err := db.SyncConn.Where("address =?", a.Address).Find(&assets)
	if err != nil {
		return assets, err
	}
	return assets, nil
}
