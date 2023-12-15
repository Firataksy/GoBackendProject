package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func login(w http.ResponseWriter, r *http.Request) {
	var userLogin Sign
	var user User

	er := json.NewDecoder(r.Body).Decode(&userLogin)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}
	id := strconv.Itoa(userLogin.ID)
	val, _ := rc.Get(context.Background(), "user:"+id).Result()

	json.Unmarshal([]byte(val), &user)
	userLogin.Pwd = md5Encode(userLogin.Pwd)

	if user.Pwd == userLogin.Pwd && user.UserName == userLogin.UserName {

		sm := SuccessMessage{
			ID:       user.ID,
			UserName: userLogin.UserName,
		}

		responseSuccess(w, sm)
		return
	}

	responseError(w, "Wrong username or password")
}
