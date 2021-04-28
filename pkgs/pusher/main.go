package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Reglament989/asynji/pkgs/pusher/api"
	"github.com/Reglament989/asynji/pkgs/pusher/avro"

	"github.com/go-redis/redis/v8"
)

const pusherChannel = "pusher"

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_URL"),
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	status := rdb.Ping(ctx)
	if status.Err() != nil {
		panic(status.Err())
	}
	println("Pusher has started")
	sub := rdb.Subscribe(ctx, pusherChannel)

	channel := sub.Channel()

	for msg := range channel {
		message, err := avro.UnmarshalMessage([]byte(msg.Payload))
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		if message != nil {
			api.SendNotify([]string{"sad"}, message.Hello, "testify")
		}
	}
}
