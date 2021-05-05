package rdb

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/Reglament989/asynji/pkgs/pusher/types"
	"github.com/go-redis/redis/v8"
)

const pusherChannel = "pusher"

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_URL"),
	Password: os.Getenv("REDIS_PASS"), // no password set
	DB:       0,                       // use default DB
})

func VerifyRdbConnection() {
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	println("\033[32mRedis has connected!\033[39m")
}

func SendToPusherChannel(message string, roomTo string) {
	pushMessage := types.Event{
		RoomTo: roomTo,
		Body:   message,
	}
	payload, err := json.Marshal(pushMessage)
	if err != nil {
		log.Fatalln(err)
	}
	rdb.Publish(ctx, pusherChannel, payload)
}
