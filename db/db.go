package db

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() error {

	return nil
}

type ManagerWorker struct {
	client *mongo.Client
}

func newManagerWorker(c *mongo.Client) *ManagerWorker {
	return &ManagerWorker{
		client: c,
	}
}

func (mw *ManagerWorker) GetCollection(collectionName string) *mongo.Collection {
	return mw.client.Database("america").Collection(collectionName)
}

func InitDb() *ManagerWorker {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	uri := viper.GetString("mongo_uri")

	options := options.Client().ApplyURI(uri)
	defer cancel()

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connected to mongodb!")

	result := newManagerWorker(client)
	return result
}
