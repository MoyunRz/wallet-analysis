package db

import (
	"errors"
	"fmt"
	"time"
	"wallet-analysis/common/conf"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func init() {
	if err := InitSyncDB2(conf.Cfg.DataBase); err != nil {
		panic(err.Error())
	}
}

var (
	SyncConn *xorm.Engine
	UserConn *xorm.Engine
)

// InitSyncDB2 连接数据库
func InitSyncDB2(cfg conf.DatabaseConfig) error {
	dburl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", cfg.User, cfg.PassWord, cfg.Url, cfg.Name)
	conn, err := initDBConn(cfg.Type, dburl, cfg.Mode)
	if err != nil {
		return err
	}
	SyncConn = conn
	return nil
}

// InitUserDB2 连接数据库
func InitUserDB2(cfg conf.DatabaseConfig) error {
	dburl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", cfg.User, cfg.PassWord, cfg.Url, cfg.Name)
	conn, err := initDBConn(cfg.Type, dburl, cfg.Mode)
	if err != nil {
		return err
	}
	UserConn = conn
	return nil
}
func initDBConn(dbType, dbUrl, dbMode string) (coin *xorm.Engine, err error) {
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
