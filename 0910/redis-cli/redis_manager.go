package main

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

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

func (r *RedisManager) HGet(key, field string) {
	result := r.Client.HGet(key, field)
	log.Println(result.String())
}

func (r *RedisManager) HSet(key, field, val string) {
	result := r.Client.HSet(key, field, val)
	log.Println(result.String())
}

func (r *RedisManager) HIncrBy(key, field string, incr int64) {
	result := r.Client.HIncrBy(key, field, incr)
	log.Println(result.String())
}

func (r *RedisManager) LPush(key, val string) {
	result := r.Client.LPush(key, val)
	log.Println(result.String())
}

func (r *RedisManager) RPop(key string) {
	result := r.Client.RPop(key)
	log.Println(result.String())
}

func (r *RedisManager) LRange(key string, start, stop int64) {
	result := r.Client.LRange(key, start, stop)
	log.Println(result.String())
}

func (r *RedisManager) LLen(key string) {
	result := r.Client.LLen(key)
	log.Println(result.String())
}

func (r *RedisManager) ZAdd(key string, members redis.Z) {
	result := r.Client.ZAdd(key, members)
	log.Println(result.String())
}

func (r *RedisManager) ZRemRangeByRank(key string, start, stop int64) {
	result := r.Client.ZRemRangeByRank(key, start, stop)
	log.Println(result.String())
}

func (r *RedisManager) ZRange(key string, start, stop int64) {
	result := r.Client.ZRange(key, start, stop)
	log.Println(result.String())
}

func (r *RedisManager) ZRevRange(key string, start, stop int64) {
	result := r.Client.ZRevRange(key, start, stop)
	log.Println(result.String())
}

func (r *RedisManager) ZIncrBy(key string, incr float64, member string) {
	result := r.Client.ZIncrBy(key, incr, member)
	log.Println(result.String())
}
