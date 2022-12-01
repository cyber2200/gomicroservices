package models

import (
	"be01/types"
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginRequest types.LoginRequest
	loginResponse := types.LoginResponse{
		Res: "NOK",
	}

	c.BindJSON(&loginRequest)
	db := getDbCon()
	dbRes, err := db.Query("SELECT * FROM `users` WHERE `email` = ?;", loginRequest.Email)
	if err != nil {
		fmt.Println(err)
	}

	var usersEntites []types.UsersEntity
	var userEntity types.UsersEntity
	for dbRes.Next() {
		dbRes.Scan(
			&userEntity.Id,
			&userEntity.Email,
			&userEntity.Password,
		)
		usersEntites = append(usersEntites, userEntity)
	}

	if len(usersEntites) == 0 {
		loginResponse.Msg = "Can't find this Email"
		c.JSON(http.StatusOK, loginResponse)
		return
	}

	inputPasswordHash := fmt.Sprintf("%x", md5.Sum([]byte(loginRequest.Password)))
	if inputPasswordHash != userEntity.Password {
		loginResponse.Msg = "Password is wrong"
		c.JSON(http.StatusOK, loginResponse)
		return
	}

	now := fmt.Sprintf("%d", time.Now().UnixMicro())
	sessionId := md5.Sum([]byte(now))
	sessionIdStr := fmt.Sprintf("%x", sessionId)

	loginResponse.Res = "OK"
	loginResponse.SessionId = sessionIdStr

	_, err = db.Exec("INSERT INTO `users_sessions` SET session_id = ?, user_id = ?", sessionIdStr, userEntity.Id)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, loginResponse)
}

func Auth(c *gin.Context) {
	var authRequest types.AuthRequest

	c.BindJSON(&authRequest)
	if authRequest.SessionId == "" {
		c.JSON(http.StatusOK, gin.H{
			"res": "NOK",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"res": "OK",
	})
}
