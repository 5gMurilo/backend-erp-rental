package http

import (
	"america-rental-backend/internal/adapter/http/handler"
	"america-rental-backend/internal/adapter/http/middleware"

	"github.com/gin-gonic/gin"
)

func Router(userHandler handler.UserHandler, authHandler handler.AuthHandler, middleware *middleware.AuthMiddleware, activitiesHandler handler.EmployeeActivityLogHandler, storageHandler handler.StorageHandler, employeeHandler handler.EmployeeHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		userRoutes := api.Group("/users")
		{
			userRoutes.POST("/login", authHandler.Login)
			userRoutes.POST("/new", userHandler.Create)

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
			employeeRoutes.GET("/all", employeeHandler.GetAll)
			employeeRoutes.GET("/:id", employeeHandler.GetById)
			employeeRoutes.POST("/new", employeeHandler.New)
		}

		onedriveRoutes := api.Group("/onedrive").Use(middleware.AuthenticationMiddleware)
		{
			onedriveRoutes.PUT("/new", storageHandler.Create)
			onedriveRoutes.GET("/all", storageHandler.List)
		}
	}

	return r
}
