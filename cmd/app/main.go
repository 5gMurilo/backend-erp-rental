package main

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/adapter/db/repository"
	"america-rental-backend/internal/adapter/http"
	"america-rental-backend/internal/adapter/http/handler"
	"america-rental-backend/internal/core/service"
	"context"
)

func main() {
	ctx := context.TODO()
	worker, err := db.InitDb(ctx)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepositoryImpl(worker)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := http.Router(userHandler)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
