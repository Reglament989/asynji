package routes

import (
	"net/http"

	models "asynji/models"

	"github.com/gin-gonic/gin"
)

type CreateUserBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func CreateUserRoute(c *gin.Context) {
	var body CreateUserBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if userId, err := models.NewUser(body.Username, body.Email, body.Password, ""); err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"userId": userId,
		})
	}
}
