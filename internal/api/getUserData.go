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

	userName, err := rc.Client.Get(context.Background(), "user:"+headerUserID).Result()
	if err != nil {
		log.Fatal("getUserData get username err :", err)
		return
	}

	val, err := rc.Client.Get(context.Background(), userName).Result()
	if err != nil {
		log.Fatal("getUserData get user data err :", err)
		return
	}

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
