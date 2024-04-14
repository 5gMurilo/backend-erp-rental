package main

import (
	"america-rental-backend/internal/adapter/Auth"
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/adapter/db/repository"
	"america-rental-backend/internal/adapter/http"
	"america-rental-backend/internal/adapter/http/handler"
	"america-rental-backend/internal/adapter/http/middleware"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/service"
	"context"
)

func main() {
	token, err := Auth.New()
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()
	worker, err := db.InitDb(ctx)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepositoryImpl(worker)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userActivitiesRepo := repository.NewEmployeeActivityLog(worker)
	userActivitiesService := service.NewEmployeeActivityLogService(userActivitiesRepo)
	userActivitiesHandler := handler.NewEmployeeActivityLogHandler(userActivitiesService)

	authService := service.NewAuthService(userRepo, token)
	authHandler := handler.NewAuthHandler(authService)
	authMiddleware := middleware.NewAuthMiddleware(token)

	storageService := service.NewStorageService(nil, domain.StorageAuthentication{
		ObjectId: "eec5618c-fa22-42c3-b827-7c5d5443f6ae",
		ClientId: "a66e44a4-0dd5-4b44-aed0-881ea37799b4",
		Username: "financeiroerh@AmericaRental932.onmicrosoft.com",
		Password: "Bj0_20285412",
		Scopes:   []string{"https://graph.microsoft.com/.default"},
	})
	storageHandler := handler.NewStorageHandler(storageService)

	router := http.Router(userHandler, authHandler, authMiddleware, userActivitiesHandler, storageHandler)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
