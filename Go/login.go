package main

import (
	"encoding/json"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userlogin Login
	var signlogin Signlogin
	var error Error

	err := json.NewDecoder(r.Body).Decode(&userlogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userlogin.Password = md5Encode(userlogin.Password)
	user := user[userlogin.UserName]

	if user.Pwd == userlogin.Password && user.UserName == userlogin.UserName {
		message := Status.StatTrue(Stat{})
		signlogin.Status, signlogin.Data.ID, signlogin.Data.UserName = message, user.ID, userlogin.UserName
		Jsonwrite(w, signlogin)
		return
	} else {
		stat := Status.StatFalse(Stat{})
		error.Status, error.Message = stat, "Wrong Username or Password"
		Jsonwrite(w, error)
		return
	}
}
