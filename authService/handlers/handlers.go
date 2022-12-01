package handlers

import (
	"be01/models"

	"github.com/gin-gonic/gin"
)

func ApiLogin(c *gin.Context) {
	models.Login(c)
}

func ApiAuth(c *gin.Context) {
	models.Auth(c)
}
