package api

import (
	"github.com/Reglament989/asynji/pkgs/pusher/fire"
	"github.com/Reglament989/asynji/pkgs/pusher/mongo"
)

func SendNotify(ids []string, message string) error {
	tokens, err := mongo.SearchFcmByIds(ids)
	if err != nil {
		return err
	}
	fire.PushNotify(map[string]string{
		"payload": message,
	}, tokens)

	return nil
}
