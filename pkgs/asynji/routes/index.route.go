package routes

import (
	"github.com/Reglament989/asynji/pkgs/asynji/rdb"
	"github.com/gin-gonic/gin"
)

func IndexRoute(c *gin.Context) {
	rdb.SendToPusherChannel("Test", "6090001791b3299b0fb1c871")
	c.JSON(200, gin.H{
		"status": "Hello was sended",
	})
}
