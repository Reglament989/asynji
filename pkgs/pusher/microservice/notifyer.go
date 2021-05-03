package microservice

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Reglament989/asynji/pkgs/pusher/api"
	"github.com/Reglament989/asynji/pkgs/pusher/types"
	"github.com/go-redis/redis/v8"
)

func ListenRedis(sub *redis.PubSub, hub *types.Hub) {
	channel := sub.Channel()

	for msg := range channel {
		message := &types.Event{}
		err := json.Unmarshal([]byte(msg.Payload), message)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		log.Println(message)
		hub.BroadcastTo <- message
		api.SendNotify(message.RoomTo, string(message.Body))
	}
}
