package routes

import (
	"asynji/src/models"

	"github.com/gin-gonic/gin"
)

type GetLatestUpdatesBody struct {
}

func GetLatestUpdates(c *gin.Context) {
	user, err := models.GetUser(c.GetString("userId"))
	if err != nil {
		c.JSON(200, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"updates": user.Updates,
	})
	user.Updates = []models.Update{}
	user.Saving()
}
