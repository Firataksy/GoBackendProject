package main

import (
	"context"
	"encoding/json"

	"log"
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

	struserid1 := strconv.Itoa(match.UserID1)
	struserid2 := strconv.Itoa(match.UserID2)
	checkuser1, _ := rc.Get(context.Background(), "user:"+struserid1).Result()
	checkuser2, _ := rc.Get(context.Background(), "user:"+struserid2).Result()

	if checkuser1 == "" || checkuser2 == "" {
		responseError(w, "User not found")
		return
	}

	if match.Score1 > match.Score2 {
		json.Unmarshal([]byte(checkuser1), &user1)
		user1.Score += 3
		users1 := jsonConvert(w, user1)
		_, user1winerror := rc.Set(context.Background(), "user:"+struserid1, users1, 0).Result()
		if user1winerror != nil {
			log.Fatal("User1 win set error", user1winerror)
			return
		}

		z := &redis.Z{
			Score:  float64(user1.Score),
			Member: user1.ID,
		}

		rc.ZAdd(context.Background(), "leaderboard", *z).Result()
		responseSuccess(w, "")
		return
	}

	if match.Score1 < match.Score2 {

		json.Unmarshal([]byte(checkuser2), &user2)
		user2.Score += 3

		users2 := jsonConvert(w, user2)
		_, user2winerror := rc.Set(context.Background(), "user:"+struserid2, users2, 0).Result()
		if user2winerror != nil {
			log.Fatal("User2 win set error", user2winerror)
			return
		}
		z := &redis.Z{
			Score:  float64(user2.Score),
			Member: user2.ID,
		}
		rc.ZAdd(context.Background(), "leaderboard", *z).Result()
		responseSuccess(w, "")
		return
	}

	if match.Score1 == match.Score2 {

		json.Unmarshal([]byte(checkuser1), &user1)
		user1.Score += 1

		users1 := jsonConvert(w, user1)
		_, user1drawerror := rc.Set(context.Background(), "user:"+struserid1, users1, 0).Result()
		if user1drawerror != nil {
			log.Fatal("User1 draw set error", user1drawerror)
			return
		}
		rz := &redis.Z{
			Score:  float64(user1.Score),
			Member: user1.ID,
		}
		rc.ZAdd(context.Background(), "leaderboard", *rz).Result()

		json.Unmarshal([]byte(checkuser2), &user2)
		user2.Score += 1

		users2 := jsonConvert(w, user2)
		_, user2drawerror := rc.Set(context.Background(), "user:"+struserid2, users2, 0).Result()
		if user2drawerror != nil {
			log.Fatal("User2 draw set error", user2drawerror)
			return
		}
		z := &redis.Z{
			Score:  float64(user2.Score),
			Member: user2.ID,
		}
		rc.ZAdd(context.Background(), "leaderboard", *z).Result()
		responseSuccess(w, "")
	}
}
