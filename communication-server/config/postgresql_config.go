package config

import "github.com/spf13/viper"

type PostgreSQL struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	PollSize int    `mapstructure:"poolSize"`
}

var _ configLoader = (*PostgreSQL)(nil)

func (PostgreSQL) loadDefault() {
	viper.SetDefault("postgres", map[string]interface{}{
		"host":     "postgres",
		"port":     5432,
		"username": "dgame",
		"password": "dgame",
		"database": "dgame",
		"pollSize": 5,
	})
}
