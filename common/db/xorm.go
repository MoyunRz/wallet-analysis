package db

import (
	"errors"
	"fmt"
	"os"
	"time"
	"wallet-analysis/common/conf"
	"wallet-analysis/common/log"
	"xorm.io/xorm/caches"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

func init() {

	err := InitSyncDB2(conf.UseDataBase)
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

	cacher := caches.NewLRUCacher2(caches.NewMemoryStore(), time.Minute, 100000)
	conn, err := xorm.NewEngine(dbType, dbUrl)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetDefaultCacher(cacher)
	conn.SetMaxIdleConns(30)
	conn.SetMaxOpenConns(15)
	conn.SetConnMaxLifetime(1 * time.Minute)
	conn.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	f, err := os.Create(conf.Cfg.Log.Sqlfile)
	if err != nil {
		println(err.Error())
		return
	}
	conn.ShowSQL(true)
	conn.SetLogger(xlog.NewSimpleLogger(f))
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
