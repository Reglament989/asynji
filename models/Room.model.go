package models

import (
	"github.com/kamva/mgm/v3"
	"github.com/rs/xid"
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
	Members  []string
	Messages []Message
}

func CreateNewRoom(roomName string, avatar string, owner string) (string, error) {
	id := xid.New().String()
	var newRoom = &Room{
		Id:       id,
		RoomName: roomName,
		Avatar:   avatar,
		Owner:    owner,
		Members:  []string{owner},
		Messages: []Message{},
	}
	err := mgm.Coll(newRoom).Create(newRoom)
	if err != nil {
		return "", err
	}
	return id, nil
}
