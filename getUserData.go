package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func getUserData(w http.ResponseWriter, r *http.Request) {
	var user User
	userID := r.Header.Get("userid")
	if userID == "" {
		return
	}

	idInt, _ := strconv.Atoi(userID)

	idUserName, _ := rc.Get(context.Background(), "user:"+userID).Result()
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
	responseError(w, "User not found")
}
