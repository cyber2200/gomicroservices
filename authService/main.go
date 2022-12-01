package main

import (
	"be01/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/api/login", handlers.ApiLogin)
	r.POST("/api/auth", handlers.ApiAuth)
	r.POST("/api/logout", handlers.ApiLogout)

	r.Run("localhost:3002")
}
