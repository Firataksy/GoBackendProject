package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func (rc *RedisClient) GetUserData(w http.ResponseWriter, r *http.Request) {
	var user User
	headerUserID := r.Header.Get("userid")
	if headerUserID == "" {
		return
	}

	idInt, _ := strconv.Atoi(headerUserID)

	userName, _ := rc.Client.Get(context.Background(), "user:"+headerUserID).Result()
	val, _ := rc.Client.Get(context.Background(), userName).Result()

	json.Unmarshal([]byte(val), &user)
	if user.ID == idInt && idInt != 0 {

		sd := SuccessData{
			ID:       user.ID,
			UserName: user.UserName,
			Name:     user.Name,
			SurName:  user.SurName,
		}

		responseSuccess(w, sd)
		return
	}
	responseFail(w, "User not found")
}
