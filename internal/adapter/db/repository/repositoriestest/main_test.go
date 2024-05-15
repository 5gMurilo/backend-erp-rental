package repositoriestest

import (
	"america-rental-backend/internal/adapter/db"
	"context"
	"fmt"
	"testing"
)

var (
	databaseName = "test"
	worker       *db.ManagerWorker
)

func TestMain(m *testing.M) {
	worker, _ = db.InitDb(context.Background(), "mongodb://localhost:27017/", databaseName)

	setup()
	m.Run()
}

func setup() {
	createCollections()
}

func createCollections() {
	err := worker.Client.Database(worker.Database).CreateCollection(context.Background(), "contract")
	if err != nil {
		fmt.Printf("error creating collection: %s", err.Error())
	}
}
