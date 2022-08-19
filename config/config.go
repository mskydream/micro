package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Database DB     `mapstructure:"db"`
	Port     string `mapstructure:"port"`
}

type DB struct {
	Host     string `mapstructure:"host"`
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	Name     string `mapstructure:"name"`
}

func LoadConfig() (config Configuration, err error) {
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	fmt.Println(config)
	return
}
