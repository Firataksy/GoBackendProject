package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func login(w http.ResponseWriter, r *http.Request) {
	var userLogin Login
	var user User

	er := json.NewDecoder(r.Body).Decode(&userLogin)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	checkUserID, _ := rc.Get(context.Background(), "userID:"+userLogin.UserName).Result()
	val, _ := rc.Get(context.Background(), "player_"+checkUserID).Result()
	json.Unmarshal([]byte(val), &user)
	w.Header().Add("token", user.Token)
	userLogin.Password = md5Encode(userLogin.Password)
	intUserID, _ := strconv.Atoi(checkUserID)
	if user.Pwd == userLogin.Password {

		sm := SuccessMessage{
			ID:       intUserID,
			UserName: userLogin.UserName,
		}

		responseSuccess(w, sm)
		return
	}

	responseError(w, "Wrong username or password")
}
