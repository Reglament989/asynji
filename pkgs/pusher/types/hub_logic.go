package types

import "log"

func (h *Hub) Run() {
	for {
		select {
		case customerInfo := <-h.Register:
			// h.Clients[customerInfo.Client] = true
			log.Printf("Client topics %v\n", customerInfo.Topics)
			for topic := range customerInfo.Topics {
				log.Printf("Register client topic %v\n", customerInfo)
				h.TopicManager.RegisterListener <- &RegisterCreds{
					Topic:  topic,
					Client: customerInfo.Client,
				}
			}
		case customerInfo := <-h.Unregister:
			log.Printf("Unregister of hub %s\n", customerInfo.Client.Id)
			delete(h.Clients, customerInfo.Client)
			for topic := range customerInfo.Client.Topics {
				h.TopicManager.UnregisterListener <- &RegisterCreds{
					Topic:  topic,
					Client: customerInfo.Client,
				}
			}
			close(customerInfo.Client.Send)
		case event := <-h.BroadcastTo:
			log.Printf("New event to %s\n", event.RoomTo)
			topic := h.TopicManager.Topics[event.RoomTo]
			if topic != nil {
				log.Printf("Listiners %v", topic.Listiners)
				log.Printf("Clients %v", h.Clients)
				for client := range topic.Listiners {

					select {
					case client.Send <- event:
					default:
						h.Unregister <- &CustomerOfClient{
							Client: client,
							Topics: client.Topics,
						}
					}
				}
			}

		}
	}
}
