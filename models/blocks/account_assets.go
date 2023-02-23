package blocks

import (
	"time"
	"wallet-analysis/common/db"
	"xorm.io/xorm"
)

type AccountAssets struct {
	Id         int64         `xorm:"pk autoincr BIGINT"`
	ContractId int           `xorm:"INT"`
	Address    string        `xorm:"VARCHAR(255)"`
	TokenId    string        `xorm:"VARCHAR(255)"`
	TokenUrl   string        `xorm:"VARCHAR(255)"`
	TokenNums  string        `xorm:"not null default 0.00 decimal(40,18)"`
	CreatedAt  time.Time     `xorm:"TIMESTAMP"`
	UpdatedAt  time.Time     `xorm:"TIMESTAMP"`
	DeletedAt  time.Time     `xorm:"TIMESTAMP"`
	Session    *xorm.Session `xorm:"-"`
}

func (a *AccountAssets) TableName() string {
	return "account_assets"
}

func MakeAssets(session *xorm.Session) (a *AccountAssets) {
	a = new(AccountAssets)
	if session != nil {
		a.Session = session
	} else {
		a.Session = db.SyncConn.NewSession()
	}
	return a
}

// Insert
// 插入
func (a *AccountAssets) Insert() error {
	_, err := a.Session.Insert(a)
	if err != nil {
		return err
	}
	return nil
}

// UpdateAssets
// 更新用户资产
func (a *AccountAssets) UpdateAssets() error {
	_, err := a.Session.Where("id=? ", a.Id).Update(a)

	if err != nil {
		return err
	}
	return nil
}

func (a *AccountAssets) GetAssets(cid int, tokenId, address string) error {
	_, err := db.SyncConn.Where("contract_id=? and token_id=? and address =?", cid, tokenId, address).Get(a)
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountAssets) GetAccountAllAssets() ([]AccountAssets, error) {
	assets := make([]AccountAssets, 0)
	err := db.SyncConn.Where("address =?", a.Address).Find(&assets)
	if err != nil {
		return assets, err
	}
	return assets, nil
}

func (a *AccountAssets) FindAllTokenByAddress(addr, contract string) ([]AccountAssets, error) {
	assets := make([]AccountAssets, 0)

	err := db.SyncConn.
		Table("account_assets").
		Join("LEFT OUTER", "block_token", "block_token.id = account_assets.contract_id").
		Where("account_assets.address =? and block_token.contract_address =?", addr, contract).
		Find(&assets)
	if err != nil {
		return assets, err
	}
	return assets, nil
}
