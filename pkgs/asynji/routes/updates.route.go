package routes

import (
	"github.com/Reglament989/asynji/pkgs/asynji/models"

	"github.com/gin-gonic/gin"
)

type GetLatestUpdatesBody struct {
}

func GetLatestUpdates(c *gin.Context) {
	user, err := models.GetUser(c.GetString("userId"))
	if err != nil {
		c.JSON(401, gin.H{
			"errors": err.Error(),
		})
		return
	}
	fullSync := c.DefaultQuery("full", "no")
	var limit int = c.GetInt("limit")
	var offset int = c.GetInt("offset")
	roomUpdates := make(map[string][]*models.Message)
	if fullSync == "yes" {
		limit = 0
	}
	for idx := range user.Rooms {
		messages, err1 := user.GetMessages(user.Rooms[idx], offset, limit)
		if err1 != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
		}
		roomUpdates[user.Rooms[idx]] = messages
	}
	c.JSON(200, gin.H{
		"room_updates": roomUpdates,
		"user_updates": user.Updates,
	})
}
