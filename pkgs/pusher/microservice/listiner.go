package microservice

import (
	"encoding/json"
	"fmt"

	"github.com/Reglament989/asynji/pkgs/pusher/api"
	"github.com/Reglament989/asynji/pkgs/pusher/ws"
	"github.com/go-redis/redis/v8"
)

func ListenRedis(sub *redis.PubSub, hub *ws.Hub) {
	channel := sub.Channel()

	for msg := range channel {
		message := &ws.Event{}
		err := json.Unmarshal([]byte(msg.Payload), message)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		api.SendNotify(message.RoomTo, string(message.Body))
	}
}
