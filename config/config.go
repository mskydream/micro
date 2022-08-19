package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Database DB     `mapstucture:"db"`
	Port     string `mapstructure:"port"`
}

type DB struct {
	Address  string `mapstructure:"address"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

func LoadConfig() (config Configuration, err error) {
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	fmt.Println(config)
	return

}
