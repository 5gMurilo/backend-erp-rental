package main

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/user/handler"
	"america-rental-backend/internal/user/repository"
	"america-rental-backend/internal/user/service"
	"context"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.TODO()
	worker, err := db.InitDb(ctx)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(worker)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	routes := r.Group("/api")
	{
		usrRoutes := routes.Group("/user")
		{
			usrRoutes.GET("/:id", userHandler.Get)
		}
	}

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
