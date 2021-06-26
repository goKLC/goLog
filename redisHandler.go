package goLog

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisHandler struct {
	Client *redis.Client
	Key    string
}

func NewRedisHandler() *RedisHandler {
	return &RedisHandler{Key: "goLog"}
}

func (rh RedisHandler) Write(log Log) {

	val := fmt.Sprintf("[%v] %v: %v", log.date, log.level, log.message)

	if log.context != nil {
		val = fmt.Sprintf("%v : %v", val, log.context)
	}

	err := rh.Client.LPush(ctx, rh.Key, val).Err()

	if err != nil {
		fmt.Println(err)
	}
}
