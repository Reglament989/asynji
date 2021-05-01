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
	statsScope := r.Group("/stats")

	statsScope.GET("/users", router.GetInfoAboutAllUsers)

	statsScope.GET("/rooms", router.GetStatsOfAllRooms)

	userScope := r.Group("/user", middlewares.Auth())

	userUpload := userScope.Group("/upload")

	userUpload.PUT("/fcm", router.UploadFcmToken)

	userUpload.PUT("/public", router.UploadPublicKey)

	// userUpload.PUT("/avatar")

	userScope.GET("/:userid/public", router.GetPublicKey)

	authScope := r.Group("/auth")

	authScope.POST("/login", router.LoginRoute)

	authScope.PATCH("/refresh", router.RefreshRoute)

	authScope.PUT("/registration", router.CreateUserRoute)

	roomScope := r.Group("/room", middlewares.Auth())

	roomScope.GET("/:roomid", router.GetInfoAboutRoom)

	roomScope.GET("/:roomid/count/messages", router.GetCountMessagesOfRoom)

	roomScope.POST("/create", router.CreateRoomRoute)

	roomScope.POST("/:roomid/invite", router.InviteRoomRoute)

	roomScope.GET("/:roomid/invite/:inviteid/resolve", router.AcceptInviteRoute)

	roomScope.DELETE("/:roomid/invite/:inviteid/discard", router.DiscardInviteRoute)

	roomScope.POST("/:roomid/send", router.NewMessageRoute)

	globalScope := r.Group("/_", middlewares.Auth())

	globalScope.GET("/sync", router.GetLatestUpdates)

	// globalScope.PUT("/upload")

	return r
}
