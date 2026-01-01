package main

import (
	"blog-platform/config"
	"blog-platform/database"
	"blog-platform/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.Connect()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":" + config.GetEnv("Server_Port"))
}
