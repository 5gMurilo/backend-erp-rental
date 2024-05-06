package main

import (
	"america-rental-backend/internal/adapter/Auth"
	"america-rental-backend/internal/adapter/config"
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

	container, err := config.New()
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()
	worker, err := db.InitDb(ctx, container.Mongo.Uri, container.Mongo.Database)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepositoryImpl(worker)
	employeeRepo := repository.NewEmployeeRepositoryImpl(worker)
	employeeActivitiesRepo := repository.NewEmployeeActivityLog(worker)
	epiRepo := repository.NewEpiRepositoryImpl(worker)
	returnableEpiRepo := repository.NewReturnableEpiRepositoryImpl(worker)
	epiExitsRepo := repository.NewEpiExitsRepositoryImpl(worker)
	storageRepo := repository.NewStorageRepository(worker)

	authService := service.NewAuthService(userRepo, token)
	userService := service.NewUserService(userRepo)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeActivitiesService := service.NewEmployeeActivityLogService(employeeActivitiesRepo, employeeRepo)
	epiService := service.NewEpiService(epiRepo)
	returnableEpiService := service.NewReturnableEpiService(returnableEpiRepo, epiRepo)
	epiExitsService := service.NewEpiExitsService(epiExitsRepo)

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	employeeHandler := handler.NewEmployeeHandler(employeeActivitiesService, employeeService)
	employeeActivitiesHandler := handler.NewEmployeeActivityLogHandler(employeeActivitiesService)
	epiHandler := handler.NewEpiHandler(epiService)
	epiExitsHandler := handler.NewEpiExitsHandler(epiExitsService)
	returnableEpiHandler := handler.NewReturnableEpiHandler(returnableEpiService)

	authMiddleware := middleware.NewAuthMiddleware(token)

	storageService := service.NewStorageService(domain.StorageAuthentication{
		ObjectId: container.Onedrive.ObjectId,
		ClientId: container.Onedrive.ClientId,
		Username: container.Onedrive.Username,
		Password: container.Onedrive.Password,
		Scopes:   []string{"https://graph.microsoft.com/.default"},
	}, storageRepo)
	storageHandler := handler.NewStorageHandler(storageService)

	router := http.Router(userHandler, authHandler, authMiddleware, employeeActivitiesHandler, storageHandler, employeeHandler, epiHandler, epiExitsHandler, returnableEpiHandler)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
