package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Address string
	}
	Database struct {
		DSN string
	}
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
