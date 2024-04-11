package http

import (
	"america-rental-backend/internal/adapter/http/handler"
	"america-rental-backend/internal/adapter/http/middleware"
	"github.com/gin-gonic/gin"
)

func Router(userHandler handler.UserHandler, authHandler handler.AuthHandler, middleware *middleware.AuthMiddleware, activitiesHandler handler.EmployeeActivityLogHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		userRoutes := api.Group("/users")
		{
			userRoutes.POST("/new", userHandler.Create)
			userRoutes.POST("/login", authHandler.Login)

			authUser := userRoutes.Group("/").Use(middleware.AuthenticationMiddleware)
			{
				authUser.GET("/:id", userHandler.Get)
				authUser.GET("/all", userHandler.GetAll)
				authUser.PUT("/update/:id", userHandler.Update)
				authUser.DELETE("/delete/:id", userHandler.Delete)
			}
		}

		employeeRoutes := api.Group("/employees").Use(middleware.AuthenticationMiddleware)
		{
			employeeRoutes.GET("/activities", activitiesHandler.Get)
		}
	}

	return r
}
