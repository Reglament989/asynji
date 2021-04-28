package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Room) NewMessage(from string, body string) (string, error) {
	col := Conn.Collection(fmt.Sprintf("%s-%s", r.Id.Hex(), "Messages"))
	message := &Message{
		From: from,
		Body: body,
	}
	err := col.Save(message)
	if err != nil {
		return "", err
	}
	return message.Id.Hex(), nil
}

func (u *User) GetMessages(room string, offset int, limit int) ([]*Message, error) {
	col := Conn.Collection(fmt.Sprintf("%s-%s", room, "Messages"))
	rp := col.Find(bson.M{})
	cursor := rp.Query.Skip(offset)
	// cursor = cursor.Sort()
	cursor = cursor.Limit(limit)
	message := &Message{}
	messages := []*Message{}
	for cursor.Iter().Next(message) {
		messages = append(messages, message)
	}
	return messages, nil
}
