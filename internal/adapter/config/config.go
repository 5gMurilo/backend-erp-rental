package config

import (
	"github.com/spf13/viper"
)

type (
	Container struct {
		Onedrive *Onedrive
		Mongo    *Mongo
	}

	Onedrive struct {
		ObjectId string
		ClientId string
		Username string
		Password string
	}

	Mongo struct {
		Uri      string
		Database string
	}
)

func New() (*Container, error) {

	viper.SetConfigName("app")
	//viper.AddConfigPath("./src")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	onedrive := &Onedrive{
		ObjectId: viper.GetString("ONEDRIVE_OBJECT_ID"),
		ClientId: viper.GetString("ONEDRIVE_CLIENT_ID"),
		Username: viper.GetString("ONEDRIVE_USERNAME"),
		Password: viper.GetString("ONEDRIVE_PASSWORD"),
	}

	mongo := &Mongo{
		Uri:      viper.GetString("MONGO_URI"),
		Database: viper.GetString("MONGO_DATABASE"),
	}

	return &Container{
		Onedrive: onedrive,
		Mongo:    mongo,
	}, nil
}
