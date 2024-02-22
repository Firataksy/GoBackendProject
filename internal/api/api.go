package api

import (
	"fmt"

	"github.com/my/repo/internal/db"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

type Api struct {
	Rdb *RedisClient
}

func NewApi(redisClient *redis.Client) *Api {
	return &Api{
		Rdb: &RedisClient{Client: redisClient},
	}
}

func ConnectRedis() (*RedisClient, error) {
	rc, err := db.New("localhost:6379")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &RedisClient{Client: rc.Client}, err
}
