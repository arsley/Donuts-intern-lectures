package main

import (
	"flag"
	"log"

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
		start int64
		stop  int64
	)
	flag.StringVar(&cmd, "cmd", "", "wanna set `cmd`")
	flag.StringVar(&key, "key", "", "wanna set `key`")
	flag.StringVar(&val, "val", "", "wanna set `value`")
	flag.StringVar(&field, "field", "", "wanna set `field`")
	flag.Int64Var(&start, "start", 0, "wanna set `start`")
	flag.Int64Var(&stop, "stop", 0, "wanna set `stop`")
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
	case "hset", "HSET":
		manager.HSet(key, field, val)
	case "hget", "HGET":
		manager.HGet(key, field)
	case "lpush", "LPUSH":
		manager.LPush(key, val)
	case "rpop", "RPOP":
		manager.RPop(key)
	case "lrange", "LRANGE":
		manager.LRange(key, start, stop)
	case "llen", "LLEN":
		manager.LLen(key)
	default:
		log.Fatalln("Command invalid.")
	}
}
