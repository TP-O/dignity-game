package config

import "github.com/spf13/viper"

const (
	DevEnv  = "development"
	ProdEnv = "production"
)

type App struct {
	Debug     bool   `mapstructure:"debug"`
	Env       string `mapstructure:"env"`
	Port      uint16 `mapstructure:"port"`
	SecretKey string `mapstructure:"secretKey"`
	Host      string `mapstructure:"host"`
}

var _ configLoader = (*App)(nil)

func (App) loadDefault() {
	viper.SetDefault("app", map[string]interface{}{
		"debug": true,
		"env":   DevEnv,
		"port":  8080,
		"host":  "http://localhost:8080/",
	})
}
