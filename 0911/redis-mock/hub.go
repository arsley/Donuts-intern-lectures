package main

type RedisClient interface {
	Publish(channel string, message interface{})
	Subscribe(string) RedisSubscription
}

type Message struct {
	Payload string
}

type RedisSubscription interface {
	Close()
	ReceiveMessage() (*Message, error)
}

type Hub struct {
	hubID     string
	broadcast chan []byte
	// subscription RedisSubscription
	redisClient RedisClient
}

func newHub(hubID string) *Hub {
	return &Hub{
		broadcast: make(chan []byte),
	}
}

// func (h *Hub) sub() {
// 	h.subscription = h.redisClient.Subscribe(h.hubID)
// 	defer h.subscription.Close()
// 	for {
// 		message, err := h.subscription.ReceiveMessage()
// 		if err != nil {
// 			log.Printf("error: %s ", err)
// 			return
// 		}
// 		h.broadcast <- []byte(message.Payload)
// 	}
// }

func (h *Hub) pub(message []byte) {
	h.redisClient.Publish(h.hubID, message)
}
