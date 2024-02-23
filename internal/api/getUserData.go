package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (rc *RedisClient) GetUserData(w http.ResponseWriter, r *http.Request) {
	var user User
	headerUserID := r.Header.Get("userid")
	if headerUserID == "" {
		return
	}

	idInt, err := strconv.Atoi(headerUserID)
	if err != nil {
		log.Fatal("getUserData convert err :", err)
		return
	}

	userName, _ := rc.Client.Get(context.Background(), "user:"+headerUserID).Result()

	val, _ := rc.Client.Get(context.Background(), userName).Result()

	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		log.Fatal("getUserData unmarshal err :", err)
		return
	}
	if user.ID == idInt && idInt != 0 {

		sd := SuccessData{
			ID:       user.ID,
			UserName: user.UserName,
			Name:     user.Name,
			SurName:  user.SurName,
		}

		ResponseSuccess(w, sd)
		return
	}
	ResponseFail(w, "User not found")
}
