package routes

import (
	"github.com/gin-gonic/gin"
)


func IndexRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Hello men",
	})
}

