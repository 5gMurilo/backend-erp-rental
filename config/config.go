package config

import (
	"github.com/spf13/viper"
)

func InitEnvironment() error {
	viper.SetConfigName("devConfig")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
