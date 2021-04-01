package main

import (
	"fmt"
	encry "gin_msg/encryption"
	router "gin_msg/routes"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	encry.LoadJweKeys()
	InitMongo()
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

	userScope := r.Group("/user")

	userScope.POST("/create", router.CreateUserRoute)
	
	authScope := r.Group("/auth")

	authScope.POST("/login", router.LoginRoute)

	r.Run()
}
