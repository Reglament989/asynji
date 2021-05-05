package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Reglament989/asynji/pkgs/pusher/middleware"
	"github.com/Reglament989/asynji/pkgs/pusher/types"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type JsonTopics struct {
	Topics []string
}

func ServeWs(hub *types.Hub, w http.ResponseWriter, r *http.Request) {
	topics := &JsonTopics{}
	rawTopics := r.Header.Get("X-Topics")
	err := json.Unmarshal([]byte(rawTopics), topics)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	realTopics, err := types.VerifyTopics(topics.Topics, r.Context().Value(middleware.ContextUserKey).(string))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	log.Printf("Topic manager: %v", hub.TopicManager.Topics)
	client := &types.Client{Hub: hub, Conn: conn, Send: make(chan *types.Event), Topics: realTopics, Id: r.Context().Value(middleware.ContextUserKey).(string)}
	client.Hub.Register <- &types.CustomerOfClient{
		Client: client,
		Topics: realTopics,
	}
	log.Println("Client registered")

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.Writer()
}

func NewHub() *types.Hub {
	topicMgr := types.NewTopicManager()
	// log.Printf("All Topics: %v", topicMgr.Topics)
	return &types.Hub{
		Clients:      make(map[*types.Client]bool),
		BroadcastTo:  make(chan *types.Event),
		TopicManager: topicMgr,
		Unregister:   make(chan *types.CustomerOfClient),
		Register:     make(chan *types.CustomerOfClient),
	}
}
