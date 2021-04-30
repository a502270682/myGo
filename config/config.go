package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Env      string `mapstructure:"env"`
	AppName  string `mapstructure:"app_name"`
	HttpPort string `mapstructure:"http_port"`
	Mysql    Mysql  `mapstructure:"mysql"`
}

type Mysql struct {
	Master ConnectionConfig `mapstructure:"master"`
	Slave  ConnectionConfig `mapstructure:"slave"`
}

type ConnectionConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       string `mapstructure:"db"`
	MaxIdle  int    `mapstructure:"max_idle"`
	MaxOpen  int    `mapstructure:"max_open"`
	Debug    bool   `mapstructure:"debug"`
}

var gConfig Config

func Load(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to load conf")
	}

	if err := v.Unmarshal(&gConfig); err != nil {
		return nil, errors.Wrap(err, "fail to Unmarshal conf")
	}

	return &gConfig, nil
}
func GetConfig() *Config {
	return &gConfig
}
