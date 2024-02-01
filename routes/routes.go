// app/routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rrahmatn/androcoffee-api.git/controller"
	"gorm.io/gorm"
)

// SetupRoutes initializes all application routes
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize handlers with the database instance

	userHandler := controller.NewUserHandler(db)

	// Example routes for users
	userGroup := router.Group("/")
	{
		userGroup.GET("/", userHandler.GetUsers)
		userGroup.GET("/:id", userHandler.GetUserById)
		userGroup.POST("/", userHandler.AddUser)
	}

}
