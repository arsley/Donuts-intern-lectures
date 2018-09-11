package main

import (
	"fmt"
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
		// h.broadcast <- []byte(message.Payload)
		msg := Parse([]byte(message.Payload))
		h.broadcast <- Stringify(msg.Username, msg.Content)
	}
}

// hubID as key
func (h *Hub) storeMessage(msg string) {
	redisClient.RPush(h.hubID, msg)
	if l, _ := redisClient.LLen(h.hubID).Result(); l > 10 {
		redisClient.LPop(h.hubID)
	}
}

func (h *Hub) getMessagesOnEnter() []string {
	l, _ := redisClient.LLen(h.hubID).Result()
	msgs, _ := redisClient.LRange(h.hubID, 0, l).Result()
	return msgs
}

// about visit counter
func (h *Hub) visit() {
	rest := redisClient.HIncrBy("visitor", "total", 1)
	resc := redisClient.HIncrBy("visitor", "current", 1)
	_, errt := rest.Result()
	_, errc := resc.Result()
	if errt != nil {
		log.Printf("visit error: %s\n", errt)
	}
	if errc != nil {
		log.Printf("visit error: %s\n", errc)
	}
}

func (h *Hub) leave() {
	res := redisClient.HIncrBy("visitor", "current", -1)
	_, err := res.Result()
	if err != nil {
		log.Printf("visit error: %s\n", err)
	}
}

func (h *Hub) totalVisitor() int64 {
	res := redisClient.HGet("visitor", "total")
	counts, err := res.Int64()
	if err != nil {
		log.Printf("HGet totalVisitor error: %s\n", err)
	}
	return counts
}

func (h *Hub) currentVisitor() int64 {
	res := redisClient.HGet("visitor", "current")
	counts, err := res.Int64()
	if err != nil {
		log.Printf("HGet currentVisitor error: %s\n", err)
	}
	return counts
}

// end of visit util

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
			msgs := h.getMessagesOnEnter()
			for _, msg := range msgs {
				client.send <- []byte(msg)
			}
			h.visit()
			visitor := fmt.Sprintf("{\"visitor\":{\"total\":%v, \"current\":%v}}", h.totalVisitor(), h.currentVisitor())
			client.send <- []byte(visitor)
		case client, ok := <-h.unregister:
			if !ok {
				return
			}
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				h.manager.exit <- client
			}
			h.leave()
		case message, ok := <-h.broadcast:
			if !ok {
				return
			}
			h.storeMessage(string(message))
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
					redisClient.Del(h.hubID)
				}
			}
		}
	}
}
