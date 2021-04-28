package main

import (
	"fmt"
	"time"

	"github.com/Reglament989/asynji/pkgs/asynji/middlewares"
	router "github.com/Reglament989/asynji/pkgs/asynji/routes"

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

	userScope.POST("/upload/fcm", router.UploadFcmToken)

	// userScope.POST("/upload/public", router.UploadFcmToken)

	// userScope.POST("/:userid/public", router.UploadFcmToken)

	authScope := r.Group("/auth")

	authScope.POST("/login", router.LoginRoute)

	authScope.POST("/refresh", router.RefreshRoute)

	roomScope := r.Group("/room", middlewares.Auth())

	roomScope.POST("/create", router.CreateRoomRoute)

	roomScope.POST("/:roomid/invite", router.InviteRoomRoute)

	roomScope.GET("/:roomid/invite/:inviteid/resolve", router.AcceptInviteRoute)

	roomScope.POST("/:roomid/invite/:inviteid/discard", router.DiscardInviteRoute)

	roomScope.POST("/:roomid/send", router.NewMessageRoute)

	r.GET("/sync", middlewares.Auth(), router.GetLatestUpdates)
	return r
}
