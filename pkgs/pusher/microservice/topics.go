package microservice

import (
	"encoding/json"
	"fmt"

	"github.com/Reglament989/asynji/pkgs/pusher/types"
	"github.com/go-redis/redis/v8"
)

func UpdateTopics(sub *redis.PubSub, t *types.TopicManager) {
	channel := sub.Channel()

	for msg := range channel {
		topic := &types.Topic{}
		err := json.Unmarshal([]byte(msg.Payload), topic)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		t.Update <- topic
	}
}
