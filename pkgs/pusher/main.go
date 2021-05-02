package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Reglament989/asynji/pkgs/pusher/microservice"
	"github.com/Reglament989/asynji/pkgs/pusher/middleware"
	"github.com/joho/godotenv"

	"github.com/Reglament989/asynji/pkgs/pusher/ws"

	"github.com/go-redis/redis/v8"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

const pusherChannel = "pusher"

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_URL"),
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	status := rdb.Ping(ctx)
	if status.Err() != nil {
		panic(status.Err())
	}
	hub := ws.NewHub()
	go hub.Run()
	sub := rdb.Subscribe(ctx, pusherChannel)
	go microservice.ListenRedis(sub, hub)

	http.Handle("/", middleware.Middleware(
		http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			ws.ServeWs(hub, rw, r)
		}),
		middleware.AuthMiddleware,
	))
	fmt.Printf("Pusher started at %s, redis connected.\n", string(os.Getenv("PORT")))
	errl := http.ListenAndServe(os.Getenv("PORT"), nil)
	if errl != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
