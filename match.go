package main

import (
	"context"
	"encoding/json"
	"fmt"

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

	strUserID1 := strconv.Itoa(match.UserID1)
	strUserID2 := strconv.Itoa(match.UserID2)
	checkUser1, _ := rc.Get(context.Background(), "user:player_"+strUserID1).Result()
	checkUser2, _ := rc.Get(context.Background(), "user:player_"+strUserID2).Result()
	fmt.Println(strUserID1, strUserID2)
	if checkUser1 == "" || checkUser2 == "" {
		responseError(w, "User not found")
		return
	}

	if match.Score1 > match.Score2 {
		json.Unmarshal([]byte(checkUser1), &user1)
		user1.Score += 3
		users1 := jsonConvert(w, user1)
		_, user1WinError := rc.Set(context.Background(), "user:player_"+strUserID1, users1, 0).Result()
		if user1WinError != nil {
			log.Fatal("User1 win set error", user1WinError)
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

		json.Unmarshal([]byte(checkUser2), &user2)
		user2.Score += 3

		users2 := jsonConvert(w, user2)
		_, user2WinError := rc.Set(context.Background(), "user:player_"+strUserID2, users2, 0).Result()
		if user2WinError != nil {
			log.Fatal("User2 win set error", user2WinError)
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

		json.Unmarshal([]byte(checkUser1), &user1)
		user1.Score += 1

		users1 := jsonConvert(w, user1)
		_, user1DrawError := rc.Set(context.Background(), "user:player_"+strUserID1, users1, 0).Result()
		if user1DrawError != nil {
			log.Fatal("User1 draw set error", user1DrawError)
			return
		}
		rz := &redis.Z{
			Score:  float64(user1.Score),
			Member: user1.ID,
		}
		rc.ZAdd(context.Background(), "leaderboard", *rz).Result()

		json.Unmarshal([]byte(checkUser2), &user2)
		user2.Score += 1

		users2 := jsonConvert(w, user2)
		_, user2DrawError := rc.Set(context.Background(), "user:player_:"+strUserID2, users2, 0).Result()
		if user2DrawError != nil {
			log.Fatal("User2 draw set error", user2DrawError)
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
