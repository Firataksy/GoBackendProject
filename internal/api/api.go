package api

import (
	"fmt"

	"github.com/my/repo/internal/db"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func ConnectRedis() (*RedisClient, error) {
	rc, err := db.New("localhost:6379")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &RedisClient{Client: rc.Client}, err
}
