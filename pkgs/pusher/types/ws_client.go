package types

import (
	"github.com/gorilla/websocket"
)

type Event struct {
	RoomTo string
	Body   string
}

type Hub struct {
	Clients      map[*Client]bool
	BroadcastTo  chan *Event
	Register     chan *CustomerOfClient
	TopicManager *TopicManager
	// Unregister requests from clients.
	Unregister chan *CustomerOfClient
}

type Client struct {
	Conn   *websocket.Conn
	Id     string
	Send   chan *Event
	Hub    *Hub
	Topics map[string]bool
}

type CustomerOfClient struct {
	Client *Client
	Topics map[string]bool
}
