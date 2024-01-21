package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Http HTTP `mapstructure:"http"`
}

type HTTP struct {
	Port string `mapstructure:"port"`
}

var config *Config

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("config not found")
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("unmarshal config error")
	}
	viper.AutomaticEnv()

	return config
}
