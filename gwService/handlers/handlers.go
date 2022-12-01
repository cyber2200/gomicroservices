package handlers

import (
	"github.com/gin-gonic/gin"
)

func ApiLogin(c *gin.Context) {
	Request(c, "/api/login")
}

func ApiAuth(c *gin.Context) {
	Request(c, "/api/auth")
}

func ApiLogout(c *gin.Context) {
	Request(c, "/api/logout")
}
