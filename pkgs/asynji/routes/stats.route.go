package routes

import (
	"github.com/Reglament989/asynji/pkgs/asynji/models"
	"github.com/gin-gonic/gin"
)

// [GET] "/stats/rooms"
func GetStatsOfAllRooms(c *gin.Context) {
	count, err := models.GetCountOfAllRooms()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"count": count,
	})
}

// [GET] "/stats/users"
func GetInfoAboutAllUsers(c *gin.Context) {
	count, err := models.GetCountOfAllUsers()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"count": count,
	})
}
