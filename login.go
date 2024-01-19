package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var userLogin Login
	var user User

	er := json.NewDecoder(r.Body).Decode(&userLogin)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	val, _ := rc.Get(context.Background(), userLogin.UserName).Result()
	json.Unmarshal([]byte(val), &user)
	userLogin.Password = md5Encode(userLogin.Password)
	if user.Pwd == userLogin.Password {

		sm := SuccessMessage{
			Token:    user.Token,
			UserName: userLogin.UserName,
		}

		responseSuccess(w, sm)
		return
	}

	responseError(w, "Wrong username or password")
}
