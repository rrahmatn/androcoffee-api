package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rrahmatn/androcoffee-api.git/database"
	"github.com/rrahmatn/androcoffee-api.git/routes"
)

func main() {
	db := database.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run() // listen and serve on 0.0.0.0:8080
}
