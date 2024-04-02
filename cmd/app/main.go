package main

import (
	"america-rental-backend/adapters/db"
	"context"
)

func main() {
	ctx := context.TODO()
	_, err := db.InitDb(ctx)
	if err != nil {
		panic(err)
	}

}
