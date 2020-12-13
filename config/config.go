package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Env      string `mapstructure:"env"`
	AppName  string `mapstructure:"app_name"`
	HttpPort string `mapstructure:"http_port"`
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
