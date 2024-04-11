package main

import (
	"america-rental-backend/internal/adapter/Auth"
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/adapter/db/repository"
	"america-rental-backend/internal/adapter/http"
	"america-rental-backend/internal/adapter/http/handler"
	"america-rental-backend/internal/adapter/http/middleware"
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

	router := http.Router(userHandler, authHandler, authMiddleware, userActivitiesHandler)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
