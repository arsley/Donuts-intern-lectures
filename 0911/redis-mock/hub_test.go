package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestHub_pub(t *testing.T) {
	type fields struct {
		hubID     string
		broadcast chan []byte
	}
	type args struct {
		message []byte
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		setupMock func(*MockRedisClient, fields, args)
	}{
		{
			name: "Do publish",
			fields: fields{
				hubID: "test_channel",
			},
			args: args{
				message: []byte("pub"),
			},
			setupMock: func(m *MockRedisClient, fields fields, args args) {
				m.EXPECT().Publish(fields.hubID, args.message).Return()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			redisMock := NewMockRedisClient(ctrl)
			tt.setupMock(redisMock, tt.fields, tt.args)
			h := &Hub{
				hubID:     tt.fields.hubID,
				broadcast: tt.fields.broadcast,
				// subscription: tt.fields.subscription,
				redisClient: redisMock,
			}
			h.pub(tt.args.message)
		})
	}
}

// func TestHub_sub(t *testing.T) {
// 	type fields struct {
// 		hubID     string
// 		broadcast chan []byte
// 		// subscription RedisSubscription
// 		// redisClient  RedisClient
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 		{
// 			name: "",
// 			fields: fields{
// 				hubID:     "test",
// 				broadcast: make(chan []byte, 10),
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &Hub{
// 				hubID:        tt.fields.hubID,
// 				broadcast:    tt.fields.broadcast,
// 				subscription: tt.fields.subscription,
// 				redisClient:  tt.fields.redisClient,
// 			}
// 			h.sub()
// 			messages := [][]byte{}
// 			for msg := range h.broadcast {
// 				messages = append(messages, msg)
// 			}
// 			// mockで流した内容と一致しているか検証する
// 		})
// 	}
// }
