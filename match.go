package main

import (
	"context"
	"encoding/json"

	"net/http"
	"strconv"
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
		redisSetJustData(w, match.UserID1, user1)
		redisZSet(user1.Score, user1.ID)
		responseSuccess(w, "")
	}

	if match.Score1 < match.Score2 {
		json.Unmarshal([]byte(checkUser2), &user2)
		user2.Score += 3
		redisSetJustData(w, match.UserID2, user2)
		redisZSet(user2.Score, user2.ID)
		responseSuccess(w, "")
	}

	if match.Score1 == match.Score2 {
		json.Unmarshal([]byte(checkUser1), &user1)
		user1.Score += 1
		redisSetJustData(w, match.UserID1, user1)
		redisZSet(user1.Score, user1.ID)

		json.Unmarshal([]byte(checkUser2), &user2)
		user2.Score += 1
		redisSetJustData(w, match.UserID2, user2)
		redisZSet(user2.Score, user2.ID)
		responseSuccess(w, "")
	}
}
