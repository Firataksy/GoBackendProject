package main

import (
	"encoding/json"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var userLogin Login

	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userLogin.Password = md5Encode(userLogin.Password)
	user := data[userLogin.UserName]

	if user.Pwd == userLogin.Password && user.UserName == userLogin.UserName {
		responseSuccess(w, user.ID, user.UserName)
		return
	}

	responseError(w, "loginerror")
}
