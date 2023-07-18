package configs

import (
	"github.com/spf13/viper"
)

type Configs struct {
	Port 	 string `mapstructure:"PORT"`
	PgDriverName string `mapstructure:"PG_DRIVER"`
	PgDataSourceName string `mapstructure:"PG_SOURCE"`
}

func LoadConfig(path string) (conf Configs, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return
	}

	return
}