package routes

import (
	"gin_msg/models"
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
	token, err := models.NewUserLogin(body.Username, body.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Please try later",
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
		"message": "Welcome back!",
	})
}
