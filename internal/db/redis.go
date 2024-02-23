package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(redisClient *redis.Client) *RedisClient {
	return &RedisClient{Client: redisClient}
}

func New(addr string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis ping err :", err)
		return nil, err
	}
	rc := &RedisClient{Client: client}
	return rc, err
}
