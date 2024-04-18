package http

import (
	"america-rental-backend/internal/adapter/http/handler"
	"america-rental-backend/internal/adapter/http/middleware"

	"github.com/gin-gonic/gin"
)

func Router(userHandler handler.UserHandler, authHandler handler.AuthHandler, middleware *middleware.AuthMiddleware, activitiesHandler handler.EmployeeActivityLogHandler, storageHandler handler.StorageHandler, employeeHandler handler.EmployeeHandler) *gin.Engine {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
		}

		ctx.Next()
	})

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
			employeeRoutes.GET("/:id", employeeHandler.GetById)
			employeeRoutes.GET("/all", employeeHandler.GetAll)
			employeeRoutes.GET("/activities", activitiesHandler.Get)
			employeeRoutes.POST("/new", employeeHandler.New)
			employeeRoutes.PUT("/update/:id", employeeHandler.Update)
		}

		onedriveRoutes := api.Group("/onedrive").Use(middleware.AuthenticationMiddleware)
		{
			onedriveRoutes.PUT("/new", storageHandler.Create)
			onedriveRoutes.GET("/all", storageHandler.List)
		}
	}

	return r
}
