package main

import (
	"gwService/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/api/login", handlers.ApiLogin)
	r.POST("/api/auth", handlers.ApiAuth)

	r.Run("localhost:3001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
