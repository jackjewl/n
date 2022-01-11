package db

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient *redis.Client

// InitRedis 初始化redis连接客户端
func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	redisClient = rdb
}
