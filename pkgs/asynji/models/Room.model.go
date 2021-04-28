package models

import (
	"errors"
	"time"

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
	InviteCodes        []string
}

func (r *Room) Save() {
	col := Conn.Collection("Rooms")
	col.Save(r)
}

func CreateNewRoom(roomName string, avatar string, owner string) (string, error) {
	col := Conn.Collection("Rooms")
	var newRoom = &Room{
		RoomName:    roomName,
		Avatar:      avatar,
		Owner:       owner,
		InviteCode:  xid.New().String(),
		Hidden:      true,
		Members:     []string{owner},
		Messages:    []Message{},
		InviteCodes: []string{},
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

func (r *Room) InviteNewMembers(invites []string, sender string) string {
	message := ""
	for idx := range invites {
		u, err := GetUser(invites[idx])
		if err != nil {
			message += invites[idx] + "Skiped because user not found.\n"
			continue
		}
		inviteId := xid.New().String()
		u.Updates = append(u.Updates, Update{
			Invite: Invite{From: sender, To: r.Id.Hex(), When: time.Now().Format(time.RubyDate), InviteId: inviteId},
		})
		u.Save()
		r.InviteCodes = append(r.InviteCodes, inviteId)
	}
	r.Save()
	if message == "" {
		message = "All invites sended"
	}
	return message
}

func (r *Room) AcceptInvite(inviteId string, user *User) error {
	if yeah, idx := StringInSlice(inviteId, r.InviteCodes); yeah {
		r.Members = append(r.Members, user.Id.Hex())
		r.FcmTokens = append(r.FcmTokens, user.FcmTokens...)
		Remove(r.InviteCodes, idx)
		r.Save()
		return nil
	} else {
		return errors.New("Invite code invalid")
	}
}

func Remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
