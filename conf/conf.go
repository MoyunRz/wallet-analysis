package conf

/*
使用toml进行配置
*/

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
)

var (
	Config *tomlConfig
	once   sync.Once
)

const (
	ADDRESS_NIL_ERR   = "address is null error"
	Valid_ADDRESS_ERR = "Parse valid address post data error"
)

func InitConfig() {
	once.Do(func() {
		filePath, err := filepath.Abs("./app.toml")
		if err != nil {
			// log.Error(err)
			panic(err)
		}
		if _, err := toml.DecodeFile(filePath, &Config); err != nil {
			// log.Errorf("Read config file error,Err=[%v]",err)
			panic(fmt.Sprintf("Read config file error,Err=[%v]", err))
		}
	})
}

type tomlConfig struct {
	Host        string `toml:"host"`
	ChainId     int    `toml:"chainId"`
	StartHeight int    `toml:"startHeight"`
}
