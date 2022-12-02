package handlers

import (
	"be01/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiLogin(c *gin.Context) {
	models.Login(c)
}

func ApiAuth(c *gin.Context) {
	models.Auth(c)
}

func ApiLogout(c *gin.Context) {
	models.Logout(c)
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"v": "0.2",
	})
}
