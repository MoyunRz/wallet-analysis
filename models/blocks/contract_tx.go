package blocks

import (
	"time"
	"wallet-analysis/common/db"
	"wallet-analysis/common/log"
	"xorm.io/xorm"
)

type ContractTx struct {
	Id            int64         `xorm:"pk autoincr BIGINT"`
	TxHash        string        `xorm:"VARCHAR(255)"`
	ContractId    int           `xorm:"INT"`
	ContractEvent string        `xorm:"VARCHAR(255)"`
	FromAddress   string        `xorm:"VARCHAR(255)"`
	ToAddress     string        `xorm:"VARCHAR(255)"`
	TokenId       string        `xorm:"VARCHAR(255)"`
	Amount        string        `xorm:"not null default 0.00 decimal(40,18)"`
	LogIndex      int           `xorm:"INT"`
	TxNonce       int           `xorm:"INT"`
	CreatedAt     time.Time     `xorm:"TIMESTAMP"`
	UpdatedAt     time.Time     `xorm:"TIMESTAMP"`
	DeletedAt     time.Time     `xorm:"TIMESTAMP"`
	ExtraData     string        `xorm:"TEXT"`
	Session       *xorm.Session `xorm:"-"`
}

func (c *ContractTx) TableName() string {
	return "contract_tx"
}

func MakeContractTx(session *xorm.Session) (c *ContractTx) {
	c = new(ContractTx)
	if session != nil {
		c.Session = session
	} else {
		c.Session = db.SyncConn.NewSession()
	}
	return c
}

func (c *ContractTx) Insert(ctx *ContractTx) error {
	log.Info(ctx)
	_, err := c.Session.Insert(ctx)
	if err != nil {
		return err
	}
	return nil
}

// GetTxByHash
// 根据 txhash 或者 用户地址 或者 区块高度获取交易
func (c *ContractTx) GetTxByHash(query string) ([]ContractTx, error) {

	blockList := make([]ContractTx, 0)
	querySql := db.SyncConn.Where("tx_hash=?", query)
	err := querySql.Find(&blockList)
	if err != nil {
		return nil, err
	}
	return blockList, nil
}

// GetTxByHashOrAddressOrHeight
// 根据 txhash 或者 用户地址 或者 区块高度获取交易
func (c *ContractTx) GetTxByHashOrAddressOrHeight(query string, limit, start int) (int64, []ContractTx, error) {

	blockList := make([]ContractTx, 0)
	querySql := db.SyncConn.Where("tx_hash=? or from_address=? or to_address=?", query, query, query)
	total, err := querySql.Count(c)
	if err != nil {
		return 0, nil, err
	}
	err = querySql.Limit(limit, start).Find(&blockList)
	if err != nil {
		return 0, nil, err
	}
	return total, blockList, nil
}

// GetTxByHashAndAddress
// 根据 txhash\用户地址\tokenId获取交易
func (c *ContractTx) GetTxByHashAndAddress(txHash, from, to string, tokenId, logIndex int64) (*ContractTx, error) {

	blockList := ContractTx{}
	isGet, err := db.SyncConn.Where("tx_hash=? and from_address=? and to_address=? and token_id=? and log_index =?",
		txHash, from, to, tokenId, logIndex).Get(&blockList)
	if err != nil {
		return nil, err
	}
	if isGet {
		log.Info("查询成功")
		return &blockList, nil
	} else {
		log.Info("查询失败")
		return nil, nil
	}
}
