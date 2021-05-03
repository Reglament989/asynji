package types

import (
	"log"

	"github.com/Reglament989/asynji/pkgs/asynji/models"
	"github.com/Reglament989/asynji/pkgs/pusher/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Topic struct {
	Id        string
	FcmTokens []string
	Listiners map[*Client]bool
}

type RegisterCreds struct {
	Topic  string
	Client *Client
}

type TopicManager struct {
	Topics             map[string]*Topic
	RegisterListener   chan *RegisterCreds
	UnregisterListener chan *RegisterCreds
	Update             chan *Topic
	Remove             chan *Topic
}

func (t *TopicManager) Run() {
	for {
		select {
		case update := <-t.Update:
			t.Topics[update.Id] = update
		case remove := <-t.Remove:
			delete(t.Topics, remove.Id)
		case creds := <-t.RegisterListener:
			log.Printf("Register %s", creds.Client.Id)
			t.Topics[creds.Topic].Listiners[creds.Client] = true
		case creds := <-t.UnregisterListener:
			log.Printf("Unregister %s", creds.Client.Id)
			delete(t.Topics[creds.Topic].Listiners, creds.Client)
		}
	}

}

func NewTopicManager() *TopicManager {
	topics, err := GetTopics()
	if err != nil {
		panic(err)
	}
	return &TopicManager{
		Topics:             topics,
		Update:             make(chan *Topic),
		Remove:             make(chan *Topic),
		RegisterListener:   make(chan *RegisterCreds),
		UnregisterListener: make(chan *RegisterCreds),
	}
}

func GetTopics() (map[string]*Topic, error) {
	curs := mongo.Conn.Collection("Rooms").Find(nil)
	query := curs.Query.Select(bson.M{"_id": 1, "FcmTokens": 1})
	room := &models.Room{}
	topics := make(map[string]*Topic)
	cursor := query.Iter()
	defer cursor.Close()
	for cursor.Next(&room) {
		topics[room.Id.Hex()] = &Topic{
			Id:        room.Id.Hex(),
			FcmTokens: room.FcmTokens,
			Listiners: make(map[*Client]bool),
		}
	}
	return topics, nil
}

func VerifyTopics(topics []string, userId string) (map[string]bool, error) {
	col := mongo.Conn.Collection("Rooms")
	realTopics := make(map[string]bool)
	for idx := range topics {
		room := &models.Room{}
		err := col.FindById(bson.ObjectIdHex(topics[idx]), room)
		if err != nil {
			return nil, err
		}
		if yeah, _ := models.StringInSlice(userId, room.Members); yeah {
			realTopics[room.Id.Hex()] = true
		}
	}
	return realTopics, nil
}
