package rdb

import (
	"context"
	"log"
	"os"

	avro "github.com/Reglament989/asynji/pkgs/pusher/avro"

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

func SendToPusherChannel(message string) {
	pushMessage := avro.PushMessage{
		Hello: message,
	}
	payload, err := pushMessage.MarshalMessage()
	if err != nil {
		log.Fatalln(err)
	}
	rdb.Publish(ctx, pusherChannel, payload)
}
