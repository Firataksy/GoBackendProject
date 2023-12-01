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
	userl[userlogin.UserName] = userlogin
	user := user[userlogin.UserName]

	if user.Pwd == userlogin.Password && user.UserName == userlogin.UserName {
		message := Status.StatTrue(Stat{})
		signlogin.Status, signlogin.Data.ID, signlogin.Data.UserName = message, user.ID, userlogin.UserName
		usersJSON, err := json.Marshal(signlogin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(usersJSON)
		return
	} else {
		stat := Status.StatFalse(Stat{})
		error.Status, error.Message = stat, "Wrong Username or Password"
		errorJSON, _ := json.Marshal(error)
		w.Write(errorJSON)
		return
	}
}
