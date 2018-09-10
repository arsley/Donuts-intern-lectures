package main

import (
	"log"

	"github.com/go-redis/redis"
)

type Hub struct {
	manager *HubManager

	hubID string

	clients map[*Client]struct{}

	// Inbound messages from the clients.
	broadcast chan []byte

	publish chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	subscription *redis.PubSub
}

func newHub(manager *HubManager, hubID string) *Hub {
	return &Hub{
		hubID:      hubID,
		manager:    manager,
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]struct{}),
	}
}

func (h *Hub) sub() {
	h.subscription = redisClient.Subscribe(h.hubID)
	defer h.subscription.Close()
	for {
		message, err := h.subscription.ReceiveMessage()
		if err != nil {
			log.Printf("error: %s ", err)
			return
		}
		h.broadcast <- []byte(message.Payload)
	}
}

func (h *Hub) pub(message []byte) {
	redisClient.Publish(h.hubID, message)
}

func (h *Hub) run() {
	for {
		select {
		case client, ok := <-h.register:
			if !ok {
				return
			}
			h.clients[client] = struct{}{}
		case client, ok := <-h.unregister:
			if !ok {
				return
			}
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				h.manager.exit <- client
			}
		case message, ok := <-h.broadcast:
			if !ok {
				return
			}
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// func (h *Hub) sub(room string) {
// 	subscription := redisClient.Subscribe(room)
// 	defer subscription.Close()
// 	for {
// 		message, err := subscription.ReceiveMessage()
// 		if err != nil {
// 			break
// 		}
// 		h.broadcast <- []byte(message.Payload)
// 	}
// }

func (h *Hub) stop() {
	close(h.register)
	close(h.unregister)
	close(h.broadcast)
	h.subscription.Close()
}

type HubManager struct {
	enter chan *Client
	exit  chan *Client

	hubs map[string]*Hub
}

func (m *HubManager) run() {
	for {
		select {
		case client := <-m.enter:
			var hub *Hub
			if h, ok := m.hubs[client.hubID]; ok {
				hub = h
			} else {
				hub = newHub(m, client.hubID)
				go hub.run()
				go hub.sub()
				m.hubs[client.hubID] = hub
			}
			client.hub = hub
			hub.register <- client
		case client := <-m.exit:
			if h, ok := m.hubs[client.hubID]; ok {
				if len(h.clients) == 0 {
					h.stop()
					delete(m.hubs, client.hubID)
				}
			}
		}
	}
}
