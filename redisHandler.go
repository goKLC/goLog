package goLog

import (
	"fmt"
	"github.com/go-redis/redis"
)

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

	err := rh.Client.LPush(rh.Key, val).Err()

	if err != nil {
		fmt.Println(err)
	}
}
