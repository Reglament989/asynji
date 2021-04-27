package main

import (
	"asynji/middlewares"
	router "asynji/routes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	r := gin.New()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("\033[34m[%s]\033[39m %s - %s \033[36m[%d]\033[39m ~ \033[33m%s\033[39m\n",
			param.TimeStamp.Format(time.Stamp),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))
	r.Use(gin.Recovery())

	r.GET("/", router.IndexRoute)

	// r.GET("/ws", ws.ListenChanges)

	userScope := r.Group("/user")

	userScope.POST("/create", router.CreateUserRoute)

	authScope := r.Group("/auth")

	authScope.POST("/login", router.LoginRoute)

	authScope.POST("/refresh", router.RefreshRoute)

	roomScope := r.Group("/room", middlewares.Auth())

	roomScope.POST("/create", router.CreateRoomRoute)

	roomScope.POST("/invite", router.InviteRoomRoute)
	return r
}
