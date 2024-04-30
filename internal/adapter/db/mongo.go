package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func (mw *ManagerWorker) StartSession() (mongo.Session, error) {
	if session, err := mw.client.StartSession(); err != nil {
		log.Println(err)
		return nil, err
	} else {
		return session, nil
	}
}

func InitDb(c context.Context) (*ManagerWorker, error) {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	//uri := "mongodb://americarentaldb:Bjc20285412@host.docker.internal:27017/"
	uri := "mongodb://localhost:27017/"
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

	result := newManagerWorker(client)
	return result, nil
}
