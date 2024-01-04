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

	checkuserid, _ := rc.Get(context.Background(), "user:"+userLogin.UserName).Result()
	val, _ := rc.Get(context.Background(), "user:"+checkuserid).Result()
	json.Unmarshal([]byte(val), &user)
	userLogin.Password = md5Encode(userLogin.Password)
	intuserid, _ := strconv.Atoi(checkuserid)
	if user.Pwd == userLogin.Password {

		sm := SuccessMessage{
			ID:       intuserid,
			UserName: userLogin.UserName,
		}

		responseSuccess(w, sm)
		return
	}

	responseError(w, "Wrong username or password")
}
