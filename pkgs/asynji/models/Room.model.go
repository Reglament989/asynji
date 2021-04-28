package models

import (
	"errors"

	"github.com/go-bongo/bongo"
	"github.com/rs/xid"
	mbson "gopkg.in/mgo.v2/bson"
	// metro "github.com/dgryski/go-metro"
)

type Media struct {
	Url string
}

type Message struct {
	bongo.DocumentBase `bson:",inline"`
	From               string
	Body               string
	Forward            string
	Reply              string
	Media              Media
}

type Room struct {
	bongo.DocumentBase `bson:",inline"`
	RoomName           string
	Avatar             string
	Owner              string
	InviteCode         string
	Hidden             bool
	Members            []string
	Messages           []Message
	FcmTokens          []string
}

func CreateNewRoom(roomName string, avatar string, owner string) (string, error) {
	col := Conn.Collection("Rooms")
	var newRoom = &Room{
		RoomName:   roomName,
		Avatar:     avatar,
		Owner:      owner,
		InviteCode: xid.New().String(),
		Hidden:     true,
		Members:    []string{owner},
		Messages:   []Message{},
	}
	err := col.Save(newRoom)
	if err != nil {
		return "", err
	}
	return newRoom.GetId().Hex(), nil
}

func GetRoom(id string) (*Room, error) {
	col := Conn.Collection("Rooms")
	room := &Room{}

	objId := mbson.ObjectIdHex(id)
	if err := col.FindById(objId, room); err != nil {
		return nil, err
	}
	return room, nil
}

func InviteNewMember(userId string, roomId string) (bool, error) {
	// var room = &Room{}
	// if err := mgm.Coll(room).First(bson.M{"id": roomId}, room); err != nil {
	// 	return false, err
	// }
	// var user = &User{}
	// if err := mgm.Coll(user).First(bson.M{"id": userId}, user); err != nil {
	// 	return false, err
	// }
	// room.FcmTokens = append(room.FcmTokens, user.FcmTokens...)
	// room.Members = append(room.Members, userId)
	// room.Saving()
	return false, errors.New("not implimented")
}
