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
	owner.Rooms = append(owner.Rooms, id)
	owner.Save()
	c.JSON(201, gin.H{
		"roomId": id,
	})
}

type InviteRoomBody struct {
	InvitedList []string `json:"invited_list"`
}

// [POST] /room/:roomid/invite
func InviteRoomRoute(c *gin.Context) {
	var body InviteRoomBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	room, err := models.GetRoom(c.Param("roomid"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if yeah, _ := models.StringInSlice(c.GetString("userId"), room.Members); yeah {
		if exists, _ := models.StringInSlice(c.GetString("userId"), body.InvitedList); exists {
			c.JSON(400, gin.H{
				"error": "You can't invite yourself",
			})
			return
		}
		message := room.InviteNewMembers(body.InvitedList, c.GetString("userId"))
		if message == "All invites sended" {
			c.JSON(201, gin.H{
				"message": message,
			})
			return
		}
		c.JSON(200, gin.H{
			"error":   message,
			"message": "Some users skipped",
		})
		return
	}
	c.JSON(400, gin.H{
		"error": "Room not found",
	})
}

// [GET] "/:roomid/invite/:inviteid/resolve"
func AcceptInviteRoute(c *gin.Context) {
	roomid := c.Param("roomid")
	inviteid := c.Param("inviteid")
	user, err := models.GetUser(c.GetString("userId"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	room, err := models.GetRoom(roomid)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = room.AcceptInvite(inviteid, user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Welcome to new room",
	})
}

// [GET] "/:roomid/invite/:inviteid/discard"
func DiscardInviteRoute(c *gin.Context) {
	roomid := c.Param("roomid")
	inviteid := c.Param("inviteid")
	room, err := models.GetRoom(roomid)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user, err := models.GetUser(c.GetString("userId"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = room.DiscardInvite(inviteid, user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "Discarded",
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
	if yeah, _ := models.StringInSlice(sender.Id.Hex(), room.Members); yeah {
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
		return
	}
	c.JSON(400, gin.H{
		"message": "Room does not exists",
	})
}

// [GET] "/:roomid"
func GetInfoAboutRoom(c *gin.Context) {
	roomid := c.Param("roomid")
	room, err := models.GetRoom(roomid)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if room.Hidden {
		user, err := models.GetUser(c.Param("userId"))
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Room not found",
			})
			return
		}
		if yeah, _ := models.StringInSlice(user.Id.Hex(), room.Members); !yeah {
			c.JSON(500, gin.H{
				"error": "Room not found",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"message": "Successfull",
		"room":    room,
	})
}

// [GET] "/:roomid/count/messages"
func GetCountMessagesOfRoom(c *gin.Context) {
	roomid := c.Param("roomid")
	room, err := models.GetRoom(roomid)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	count, err := room.GetCountMessages()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if room.Hidden {
		user, err := models.GetUser(c.Param("userId"))
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Room not found",
			})
			return
		}
		if yeah, _ := models.StringInSlice(user.Id.Hex(), room.Members); !yeah {
			c.JSON(500, gin.H{
				"error": "Room not found",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"message": "Successfull",
		"count":   count,
	})
}
