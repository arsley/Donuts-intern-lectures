package main

import (
	"flag"
	"log"
	"time"

	"github.com/go-redis/redis"
)

// RedisManager : to handle Redis
type RedisManager struct {
	Client *redis.Client
}

func main() {
	// flag defination, settings
	var (
		cmd   string
		key   string
		val   string
		field string
	)
	flag.StringVar(&cmd, "cmd", "", "wanna set `cmd`")
	flag.StringVar(&key, "key", "", "wanna set `key`")
	flag.StringVar(&val, "val", "", "wanna set `value`")
	flag.StringVar(&field, "field", "", "wanna set `field`")
	flag.Parse()

	// if cmd == "" || key == "" || val == "" {
	// 	flag.Usage()
	// 	return
	// }
	// flag end

	// connect to redis (on docker)
	manager := &RedisManager{
		Client: redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
	}
	// connect end

	// operate
	switch cmd {
	case "set", "SET":
		manager.Set(key, val)
	case "get", "GET":
		manager.Get(key)
	case "del", "DEL":
		manager.Del(key)
	default:
		log.Fatalln("Command invalid.")
	}
}

// Get : get value from key
func (r *RedisManager) Get(key string) {
	result := r.Client.Get(key)
	log.Println(result.String())
}

// Set : set key with value
func (r *RedisManager) Set(key string, val string) {
	result := r.Client.Set(key, val, time.Minute)
	log.Println(result.String())
}

// Del : delete key from key
func (r *RedisManager) Del(key string) {
	result := r.Client.Del(key)
	log.Println(result.String())
}
