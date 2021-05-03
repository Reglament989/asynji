package api

import (
	"github.com/Reglament989/asynji/pkgs/pusher/fire"
	"github.com/Reglament989/asynji/pkgs/pusher/mongo"
)

func SendNotify(roomTo string, message string) error {
	tokens, err := mongo.SearchFcmByRoom(roomTo)
	if err != nil {
		return err
	}
	fire.PushNotify(map[string]string{
		"payload": message,
	}, tokens)

	return nil
}
