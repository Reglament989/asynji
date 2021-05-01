package routes

import (
	"strconv"

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
	limitString := c.DefaultQuery("limit", "0")
	limit, err := strconv.ParseUint(limitString, 0, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Limit can be only > 0 and int64",
		})
		return
	}
	if limit == 0 {
		limit = 60
	}
	offsetString := c.DefaultQuery("offset", "0")
	offset, err := strconv.ParseUint(offsetString, 0, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Offset can be only > 0 and int64",
		})
		return
	}
	roomUpdates := make(map[string][]*models.Message)
	if fullSync == "yes" {
		limit = 2000
	}
	for idx := range user.Rooms {
		messages, err1 := user.GetMessages(user.Rooms[idx], int(offset), int(limit))
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
