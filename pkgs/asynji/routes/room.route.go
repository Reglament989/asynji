package routes

import (
	"net/http"

	"github.com/Reglament989/asynji/pkgs/asynji/models"
	val "github.com/Reglament989/asynji/pkgs/asynji/validators"

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
	validation := val.MultiValidation{}
	roomName := val.Validator{Data: body.RoomName, Name: "room name"}
	roomName.Length(4, 18)
	validation.Add(roomName)
	avatar := val.Validator{Data: body.Avatar, Name: "avatar"}
	avatar.IsUrl()
	validation.Add(avatar)
	errors := validation.Result()
	if errors != nil {
		stringsOfErrors := []string{}
		for i := range errors {
			stringsOfErrors = append(stringsOfErrors, errors[i].Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": stringsOfErrors,
		})
		return
	}
	owner, err1 := models.GetUser(c.GetString("userId"))
	if err1 != nil {
		c.JSON(500, gin.H{"error": err1.Error()})
		return
	}
	id, err := models.CreateNewRoom(body.RoomName, body.Avatar, owner.Id.Hex())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{
		"roomId": id,
	})
}

type InviteRoomBody struct {
	roomIdToJoin string
}

func InviteRoomRoute(c *gin.Context) {
	var body InviteRoomBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	invited, err := models.InviteNewMember(c.GetString("userId"), body.roomIdToJoin)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if invited {
		c.JSON(201, gin.H{
			"message": "Welcome to new room",
		})
		return
	}
	c.JSON(401, gin.H{
		"message": "You not invited to room",
	})
}

type SendMessageBody struct {
	Body string `json:"body"`
}

// [POST] /room/:roomid/send
func NewMessageRoute(c *gin.Context) {
	var body SendMessageBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	room, err := models.GetRoom(c.Param("roomid"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	sender, err1 := models.GetUser(c.GetString("userId"))
	if err1 != nil {
		c.JSON(500, gin.H{"error": err1.Error()})
		return
	}
	if models.StringInSlice(sender.Id.Hex(), room.Members) {
		ids, err := room.NewMessage(sender.Id.Hex(), body.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": err1.Error()})
			return
		}

		// rdb.SendToPusherChannel(room.FcmTokens)

		c.JSON(201, gin.H{
			"message": "Sended",
			"id":      ids,
		})
	}
	c.JSON(400, gin.H{
		"message": "Room does not exists",
	})
}
