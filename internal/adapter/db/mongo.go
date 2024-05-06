package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ManagerWorker struct {
	client   *mongo.Client
	database string
}

func newManagerWorker(c *mongo.Client, database string) *ManagerWorker {
	return &ManagerWorker{
		client:   c,
		database: database,
	}
}

func (mw *ManagerWorker) GetCollection(collectionName string) *mongo.Collection {
	return mw.client.Database(mw.database).Collection(collectionName)
}

func (mw *ManagerWorker) StartSession() (mongo.Session, error) {
	if session, err := mw.client.StartSession(); err != nil {
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
