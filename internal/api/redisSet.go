package api

import (
	"context"
	"log"

	"github.com/my/repo/internal/utils"
	"github.com/redis/go-redis/v9"
)

func (rc *RedisClient) redisSetJustData(data *Sign) {
	jsonData := utils.JsonConvert(data)

	_, err := rc.Client.Set(context.Background(), data.UserName, jsonData, 0).Result()
	if err != nil {
		log.Fatal("redis set User data err: ", err)
	}
}

func (rc *RedisClient) redisSetUserNameAndID(username string, id int) {
	strID := utils.JsonConvert(id)
	_, err := rc.Client.Set(context.Background(), "userID:"+username, id, 0).Result()
	if err != nil {
		log.Fatal("redis set User ID err: ", err)
	}

	_, err = rc.Client.Set(context.Background(), "user:"+string(strID), username, 0).Result()
	if err != nil {
		log.Fatal("redis set User ID err: ", err)
	}

}

func (rc *RedisClient) redisSetDataAndID(data *Sign) {
	rc.redisSetJustData(data)
	rc.redisSetUserNameAndID(data.UserName, data.ID)
}

func (rc *RedisClient) redisSetLeaderBoard(user *Sign) {

	z := &redis.Z{
		Score:  float64(user.Score),
		Member: user.ID,
	}

	rc.Client.ZAdd(context.Background(), "leaderboard", *z).Result()
}
