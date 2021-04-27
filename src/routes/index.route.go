package routes

import (
	"asynji/src/rdb"

	"github.com/gin-gonic/gin"
)

func IndexRoute(c *gin.Context) {
	rdb.SendToPusherChannel("world")
	c.JSON(200, gin.H{
		"status": "Hello was sended",
	})
}
