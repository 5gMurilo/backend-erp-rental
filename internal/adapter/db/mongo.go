package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ManagerWorker struct {
	Client   *mongo.Client
	Database string
}

func newManagerWorker(c *mongo.Client, database string) *ManagerWorker {
	return &ManagerWorker{
		Client:   c,
		Database: database,
	}
}

func (mw *ManagerWorker) GetCollection(collectionName string) *mongo.Collection {
	return mw.Client.Database(mw.Database).Collection(collectionName)
}

func (mw *ManagerWorker) StartSession() (mongo.Session, error) {
	if session, err := mw.Client.StartSession(); err != nil {
		log.Println(err)
		return nil, err
	} else {
		return session, nil
	}
}

func InitDb(c context.Context, uri string, database string) (*ManagerWorker, error) {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to mongodb!")

	result := newManagerWorker(client, database)
	return result, nil
}
