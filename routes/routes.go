// app/routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rrahmatn/androcoffee-api.git/controller"
	"gorm.io/gorm"
)

// SetupRoutes initializes all application routes
func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	userHandler := controller.NewUserHandler(db)
	authHandler := controller.NewAuthHandler(db)

	userGroup := router.Group("/")
	{
		userGroup.GET("/", userHandler.GetUsers)
		userGroup.GET("/:id", userHandler.GetUserById)
		userGroup.POST("/", userHandler.AddUser)
	}

	authGroup := router.Group("auth")
	{
		authGroup.POST("/signin", authHandler.Signin)
	}

}
