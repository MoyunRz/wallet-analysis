package db

import (
	"errors"
	"fmt"
	"time"
	"wallet-analysis/common/conf"
	"wallet-analysis/common/log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func init() {
	err := InitSyncDB2(conf.Cfg.DataBase)
	if err != nil {
		panic(err.Error())
	}
}

var SyncConn *xorm.Engine

// InitSyncDB2 连接数据库
func InitSyncDB2(cfg conf.DatabaseConfig) error {
	dburl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", cfg.User, cfg.PassWord, cfg.Url, cfg.Name)
	conn, err := initDBConn(cfg.Type, dburl)
	if err != nil {
		return err
	}
	SyncConn = conn
	return nil
}

func initDBConn(dbType, dbUrl string) (coin *xorm.Engine, err error) {
	if dbUrl == "" || dbType == "" {
		return nil, errors.New("empty databases config")
	}
	conn, err := xorm.NewEngine(dbType, dbUrl)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(2)
	conn.SetMaxOpenConns(6)
	conn.SetConnMaxLifetime(60 * time.Second)
	//conn.ShowSQL(true)
	//conn.ShowExecTime(true)
	return conn, nil
}

func RollbackSession(session *xorm.Session, err error) {

	if err != nil {

		err1 := session.Rollback()
		if err1 != nil {
			log.Fatal(err1.Error())
			return
		}
		log.Fatal(err.Error())
		return
	}
}
