package routes

import (
	"github.com/Reglament989/asynji/pkgs/asynji/rdb"

	"github.com/gin-gonic/gin"
)

func IndexRoute(c *gin.Context) {
	rdb.SendToPusherChannel("world")
	c.JSON(200, gin.H{
		"status": "Hello was sended",
	})
}
