package main

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func generateClient() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: "localhost:6380"})
}

func storeFixture(r *redis.Client) {
	r.Set("key", "value", time.Second)
}

func flushAll(r *redis.Client) {
	r.FlushAll()
}

func TestRedisManager_Get(t *testing.T) {
	type fields struct {
		Client *redis.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "If no key-value data, return null-string with error",
			fields: fields{Client: generateClient()},
			args:   args{key: "no-stored-key"},
			want:   "",
		},
		{
			name:   "If key-value stored, return gets value",
			fields: fields{Client: generateClient()},
			args:   args{key: "key"},
			want:   "value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisManager{
				Client: tt.fields.Client,
			}
			flushAll(r.Client)
			storeFixture(r.Client)

			got, err := r.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("RedisManager.Get() error: %v", err)
				t.Errorf("RedisManager.Get() return `%v`, want: %v", got, tt.want)
			}
		})
	}
}

func TestRedisManager_Set(t *testing.T) {
	type fields struct {
		Client *redis.Client
	}
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "If no key-value data, return OK",
			fields: fields{Client: generateClient()},
			args:   args{key: "key", val: "value"},
			want:   "OK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisManager{
				Client: tt.fields.Client,
			}
			flushAll(r.Client)
			storeFixture(r.Client)

			got, err := r.Set(tt.args.key, tt.args.val)
			if got != tt.want {
				t.Errorf("RedisManager.Set() error: %v", err)
				t.Errorf("RedisManager.Set() return `%v`, want: `%v`", got, tt.want)
			}
		})
	}
}
