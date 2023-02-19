package conf

import (
	"github.com/BurntSushi/toml"
)

var Cfg = new(Config)

var defConfFile = "app.toml"

func init() {
	if err := LoadConfig(defConfFile, Cfg); err != nil {

	}
}

type Config struct {
	Mode        string         `toml:"mode"`
	Host        string         `toml:"host"`
	ChainId     int            `toml:"chainId"`
	StartHeight int            `toml:"startHeight"`
	Log         LogConfig      `toml:"log"`
	DataBase    DatabaseConfig `toml:"database"`
}

type LogConfig struct {
	Level     string `toml:"level"       json:"level"`
	Formatter string `toml:"formatter"   json:"formatter"`
	OutFile   string `toml:"outfile"    json:"outfile"`
	ErrFile   string `toml:"errfile"    json:"errfile"`
}

type DatabaseConfig struct {
	Name     string `toml:"name"`
	Type     string `toml:"type"`
	Url      string `toml:"url"`
	User     string `toml:"user"`
	PassWord string `toml:"password"`
	Mode     string `toml:"mode"`
}

// 从相对路径Load conf
// 请传入指针类型
func LoadConfig(cfgPath string, cfg *Config) error {

	if cfgPath == "" {
		cfgPath = defConfFile
	}

	if _, err := toml.DecodeFile(cfgPath, cfg); err != nil {
		return err
	}
	return nil
}
