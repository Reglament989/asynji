package mongo

import (
	"github.com/Reglament989/asynji/pkgs/asynji/models"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type SearchResult struct {
	bongo.DocumentBase `bson:",inline"`
	FcmTokens          []string
}

func SearchFcmByRoom(roomToId string) ([]string, error) {
	room := &models.Room{}
	err := Conn.Collection("Rooms").FindById(bson.ObjectIdHex(roomToId), room)
	if err != nil {
		return nil, err
	}
	return room.FcmTokens, nil
}
