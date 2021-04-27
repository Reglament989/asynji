package middlewares

import (
	"asynji/encryption"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")
		//normally Authorization the_token_xxx
		strArr := strings.Split(bearToken, " ")
		rawToken := ""
		if len(strArr) == 2 {
			rawToken = strArr[1]
		}
		if rawToken != "" {
			userId, err := encryption.VerifyToken(rawToken)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				c.Abort()
				return
			}
			c.Set("userId", userId)
			c.Next()
		} else {
			c.JSON(401, gin.H{
				"message": "Token not found!",
			})
			c.Abort()
			return
		}

	}
}
