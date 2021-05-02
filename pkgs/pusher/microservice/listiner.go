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
		go CheckIfOnlineAndPush(hub, message)
	}
}

func CheckIfOnlineAndPush(hub *ws.Hub, message *ws.Event) {
	clientsOnline := []*ws.Client{}
	for client := range hub.Clients {
		if yeah, idx := StringInSlice(client.Id, message.Recipients); yeah {
			clientsOnline = append(clientsOnline, client)
			message.Recipients = Remove(message.Recipients, idx)
		}
	}
	for idx := range clientsOnline {
		clientsOnline[idx].Send <- message
	}
	api.SendNotify(message.Recipients, string(message.Body))
}

func StringInSlice(a string, list []string) (bool, int) {
	for i, b := range list {
		if b == a {
			return true, i
		}
	}
	return false, 0
}

func Remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
