package models

import (
	"github.com/kamva/mgm/v3"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

type Message struct {
	mgm.DefaultModel
	Id   string
	From string
	Body string
}

type Room struct {
	mgm.DefaultModel
	Id       string
	RoomName string
	Avatar   string
	Owner    string
	InviteCode string
	Hidden 	bool
	Members  []string
	Messages []Message
	FcmTokens []string
}

func CreateNewRoom(roomName string, avatar string, owner string) (string, error) {
	id := xid.New().String()
	var newRoom = &Room{
		Id:       id,
		RoomName: roomName,
		Avatar:   avatar,
		Owner:    owner,
		InviteCode: xid.New().String(),
		Hidden: true,
		Members:  []string{owner},
		Messages: []Message{},
	}
	err := mgm.Coll(newRoom).Create(newRoom)
	if err != nil {
		return "", err
	}
	return id, nil
}

func InviteNewMember(userId string, roomId string) (bool, error) {
	var room = &Room{}
	if err := mgm.Coll(room).First(bson.M{"id": roomId}, room); err != nil {
		return false, err
	}
	var user = &User{}
	if err := mgm.Coll(user).First(bson.M{"id": userId}, user); err != nil {
		return false, err
	}
	room.FcmTokens = append(room.FcmTokens, user.FcmTokens...)
	room.Members = append(room.Members, userId)
	room.Saving()
	return true, nil
}
