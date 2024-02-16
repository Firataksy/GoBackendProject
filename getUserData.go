package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func getUserData(w http.ResponseWriter, r *http.Request) {
	var user User
	headerUserID := r.Header.Get("userid")
	if headerUserID == "" {
		return
	}

	idInt, _ := strconv.Atoi(headerUserID)

	idUserName, _ := rc.Get(context.Background(), "user:"+headerUserID).Result()
	val, _ := rc.Get(context.Background(), idUserName).Result()

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
