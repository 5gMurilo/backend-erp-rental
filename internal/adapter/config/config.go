package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		Onedrive *Onedrive
	}

	Onedrive struct {
		ObjectId string
		ClientId string
		Username string
		Password string
	}
)

func New() (*Container, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	onedrive := &Onedrive{
		ObjectId: os.Getenv("ONEDRIVE_OBJECT_ID"),
		ClientId: os.Getenv("ONEDRIVE_CLIENT_ID"),
		Username: os.Getenv("ONEDRIVE_USERNAME"),
		Password: os.Getenv("ONEDRIVE_PASSWORD"),
	}

	return &Container{
		Onedrive: onedrive,
	}, nil
}
