package models

import (
	"fmt"
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

func (r *Room) GetCountMessages() (int, error) {
	return Conn.Collection(fmt.Sprintf("%s-%s", r.Id.Hex(), "Messages")).Collection().Count()
}

func (u *User) GetMessages(room string, offset int, limit int) ([]*Message, error) {
	col := Conn.Collection(fmt.Sprintf("%s-%s", room, "Messages"))
	rp := col.Find(nil)
	cursor := rp.Query.Skip(offset)
	// cursor = cursor.Sort()
	cursor = cursor.Limit(limit)
	message := Message{}
	messages := []*Message{}
	iter := cursor.Iter()
	for iter.Next(&message) {
		messages = append(messages, &message)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return messages, nil
}
