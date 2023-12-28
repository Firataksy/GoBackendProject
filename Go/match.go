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
	if match.Userid1 == match.Userid2 {
		responseError(w, "It cannot be the same in 2 users")
		return
	}

	struserid1 := strconv.Itoa(match.Userid1)
	struserid2 := strconv.Itoa(match.Userid2)
	checkuser1, _ := rc.Get(context.Background(), "user:"+struserid1).Result()
	checkuser2, _ := rc.Get(context.Background(), "user:"+struserid2).Result()

	if checkuser1 == "" || checkuser2 == "" {
		responseError(w, "User not found")
		return
	}

	if match.Score1 > match.Score2 {
		rc.ZRem(context.Background(), "leaderboard", user1.UserName)
		json.Unmarshal([]byte(checkuser1), &user1)

		user1.Puan += 3
		z := &UserLeaderBoard{
			UserName: user1.UserName,
			ID:       user1.ID,
			Puan:     user1.Puan,
		}
		users1 := jsonConvert(w, user1)
		_, user1winerror := rc.Set(context.Background(), "user:"+struserid1, users1, 0).Result()
		if user1winerror != nil {
			log.Fatal("User1 win set error", user1winerror)
			return
		}
		rc.Set(context.Background(), "leaderboard", z, 0)

		responseSuccess(w, "")
		return
	}

	if match.Score1 < match.Score2 {
		rc.ZRem(context.Background(), "leaderboard", user2.UserName)
		json.Unmarshal([]byte(checkuser2), &user2)
		user2.Puan += 3
		z := &redis.Z{
			Score:  float64(user2.Puan),
			Member: user2.UserName,
		}
		users2 := jsonConvert(w, user2)
		_, user2winerror := rc.Set(context.Background(), "user:"+struserid2, users2, 0).Result()
		if user2winerror != nil {
			log.Fatal("User2 win set error", user2winerror)
			return
		}

		rc.ZAdd(context.Background(), "leaderboard", *z)
		responseSuccess(w, "")
	}

	if match.Score1 == match.Score2 {
		rc.ZRem(context.Background(), "leaderboard", user1.UserName)
		json.Unmarshal([]byte(checkuser1), &user1)
		user1.Puan += 1
		rz := &redis.Z{
			Score:  float64(user1.Puan),
			Member: user1.UserName,
		}
		users1 := jsonConvert(w, user1)
		_, user1drawerror := rc.Set(context.Background(), "user:"+struserid1, users1, 0).Result()
		if user1drawerror != nil {
			log.Fatal("User1 draw set error", user1drawerror)
			return
		}
		rc.ZAdd(context.Background(), "leaderboard", *rz)

		rc.ZRem(context.Background(), "leaderboard", user2.UserName)
		json.Unmarshal([]byte(checkuser2), &user2)
		user2.Puan += 1
		z := &redis.Z{
			Score:  float64(user2.Puan),
			Member: user2.UserName,
		}
		users2 := jsonConvert(w, user2)
		_, user2drawerror := rc.Set(context.Background(), "user:"+struserid2, users2, 0).Result()
		if user2drawerror != nil {
			log.Fatal("User2 draw set error", user2drawerror)
			return
		}
		rc.ZAdd(context.Background(), "leaderboard", *z)
		responseSuccess(w, "")
	}
}
