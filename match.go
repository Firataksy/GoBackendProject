package main

import (
	"context"
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func match(w http.ResponseWriter, r *http.Request) {
	var match Match
	var user1 User1
	var user2 User2

	er := json.NewDecoder(r.Body).Decode(&match)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}
	if match.UserID1 == match.UserID2 {
		responseError(w, "It cannot be the same in 2 users")
		return
	}

	strUserID1 := strconv.Itoa(match.UserID1)
	strUserID2 := strconv.Itoa(match.UserID2)
	checkUser1, _ := rc.Get(context.Background(), "player_"+strUserID1).Result()
	checkUser2, _ := rc.Get(context.Background(), "player_"+strUserID2).Result()
	if checkUser1 == "" || checkUser2 == "" {
		responseError(w, "User not found")
		return
	}

	if match.Score1 > match.Score2 {
		json.Unmarshal([]byte(checkUser1), &user1)
		user1.Score += 3
		users1 := jsonConvert(w, user1)
		redisSetJustData(match.UserID1, users1)

		z := &redis.Z{
			Score:  float64(user1.Score),
			Member: user1.ID,
		}

		rc.ZAdd(context.Background(), "leaderboard", *z).Result()
		responseSuccess(w, "")
		return
	}

	if match.Score1 < match.Score2 {

		json.Unmarshal([]byte(checkUser2), &user2)
		user2.Score += 3

		users2 := jsonConvert(w, user2)
		redisSetJustData(match.UserID2, users2)
		z := &redis.Z{
			Score:  float64(user2.Score),
			Member: user2.ID,
		}
		rc.ZAdd(context.Background(), "leaderboard", *z).Result()
		responseSuccess(w, "")
		return
	}

	if match.Score1 == match.Score2 {

		json.Unmarshal([]byte(checkUser1), &user1)
		user1.Score += 1

		users1 := jsonConvert(w, user1)
		redisSetJustData(match.UserID1, users1)

		rz := &redis.Z{
			Score:  float64(user1.Score),
			Member: user1.ID,
		}
		rc.ZAdd(context.Background(), "leaderboard", *rz).Result()

		json.Unmarshal([]byte(checkUser2), &user2)
		user2.Score += 1

		users2 := jsonConvert(w, user2)
		redisSetJustData(match.UserID2, users2)
		z := &redis.Z{
			Score:  float64(user2.Score),
			Member: user2.ID,
		}
		rc.ZAdd(context.Background(), "leaderboard", *z).Result()
		responseSuccess(w, "")
	}
}
