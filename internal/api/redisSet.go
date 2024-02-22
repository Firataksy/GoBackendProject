package api

import (
	"context"
	"log"
	"net/http"

	"github.com/my/repo/internal/utils"
	"github.com/redis/go-redis/v9"
)

func (rc *RedisClient) redisSetJustData(w http.ResponseWriter, data *Sign) {
	jsonData := utils.JsonConvert(w, data)

	_, er := rc.Client.Set(context.Background(), data.UserName, jsonData, 0).Result()
	if er != nil {
		log.Fatal("Set User data err: ", er)
	}
}

func (rc *RedisClient) redisSetUserNameAndID(w http.ResponseWriter, username string, id int) {
	strID := utils.JsonConvert(w, id)
	_, er := rc.Client.Set(context.Background(), "userID:"+username, id, 0).Result()
	if er != nil {
		log.Fatal("Set User ID err: ", er)
	}

	_, err := rc.Client.Set(context.Background(), "user:"+string(strID), username, 0).Result()
	if err != nil {
		log.Fatal("Set User ID err: ", err)
	}

}

func (rc *RedisClient) redisSetDataAndID(w http.ResponseWriter, data *Sign) {
	rc.redisSetJustData(w, data)
	rc.redisSetUserNameAndID(w, data.UserName, data.ID)
}

func (rc *RedisClient) redisSetLeaderBoard(user *Sign) {

	z := &redis.Z{
		Score:  float64(user.Score),
		Member: user.ID,
	}

	rc.Client.ZAdd(context.Background(), "leaderboard", *z).Result()
}
