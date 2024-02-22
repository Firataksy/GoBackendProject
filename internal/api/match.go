package api

import (
	"context"
	"encoding/json"

	"net/http"
	"strconv"
)

func (rc *RedisClient) Match(w http.ResponseWriter, r *http.Request) {
	var match Match
	var user *Sign

	er := json.NewDecoder(r.Body).Decode(&match)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}
	if match.UserID1 == match.UserID2 {
		ResponseFail(w, "It cannot be the same in 2 users")
		return
	}

	strUserID1 := strconv.Itoa(match.UserID1)
	strUserID2 := strconv.Itoa(match.UserID2)

	checkUserName1, _ := rc.Client.Get(context.Background(), "user:"+strUserID1).Result()
	checkUserName2, _ := rc.Client.Get(context.Background(), "user:"+strUserID2).Result()
	checkUser1, _ := rc.Client.Get(context.Background(), checkUserName1).Result()
	checkUser2, _ := rc.Client.Get(context.Background(), checkUserName2).Result()

	if checkUser1 == "" || checkUser2 == "" {
		ResponseFail(w, "User not found")
		return
	}

	if match.Score1 > match.Score2 {
		json.Unmarshal([]byte(checkUser1), &user)
		rc.win(w, user)
	}

	if match.Score1 < match.Score2 {
		json.Unmarshal([]byte(checkUser2), &user)
		rc.win(w, user)
	}

	if match.Score1 == match.Score2 {
		json.Unmarshal([]byte(checkUser1), &user)

		json.Unmarshal([]byte(checkUser2), &user)

		rc.draw(w, user, user)
	}
}
