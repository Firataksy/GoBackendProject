package api

import (
	"context"
	"encoding/json"
	"log"

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

	checkUserName1, err := rc.Client.Get(context.Background(), "user:"+strUserID1).Result()
	if err != nil {
		log.Fatal("match get userName control1 err :", err)
		return
	}

	checkUserName2, err := rc.Client.Get(context.Background(), "user:"+strUserID2).Result()
	if err != nil {
		log.Fatal("match get userName control2 err :", err)
		return
	}

	checkUser1, err := rc.Client.Get(context.Background(), checkUserName1).Result()
	if err != nil {
		log.Fatal("match get user control1 err :", err)
		return
	}

	checkUser2, err := rc.Client.Get(context.Background(), checkUserName2).Result()
	if err != nil {
		log.Fatal("match get user control2 err :", err)
		return
	}

	if checkUser1 == "" || checkUser2 == "" {
		ResponseFail(w, "User not found")
		return
	}

	if match.Score1 > match.Score2 {
		err := json.Unmarshal([]byte(checkUser1), &user)
		if err != nil {
			log.Fatal("match win user1 unmarshal err :", err)
			return
		}

		rc.win(user)
	}

	if match.Score1 < match.Score2 {
		err := json.Unmarshal([]byte(checkUser2), &user)
		if err != nil {
			log.Fatal("match win user2 unmarshal err :", err)
			return
		}

		rc.win(user)
	}

	if match.Score1 == match.Score2 {
		err := json.Unmarshal([]byte(checkUser1), &user)
		if err != nil {
			log.Fatal("match draw user1 unmarshal err :", err)
			return
		}

		err = json.Unmarshal([]byte(checkUser2), &user)
		if err != nil {
			log.Fatal("match draw user2 unmarshal err :", err)
			return
		}

		rc.draw(user, user)
	}
}
