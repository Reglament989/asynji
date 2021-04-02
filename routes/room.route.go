package routes

import (
	"gin_msg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRoomBody struct {
	RoomName string `json:"roomName"`
	Avatar   string `json:"avatar"`
}

func CreateRoomRoute(c *gin.Context) {
	var body CreateRoomBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	owner, err1 := models.GetUser(c.GetString("userId"))
	if err1 != nil {
		c.JSON(500, gin.H{"error": err1.Error()})
		return
	}
	id, err := models.CreateNewRoom(body.RoomName, body.Avatar, owner.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{
		"roomId": id,
	})
}

func InviteRoomRoute(c *gin.Context) {

}
