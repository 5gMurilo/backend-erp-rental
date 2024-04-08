package http

import (
	"america-rental-backend/internal/adapter/http/handler"
	"github.com/gin-gonic/gin"
)

func Router(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/:id", userHandler.Get)
		userRoutes.GET("/all", userHandler.GetAll)
		userRoutes.POST("/new", userHandler.Create)
		userRoutes.PUT("/update/:id", userHandler.Update)
		userRoutes.DELETE("/delete/:id", userHandler.Delete)
	}

	return r
}
