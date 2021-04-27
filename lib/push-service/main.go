package main

import (
	"context"
	"fmt"

	"asynji/lib/push-service/api"
	"asynji/lib/push-service/avro"

	"github.com/go-redis/redis/v8"
)

const pusherChannel = "pusher"

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
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
