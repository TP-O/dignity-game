package config

import "github.com/spf13/viper"

type Mail struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

var _ configLoader = (*Mail)(nil)

func (Mail) loadDefault() {
	viper.SetDefault("mail", map[string]interface{}{
		//
	})
}
