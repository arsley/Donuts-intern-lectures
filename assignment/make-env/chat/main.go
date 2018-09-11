package main

import (
	"flag"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"

	"log"
	"net/http"
	"strings"
)

var redisClient *redis.Client

func main() {
	// definations
	var port int
	flag.IntVar(&port, "p", 8080, "App port")
	flag.Parse()
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	manager := HubManager{
		hubs:  map[string]*Hub{},
		enter: make(chan *Client),
		exit:  make(chan *Client),
	}
	redisClient = redis.NewClient(&redis.Options{Addr: "redis:6379"})
	// reset visitors
	redisClient.HSet("visitor", "total", 0)
	redisClient.HSet("visitor", "current", 0)

	// main
	go manager.run()
	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			respondText(w, 400, []byte("bad parameter"))
			return
		}
		hubID := parts[2]
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := &Client{
			hubID: hubID,
			conn:  conn,
			send:  make(chan []byte, 256),
		}
		manager.enter <- client
		go client.readPump()
		go client.writePump()
	})

	// run server
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func respondText(w http.ResponseWriter, code int, body []byte) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	w.Write(body)
}
