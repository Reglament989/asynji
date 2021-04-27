package routes

import (
	"asynji/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginRoute(c *gin.Context) {
	var body LoginBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, refresh, err := models.NewUserLogin(body.Username, body.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Please try later",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token":        token,
		"refreshToken": refresh,
		"message":      "Welcome back!",
	})
}

type RefreshBody struct {
	RefreshToken string `json:"refresh_token"`
}

func RefreshRoute(c *gin.Context) {
	var body RefreshBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, refresh, err := models.RefreshTokens(body.RefreshToken)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token":        token,
		"refreshToken": refresh,
	})
}
